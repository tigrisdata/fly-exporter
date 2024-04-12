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

package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ApplicationJson(t *testing.T) {
	applicationJson := []byte(`
	{
		"id": "jg589zorv4zqkze7",
		"name": "fly-builder-dark-sound-4364",
		"machine_count": 1,
		"network": "default"
	}
	`)
	app := Application{}
	err := json.Unmarshal(applicationJson, &app)
	assert.Equal(t, err, nil)
}

func Test_ApplicationsJson(t *testing.T) {
	applicationsJson := []byte(`
	{
		"total_apps": 30,
		"apps": [
		  {
			"id": "zgrlqnx0jg41vd47",
			"name": "warp-app",
			"machine_count": 5,
			"network": "default"
		  },
		  {
			"id": "vxemq2vr8739gz63",
			"name": "ptest-js",
			"machine_count": 2,
			"network": "default"
		  },
		  {
			"id": "g3zmqxx6jegqdlp4",
			"name": "dev-k3s-cluster-iad-01-cp",
			"machine_count": 3,
			"network": "default"
		  },
		  {
			"id": "8gnr14p4e481y7wm",
			"name": "dev-k3s-cluster-iad-01-ng-10",
			"machine_count": 15,
			"network": "default"
		  },
		  {
			"id": "zgrlqnjylj89vd47",
			"name": "dev-tigris-os-console",
			"machine_count": 1,
			"network": "default"
		  },
		  {
			"id": "yexkqwjegxj9m38d",
			"name": "testbox",
			"machine_count": 4,
			"network": "default"
		  },
		  {
			"id": "vo4x1onjw04ql5yd",
			"name": "dev-tigris-logshippper",
			"machine_count": 2,
			"network": "default"
		  },
		  {
			"id": "ewpm90zrx3017zoy",
			"name": "dev-docker-registry",
			"machine_count": 4,
			"network": "default"
		  },
		  {
			"id": "jg589zjy3lx1kze7",
			"name": "idev-k3s-central-iad-01-cp",
			"machine_count": 3,
			"network": "default"
		  },
		  {
			"id": "jpon17p7jex1dgr4",
			"name": "idev-k3s-regional-iad-01-cp",
			"machine_count": 3,
			"network": "default"
		  },
		  {
			"id": "p7j3qgwzr8x1mz5x",
			"name": "tigris-data-credentials",
			"machine_count": 1,
			"network": "default"
		  },
		  {
			"id": "704e937ek7jq3xgn",
			"name": "idev-tigris-os",
			"machine_count": 6,
			"network": "default"
		  },
		  {
			"id": "jlyv9r5xmjr18xrg",
			"name": "fly-test-app-tigris",
			"machine_count": 0,
			"network": "default"
		  },
		  {
			"id": "jpon17dpzj81dgr4",
			"name": "dev-tigris-internal",
			"machine_count": 3,
			"network": "default"
		  },
		  {
			"id": "p7j3qgo7jlx9mz5x",
			"name": "tigris-gateway-image-builder",
			"machine_count": 0,
			"network": "default"
		  },
		  {
			"id": "jg589zvke2x1kze7",
			"name": "idev-otel-collector",
			"machine_count": 2,
			"network": "default"
		  },
		  {
			"id": "vo4x1ow6pnzql5yd",
			"name": "idev-grafana-agent",
			"machine_count": 4,
			"network": "default"
		  },
		  {
			"id": "r85yqlzk4v092pvl",
			"name": "idev-warp-runner",
			"machine_count": 3,
			"network": "default"
		  },
		  {
			"id": "g3zmqx0y7k6qdlp4",
			"name": "tigris-integration-test-app",
			"machine_count": 0,
			"network": "default"
		  },
		  {
			"id": "mkp298wonkr9orx5",
			"name": "multi-modal-starter-kit",
			"machine_count": 2,
			"network": "default"
		  },
		  {
			"id": "j3g8qvr4k689dvxy",
			"name": "idev-k3s-central-iad-01-ng-1",
			"machine_count": 15,
			"network": "default"
		  },
		  {
			"id": "p7j3qgjpmnxqmz5x",
			"name": "idev-tigris-ecache",
			"machine_count": 8,
			"network": "default"
		  },
		  {
			"id": "yexkqwemrejqm38d",
			"name": "idev-k3s-regional-lax-01-cp",
			"machine_count": 3,
			"network": "default"
		  },
		  {
			"id": "yk4g9k63rxdqzom5",
			"name": "idev-k3s-regional-lax-01-ng-0",
			"machine_count": 10,
			"network": "default"
		  },
		  {
			"id": "36dyqm74gzkqgwme",
			"name": "idev-k3s-regional-iad-01-ng-0",
			"machine_count": 12,
			"network": "default"
		  },
		  {
			"id": "36dyqm767nkqgwme",
			"name": "dev-tigris-ugrafana",
			"machine_count": 2,
			"network": "default"
		  },
		  {
			"id": "jg589zooeozqkze7",
			"name": "idev-tigris-replication-cmd",
			"machine_count": 1,
			"network": "default"
		  },
		  {
			"id": "jg589zooekxqkze7",
			"name": "tigris-replication-cmd-builder",
			"machine_count": 0,
			"network": "default"
		  },
		  {
			"id": "4nml16r40wr1vp2y",
			"name": "tigris-text-to-image",
			"machine_count": 1,
			"network": "default"
		  },
		  {
			"id": "jg589zorv4zqkze7",
			"name": "fly-builder-dark-sound-4364",
			"machine_count": 1,
			"network": "default"
		  }
		]
	  }
	`)

	apps := Applications{}
	err := json.Unmarshal(applicationsJson, &apps)
	assert.Equal(t, err, nil)
}
