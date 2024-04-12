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

type CheckHeader struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

type CheckConfig struct {
	Type          string        `json:"type"`
	Protocol      string        `json:"protocol,omitempty"`
	Port          int           `json:"port,omitempty"`
	Interval      string        `json:"interval"`
	Timeout       string        `json:"timeout"`
	Method        string        `json:"method,omitempty"`
	Path          string        `json:"path,omitempty"`
	TLSSkipVerify *bool         `json:"tls_skip_verify,omitempty"`
	Headers       []CheckHeader `json:"headers,omitempty"`
}

type ConcurrencyConfig struct {
	Type      string `json:"type,omitempty"`
	SoftLimit *int   `json:"soft_limit,omitempty"`
	HardLimit *int   `json:"hard_limit,omitempty"`
}

type InitConfig struct {
	Exec       string   `json:"exec,omitempty"`
	Cmd        []string `json:"cmd,omitempty"`
	Entrypoint []string `json:"entrypoint,omitempty"`
	TTY        bool     `json:"tty"`
}

type GuestConfig struct {
	CPUKind string `json:"cpu_kind"`
	CPUs    uint   `json:"cpus"`
	Memory  uint   `json:"memory_mb"`
}

type MetadataConfig struct {
	FlyCtlVersion      string `json:"fly_flyctl_version,omitempty"`
	FlyPlatformVersion string `json:"fly_platform_version,omitempty"`
	FlyProcessGroup    string `json:"fly_process_group,omitempty"`
	FlyReleaseID       string `json:"fly_release_id,omitempty"`
	FlyReleaseVersion  string `json:"fly_release_version,omitempty"`
}

type MountConfig struct {
	VolumeID string `json:"volume"`
	Path     string `json:"path"`
}

type PortConfig struct {
	Port       *int     `json:"port,omitempty"`
	Handlers   []string `json:"handlers,omitempty"`
	ForceHTTPS *bool    `json:"force_https,omitempty"`
}

type RestartConfig struct {
	Policy     string `json:"policy,omitempty"`
	MaxRetries int    `json:"max_retries,omitempty"`
}

type ServiceConfig struct {
	Protocol     string             `json:"protocol"`                // Networking protocol (TCP/HTTP)
	Concurrency  *ConcurrencyConfig `json:"concurrency,omitempty"`   // Load balancing concurrency settings
	InternalPort uint               `json:"internal_port,omitempty"` // Port the machine VM listens on
	Ports        []PortConfig       `json:"ports,omitempty"`         // Service's ports and associated handler
}

type Config struct {
	Checks   map[string]CheckConfig `json:"checks,omitempty"`
	Env      map[string]string      `json:"env"`
	Image    string                 `json:"image"`
	Init     *InitConfig            `json:"init"`
	Guest    *GuestConfig           `json:"guest"`
	Metadata *MetadataConfig        `json:"metadata,omitempty"`
	Mounts   []MountConfig          `json:"mounts,omitempty"`
	Restart  *RestartConfig         `json:"restart"`
	Services []ServiceConfig        `json:"services,omitempty"`
}
