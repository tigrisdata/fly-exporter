// Copyright 2022-2024 Tigris Data, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package reporter

import (
	"github.com/tigrisdata/fly-exporter/models"
	"github.com/uber-go/tally"
)

// Shared class and methods for metric groups
type metricGroup struct {
	name        string
	parentScope tally.Scope
	reporter    *Reporter // parent reference
	scoped
}

func newMetricGroup(name string, parentScope tally.Scope, reporter *Reporter) *metricGroup {
	m := metricGroup{name: name, parentScope: parentScope, reporter: reporter}
	m.scopes = make(map[string]tally.Scope)
	m.AddScope(parentScope, "default", name)
	return &m
}

// Each group that collects metrics needs to implement the Collectable interface
type Collectable interface {
	GetMetrics(report *models.Report)
}
