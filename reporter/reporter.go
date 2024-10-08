// Copyright 2024 Tigris Data, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package reporter

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"

	fly "github.com/tigrisdata/fly-exporter/client"
	"github.com/tigrisdata/fly-exporter/models"

	pclient "github.com/m3db/prometheus_client_golang/prometheus"
	"github.com/uber-go/tally"
	"github.com/uber-go/tally/prometheus"
	promreporter "github.com/uber-go/tally/prometheus"
)

type Reporter struct {
	closer   io.Closer
	groups   []Collectable
	reporter prometheus.Reporter
	scoped
}

func assertRequiredEnv() {
	requiredEnvVars := []string{
		"ENVIRONMENT",
		"FLY_API_TOKEN",
		"FLY_ORG_SLUG",
	}

	for i := 0; i < len(requiredEnvVars); i++ {
		if _, exists := os.LookupEnv(requiredEnvVars[i]); !exists {
			fmt.Printf("Please set %s\n", requiredEnvVars[i])
			os.Exit(3)
		}
	}
}

func NewReporter() *Reporter {
	assertRequiredEnv()
	reg := pclient.NewRegistry()

	r := Reporter{
		reporter: promreporter.NewReporter(
			promreporter.Options{Registerer: reg},
		),
	}

	// Create root scope
	scope, closer := tally.NewRootScope(tally.ScopeOptions{
		Tags:                   GetBaseTags(),
		CachedReporter:         r.reporter,
		Separator:              promreporter.DefaultSeparator,
		OmitCardinalityMetrics: true,
	}, 1*time.Second)

	// Add closer to reporter
	r.closer = closer

	// Add root scope to reporter
	r.scopes = make(map[string]tally.Scope)
	r.scopes["root"] = scope

	// Add prometheus prefix for root scope
	r.AddScope(r.scopes["root"], "fly", "fly")

	// Add prometheus prefix once for machine scope, collected below
	r.AddScope(r.scopes["fly"], "machine", "machine")

	// Collect metric groups
	r.groups = []Collectable{
		NewMachineMetricGroup(&r),
	}

	return &r
}

func (r *Reporter) getReport() (*models.Report, error) {
	report := models.Report{}

	c := fly.NewClient()
	ctx := context.Background()
	apps, err := c.GetApps(ctx)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(apps); i++ {
		machines, err := c.GetMachines(ctx, &apps[i])
		if err != nil {
			return nil, err
		}

		report.AppData = append(report.AppData, models.AppData{
			App:      apps[i],
			Machines: machines,
		})
	}

	return &report, nil
}

func (r *Reporter) CollectOnce() error {
	report, err := r.getReport()
	if err != nil {
		return err
	}

	if len(r.groups) == 0 {
		slog.Error("no metric groups detected")
	}

	for _, group := range r.groups {
		group.GetMetrics(report)
	}

	return nil
}

func (r *Reporter) Close() {
	slog.Error("failed to close server:", "error", r.closer.Close())
}

func (r *Reporter) ServeHttp() {
	port := os.Getenv("FLY_EXPORTER_PORT")
	if port == "" {
		port = "8081"
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), r.reporter.HTTPHandler()); err != nil {
		slog.Error("server failed:", "error", err)
		os.Exit(1)
	}
}
