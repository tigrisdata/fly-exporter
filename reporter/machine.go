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
	"github.com/tigrisdata/fly-exporter/models"
)

type MachineMetricGroup struct {
	metricGroup
}

func NewMachineMetricGroup(reporter *Reporter) *MachineMetricGroup {
	parentScope := reporter.GetScopeOrExit("fly")
	g := MachineMetricGroup{
		metricGroup: *newMetricGroup("machine", parentScope, reporter),
	}
	g.AddScope(parentScope, "created", "machine")
	g.AddScope(parentScope, "destroyed", "machine")
	g.AddScope(parentScope, "destroying", "machine")
	g.AddScope(parentScope, "replacing", "machine")
	g.AddScope(parentScope, "started", "machine")
	g.AddScope(parentScope, "starting", "machine")
	g.AddScope(parentScope, "stopped", "machine")
	g.AddScope(parentScope, "stopping", "machine")

	return &g
}

func (m *MachineMetricGroup) GetMetrics(report *models.Report) {
	scope := m.GetScopeOrExit("default")
	for i := 0; i < len(report.AppData); i++ {
		appData := report.AppData[i]
		for j := 0; j < len(appData.Machines); j++ {
			machine := report.AppData[i].Machines[j]
			tags := map[string]string{
				"name":    machine.Name,
				"id":      machine.ID,
				"region":  machine.Region,
				"fly_app": appData.App.Name,
			}
			SetGauge(scope, machine.State, tags, 1)
		}
	}
}
