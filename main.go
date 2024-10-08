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

// Collects machine state from fly API and reports in prometheus metrics format

package main

import (
	"sync"

	"github.com/tigrisdata/fly-exporter/provider"
)

func main() {
	var wg sync.WaitGroup

	// Create reporter to serve up prometheus metrics
	m := provider.NewMetricProvider()
	defer m.Close()
	go m.ServeHttp()
	wg.Add(1)

	// Collect data from fly
	go m.Collect()
	wg.Add(1)

	wg.Wait()
}
