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

func TestEmptyEventUnmarshal(t *testing.T) {
	test := []byte(`{}`)
	object := Event{}
	err := json.Unmarshal(test, &object)
	assert.Equal(t, err, nil)
}

func TestEventUnmarshal(t *testing.T) {
	test := []byte(`
	{
	"type": "launch",
	"status": "created",
	"source": "user",
	"timestamp": 1698719723203
	}
	`)
	object := Event{}
	err := json.Unmarshal(test, &object)
	assert.Equal(t, err, nil)
}

func TestEventsUnmarshal(t *testing.T) {
	test := []byte(`
	[
        {
        "type": "start",
        "status": "started",
        "source": "flyd",
        "timestamp": 1698719726615
        },
        {
        "type": "launch",
        "status": "created",
        "source": "user",
        "timestamp": 1698719723203
        }
    ]
	`)
	object := []Event{}
	err := json.Unmarshal(test, &object)
	assert.Equal(t, err, nil)
}

func TestEmptyImageRefUnmarshal(t *testing.T) {
	test := []byte(`{}`)
	object := ImageRef{}
	err := json.Unmarshal(test, &object)
	assert.Equal(t, err, nil)
}

func TestImageRefUnmarshal(t *testing.T) {
	test := []byte(`
	{
		"registry": "registry.fly.io",
		"repository": "idev-tigris-ecache",
		"tag": "deployment-01HSMPM95GC25XTPJDZS9RBZT9",
		"digest": "sha256:b05aa81491ec599c3e52835772d2ac5f1ddbdfbc073b70f2b430199f3a59faa5",
		"labels": {
			"org.opencontainers.image.created": "2023-07-21T23:34:15.259Z",
			"org.opencontainers.image.description": "Kvrocks is a distributed key value NoSQL database that uses RocksDB as storage engine and is compatible with Redis protocol.",
			"org.opencontainers.image.licenses": "Apache-2.0",
			"org.opencontainers.image.revision": "93d2f7e775f155dda9d388a86b89b6684fba1c9a",
			"org.opencontainers.image.source": "https://github.com/tigrisdata/kvrocks",
			"org.opencontainers.image.title": "kvrocks",
			"org.opencontainers.image.url": "https://github.com/tigrisdata/kvrocks",
			"org.opencontainers.image.version": "nightly"
		}
	}
	`)
	object := ImageRef{}
	err := json.Unmarshal(test, &object)
	assert.Equal(t, err, nil)
}

func TestMachineUnmarshal(t *testing.T) {
	test := []byte(`
	{
		"id": "a5c5de9ce64ca12",
		"name": "aged-wind-2649",
		"state": "started",
		"region": "ord",
		"image_ref": {
			"registry": "registry-1.docker.io",
			"repository": "rebelthor/sleep",
			"tag": "latest",
			"digest": "sha256:597c3e12f830132be2aa69b4c0deccb0657ea4253e6d59c6f38e41e9f69a0add"
		},
		"instance_id": "1RREBN3T5K95DK9IVP4XHTTPEY2",
		"private_ip": "fdaa:0:18:a7b:196:e274:9ce1:2",
		"created_at": "2023-10-31T02:30:10Z",
		"updated_at": "2023-10-31T02:35:26Z",
		"config": {},
		"events": [
			{
			"type": "start",
			"status": "started",
			"source": "flyd",
			"timestamp": 1698719726615
			},
			{
			"type": "launch",
			"status": "created",
			"source": "user",
			"timestamp": 1698719723203
			}
		]
	}
	`)
	object := Machine{}
	err := json.Unmarshal(test, &object)
	assert.Equal(t, err, nil)
}

func TestMachinesUnmarshal(t *testing.T) {
	test := []byte(`
	[
	{
		"id": "080e335f175068",
		"name": "long-waterfall-3668",
		"state": "started",
		"region": "iad",
		"instance_id": "01HV05HP06B552N3ET8XTNV1BB",
		"private_ip": "fdaa:2:3d6:a7b:1d8:bba5:5cb7:2",
		"config": {
		"env": { "FLY_PROCESS_GROUP": "app", "PRIMARY_REGION": "iad" },
		"init": {},
		"guest": { "cpu_kind": "shared", "cpus": 4, "memory_mb": 2048 },
		"metadata": {
			"fly_flyctl_version": "0.2.23",
			"fly_platform_version": "v2",
			"fly_process_group": "app",
			"fly_release_id": "QgDeJ3o11V2D7fmo8DMDomQQq",
			"fly_release_version": "10"
		},
		"mounts": [
			{
			"encrypted": true,
			"path": "/data",
			"size_gb": 12,
			"volume": "vol_4503250zj9m31z8r",
			"name": "idev_tigris_ecache_data"
			}
		],
		"services": [
			{
			"internal_port": 6379,
			"autostop": false,
			"autostart": true,
			"force_instance_key": null
			}
		],
		"checks": {
			"kvrocks": {
			"port": 6379,
			"type": "tcp",
			"interval": "15s",
			"timeout": "5s",
			"grace_period": "30s"
			}
		},
		"image": "registry.fly.io/idev-tigris-ecache:deployment-01HSM9K5B3A23TR6Y5KFJ2ZHAZ",
		"restart": {}
		},
		"image_ref": {
		"registry": "registry.fly.io",
		"repository": "idev-tigris-ecache",
		"tag": "deployment-01HSM9K5B3A23TR6Y5KFJ2ZHAZ",
		"digest": "sha256:7223eb7b751069c32c568de24d2fe24d9dded5362cc9f9a666f308a676a84832",
		"labels": {
			"org.opencontainers.image.created": "2023-07-21T23:34:15.259Z",
			"org.opencontainers.image.description": "Kvrocks is a distributed key value NoSQL database that uses RocksDB as storage engine and is compatible with Redis protocol.",
			"org.opencontainers.image.licenses": "Apache-2.0",
			"org.opencontainers.image.revision": "93d2f7e775f155dda9d388a86b89b6684fba1c9a",
			"org.opencontainers.image.source": "https://github.com/tigrisdata/kvrocks",
			"org.opencontainers.image.title": "kvrocks",
			"org.opencontainers.image.url": "https://github.com/tigrisdata/kvrocks",
			"org.opencontainers.image.version": "nightly"
		}
		},
		"created_at": "2024-04-09T00:53:51Z",
		"updated_at": "2024-04-09T00:54:04Z",
		"events": [
		{
			"id": "01HV05J25TE4B5RQQ1SFEEKADS",
			"type": "start",
			"status": "started",
			"source": "flyd",
			"timestamp": 1712624044218
		},
		{
			"id": "01HV05HP1GZ5Z6QDBK22HAT7JP",
			"type": "launch",
			"status": "created",
			"source": "user",
			"timestamp": 1712624031792
		}
		],
		"checks": [
		{
			"name": "kvrocks",
			"status": "passing",
			"output": "Success",
			"updated_at": "2024-04-09T00:54:07.984Z"
		}
		]
	},
	{
		"id": "1857473b7e7178",
		"name": "wandering-paper-6043",
		"state": "started",
		"region": "iad",
		"instance_id": "01HV0F715PH07K046DTKFTZWTR",
		"private_ip": "fdaa:2:3d6:a7b:1b7:f090:404a:2",
		"config": {
		"env": { "FLY_PROCESS_GROUP": "app", "PRIMARY_REGION": "iad" },
		"init": {},
		"guest": { "cpu_kind": "shared", "cpus": 4, "memory_mb": 2048 },
		"metadata": {
			"fly_flyctl_version": "0.2.23",
			"fly_platform_version": "v2",
			"fly_process_group": "app",
			"fly_release_id": "51azqDBYY72aRcNP8bOKKgDaP",
			"fly_release_version": "11"
		},
		"mounts": [
			{
			"encrypted": true,
			"path": "/data",
			"size_gb": 12,
			"volume": "vol_49k30kozowx53n3v",
			"name": "idev_tigris_ecache_data"
			}
		],
		"services": [
			{
			"internal_port": 6379,
			"autostop": false,
			"autostart": true,
			"force_instance_key": null
			}
		],
		"checks": {
			"kvrocks": {
			"port": 6379,
			"type": "tcp",
			"interval": "15s",
			"timeout": "5s",
			"grace_period": "30s"
			}
		},
		"image": "registry.fly.io/idev-tigris-ecache:deployment-01HSMPM95GC25XTPJDZS9RBZT9",
		"restart": {}
		},
		"image_ref": {
		"registry": "registry.fly.io",
		"repository": "idev-tigris-ecache",
		"tag": "deployment-01HSMPM95GC25XTPJDZS9RBZT9",
		"digest": "sha256:b05aa81491ec599c3e52835772d2ac5f1ddbdfbc073b70f2b430199f3a59faa5",
		"labels": {
			"org.opencontainers.image.created": "2023-07-21T23:34:15.259Z",
			"org.opencontainers.image.description": "Kvrocks is a distributed key value NoSQL database that uses RocksDB as storage engine and is compatible with Redis protocol.",
			"org.opencontainers.image.licenses": "Apache-2.0",
			"org.opencontainers.image.revision": "93d2f7e775f155dda9d388a86b89b6684fba1c9a",
			"org.opencontainers.image.source": "https://github.com/tigrisdata/kvrocks",
			"org.opencontainers.image.title": "kvrocks",
			"org.opencontainers.image.url": "https://github.com/tigrisdata/kvrocks",
			"org.opencontainers.image.version": "nightly"
		}
		},
		"created_at": "2024-04-09T03:42:48Z",
		"updated_at": "2024-04-09T03:43:06Z",
		"events": [
		{
			"id": "01HV0F7JQ341HHT5QJFA2MAEKX",
			"type": "start",
			"status": "started",
			"source": "flyd",
			"timestamp": 1712634186467
		},
		{
			"id": "01HV0F71769DCK3PHJ7M75RWVF",
			"type": "launch",
			"status": "created",
			"source": "user",
			"timestamp": 1712634168550
		}
		],
		"checks": [
		{
			"name": "kvrocks",
			"status": "passing",
			"output": "Success",
			"updated_at": "2024-04-09T03:43:22.241Z"
		}
		]
	},
	{
		"id": "148e442b292078",
		"name": "polished-fog-527",
		"state": "started",
		"region": "ord",
		"instance_id": "01HV0FSXFREXEKP4TDCV7BHV7T",
		"private_ip": "fdaa:2:3d6:a7b:b1ce:3371:6336:2",
		"config": {
		"env": { "FLY_PROCESS_GROUP": "app", "PRIMARY_REGION": "iad" },
		"init": {},
		"guest": { "cpu_kind": "shared", "cpus": 4, "memory_mb": 2048 },
		"metadata": {
			"fly_flyctl_version": "0.2.23",
			"fly_platform_version": "v2",
			"fly_process_group": "app",
			"fly_release_id": "51azqDBYY72aRcNP8bOKKgDaP",
			"fly_release_version": "11"
		},
		"mounts": [
			{
			"encrypted": true,
			"path": "/data",
			"size_gb": 12,
			"volume": "vol_v30lq03973nkxx1v",
			"name": "idev_tigris_ecache_data"
			}
		],
		"services": [
			{
			"internal_port": 6379,
			"autostop": false,
			"autostart": true,
			"force_instance_key": null
			}
		],
		"checks": {
			"kvrocks": {
			"port": 6379,
			"type": "tcp",
			"interval": "15s",
			"timeout": "5s",
			"grace_period": "30s"
			}
		},
		"image": "registry.fly.io/idev-tigris-ecache:deployment-01HSMPM95GC25XTPJDZS9RBZT9",
		"restart": {}
		},
		"image_ref": {
		"registry": "registry.fly.io",
		"repository": "idev-tigris-ecache",
		"tag": "deployment-01HSMPM95GC25XTPJDZS9RBZT9",
		"digest": "sha256:b05aa81491ec599c3e52835772d2ac5f1ddbdfbc073b70f2b430199f3a59faa5",
		"labels": {
			"org.opencontainers.image.created": "2023-07-21T23:34:15.259Z",
			"org.opencontainers.image.description": "Kvrocks is a distributed key value NoSQL database that uses RocksDB as storage engine and is compatible with Redis protocol.",
			"org.opencontainers.image.licenses": "Apache-2.0",
			"org.opencontainers.image.revision": "93d2f7e775f155dda9d388a86b89b6684fba1c9a",
			"org.opencontainers.image.source": "https://github.com/tigrisdata/kvrocks",
			"org.opencontainers.image.title": "kvrocks",
			"org.opencontainers.image.url": "https://github.com/tigrisdata/kvrocks",
			"org.opencontainers.image.version": "nightly"
		}
		},
		"created_at": "2024-04-09T03:53:07Z",
		"updated_at": "2024-04-09T03:53:22Z",
		"events": [
		{
			"id": "01HV0FTBYPS847D5M16B2KRNBY",
			"type": "start",
			"status": "started",
			"source": "flyd",
			"timestamp": 1712634802134
		},
		{
			"id": "01HV0FSXJMX03SR1XQ44XTYD2V",
			"type": "launch",
			"status": "created",
			"source": "user",
			"timestamp": 1712634787412
		}
		],
		"checks": [
		{
			"name": "kvrocks",
			"status": "passing",
			"output": "Success",
			"updated_at": "2024-04-09T03:53:26.817Z"
		}
		]
	},
	{
		"id": "d89d920a715738",
		"name": "patient-haze-9838",
		"state": "started",
		"region": "ord",
		"instance_id": "01HV0FSXJTSXNPB8CNYTKX7K20",
		"private_ip": "fdaa:2:3d6:a7b:255:4195:b862:2",
		"config": {
		"env": { "FLY_PROCESS_GROUP": "app", "PRIMARY_REGION": "iad" },
		"init": {},
		"guest": { "cpu_kind": "shared", "cpus": 4, "memory_mb": 2048 },
		"metadata": {
			"fly_flyctl_version": "0.2.23",
			"fly_platform_version": "v2",
			"fly_process_group": "app",
			"fly_release_id": "51azqDBYY72aRcNP8bOKKgDaP",
			"fly_release_version": "11"
		},
		"mounts": [
			{
			"encrypted": true,
			"path": "/data",
			"size_gb": 12,
			"volume": "vol_r6g30gxmlqzyklpv",
			"name": "idev_tigris_ecache_data"
			}
		],
		"services": [
			{
			"internal_port": 6379,
			"autostop": false,
			"autostart": true,
			"force_instance_key": null
			}
		],
		"checks": {
			"kvrocks": {
			"port": 6379,
			"type": "tcp",
			"interval": "15s",
			"timeout": "5s",
			"grace_period": "30s"
			}
		},
		"image": "registry.fly.io/idev-tigris-ecache:deployment-01HSMPM95GC25XTPJDZS9RBZT9",
		"restart": {}
		},
		"image_ref": {
		"registry": "registry.fly.io",
		"repository": "idev-tigris-ecache",
		"tag": "deployment-01HSMPM95GC25XTPJDZS9RBZT9",
		"digest": "sha256:b05aa81491ec599c3e52835772d2ac5f1ddbdfbc073b70f2b430199f3a59faa5",
		"labels": {
			"org.opencontainers.image.created": "2023-07-21T23:34:15.259Z",
			"org.opencontainers.image.description": "Kvrocks is a distributed key value NoSQL database that uses RocksDB as storage engine and is compatible with Redis protocol.",
			"org.opencontainers.image.licenses": "Apache-2.0",
			"org.opencontainers.image.revision": "93d2f7e775f155dda9d388a86b89b6684fba1c9a",
			"org.opencontainers.image.source": "https://github.com/tigrisdata/kvrocks",
			"org.opencontainers.image.title": "kvrocks",
			"org.opencontainers.image.url": "https://github.com/tigrisdata/kvrocks",
			"org.opencontainers.image.version": "nightly"
		}
		},
		"created_at": "2024-04-09T03:53:07Z",
		"updated_at": "2024-04-09T03:53:25Z",
		"events": [
		{
			"id": "01HV0FTESB37SQSDBGSQWS6D7D",
			"type": "start",
			"status": "started",
			"source": "flyd",
			"timestamp": 1712634805035
		},
		{
			"id": "01HV0FSXR6TGF7MRHJ5MWFNGFQ",
			"type": "launch",
			"status": "created",
			"source": "user",
			"timestamp": 1712634787590
		}
		],
		"checks": [
		{
			"name": "kvrocks",
			"status": "passing",
			"output": "Success",
			"updated_at": "2024-04-09T03:53:34.174Z"
		}
		]
	},
	{
		"id": "d8d7667fee3768",
		"name": "ancient-morning-2081",
		"state": "started",
		"region": "lax",
		"instance_id": "01HV0FGRXKSGHP5CVMEMR2VSN3",
		"private_ip": "fdaa:2:3d6:a7b:1e2:5e1d:3c5d:2",
		"config": {
		"env": { "FLY_PROCESS_GROUP": "app", "PRIMARY_REGION": "iad" },
		"init": {},
		"guest": { "cpu_kind": "shared", "cpus": 4, "memory_mb": 2048 },
		"metadata": {
			"fly_flyctl_version": "0.2.23",
			"fly_platform_version": "v2",
			"fly_process_group": "app",
			"fly_release_id": "51azqDBYY72aRcNP8bOKKgDaP",
			"fly_release_version": "11"
		},
		"mounts": [
			{
			"encrypted": true,
			"path": "/data",
			"size_gb": 12,
			"volume": "vol_4mn8zn6ejzjld2zr",
			"name": "idev_tigris_ecache_data"
			}
		],
		"services": [
			{
			"internal_port": 6379,
			"autostop": false,
			"autostart": true,
			"force_instance_key": null
			}
		],
		"checks": {
			"kvrocks": {
			"port": 6379,
			"type": "tcp",
			"interval": "15s",
			"timeout": "5s",
			"grace_period": "30s"
			}
		},
		"image": "registry.fly.io/idev-tigris-ecache:deployment-01HSMPM95GC25XTPJDZS9RBZT9",
		"restart": {}
		},
		"image_ref": {
		"registry": "registry.fly.io",
		"repository": "idev-tigris-ecache",
		"tag": "deployment-01HSMPM95GC25XTPJDZS9RBZT9",
		"digest": "sha256:b05aa81491ec599c3e52835772d2ac5f1ddbdfbc073b70f2b430199f3a59faa5",
		"labels": {
			"org.opencontainers.image.created": "2023-07-21T23:34:15.259Z",
			"org.opencontainers.image.description": "Kvrocks is a distributed key value NoSQL database that uses RocksDB as storage engine and is compatible with Redis protocol.",
			"org.opencontainers.image.licenses": "Apache-2.0",
			"org.opencontainers.image.revision": "93d2f7e775f155dda9d388a86b89b6684fba1c9a",
			"org.opencontainers.image.source": "https://github.com/tigrisdata/kvrocks",
			"org.opencontainers.image.title": "kvrocks",
			"org.opencontainers.image.url": "https://github.com/tigrisdata/kvrocks",
			"org.opencontainers.image.version": "nightly"
		}
		},
		"created_at": "2024-04-09T03:48:07Z",
		"updated_at": "2024-04-09T03:48:25Z",
		"events": [
		{
			"id": "01HV0FHAMWFPY2WJKSFNA4FER8",
			"type": "start",
			"status": "started",
			"source": "flyd",
			"timestamp": 1712634505884
		},
		{
			"id": "01HV0FGRZ73AT82SKGGV05TESR",
			"type": "launch",
			"status": "created",
			"source": "user",
			"timestamp": 1712634487783
		}
		],
		"checks": [
		{
			"name": "kvrocks",
			"status": "passing",
			"output": "Success",
			"updated_at": "2024-04-09T03:48:35.093Z"
		}
		]
	},
	{
		"id": "3d8d7947f55738",
		"name": "still-silence-9494",
		"state": "started",
		"region": "lax",
		"instance_id": "01HV0FGRXX902H8G0W4CE5BWTP",
		"private_ip": "fdaa:2:3d6:a7b:21a:9b96:6b9b:2",
		"config": {
		"env": { "FLY_PROCESS_GROUP": "app", "PRIMARY_REGION": "iad" },
		"init": {},
		"guest": { "cpu_kind": "shared", "cpus": 4, "memory_mb": 2048 },
		"metadata": {
			"fly_flyctl_version": "0.2.23",
			"fly_platform_version": "v2",
			"fly_process_group": "app",
			"fly_release_id": "51azqDBYY72aRcNP8bOKKgDaP",
			"fly_release_version": "11"
		},
		"mounts": [
			{
			"encrypted": true,
			"path": "/data",
			"size_gb": 12,
			"volume": "vol_vjpekp32zq9q6x9v",
			"name": "idev_tigris_ecache_data"
			}
		],
		"services": [
			{
			"internal_port": 6379,
			"autostop": false,
			"autostart": true,
			"force_instance_key": null
			}
		],
		"checks": {
			"kvrocks": {
			"port": 6379,
			"type": "tcp",
			"interval": "15s",
			"timeout": "5s",
			"grace_period": "30s"
			}
		},
		"image": "registry.fly.io/idev-tigris-ecache:deployment-01HSMPM95GC25XTPJDZS9RBZT9",
		"restart": {}
		},
		"image_ref": {
		"registry": "registry.fly.io",
		"repository": "idev-tigris-ecache",
		"tag": "deployment-01HSMPM95GC25XTPJDZS9RBZT9",
		"digest": "sha256:b05aa81491ec599c3e52835772d2ac5f1ddbdfbc073b70f2b430199f3a59faa5",
		"labels": {
			"org.opencontainers.image.created": "2023-07-21T23:34:15.259Z",
			"org.opencontainers.image.description": "Kvrocks is a distributed key value NoSQL database that uses RocksDB as storage engine and is compatible with Redis protocol.",
			"org.opencontainers.image.licenses": "Apache-2.0",
			"org.opencontainers.image.revision": "93d2f7e775f155dda9d388a86b89b6684fba1c9a",
			"org.opencontainers.image.source": "https://github.com/tigrisdata/kvrocks",
			"org.opencontainers.image.title": "kvrocks",
			"org.opencontainers.image.url": "https://github.com/tigrisdata/kvrocks",
			"org.opencontainers.image.version": "nightly"
		}
		},
		"created_at": "2024-04-09T03:48:07Z",
		"updated_at": "2024-04-09T03:48:24Z",
		"events": [
		{
			"id": "01HV0FH8VVD8WJEMSE9R97RGXD",
			"type": "start",
			"status": "started",
			"source": "flyd",
			"timestamp": 1712634504059
		},
		{
			"id": "01HV0FGRZ9052JPK80YS7BQ1D9",
			"type": "launch",
			"status": "created",
			"source": "user",
			"timestamp": 1712634487785
		}
		],
		"checks": [
		{
			"name": "kvrocks",
			"status": "passing",
			"output": "Success",
			"updated_at": "2024-04-09T03:48:40.817Z"
		}
		]
	},
	{
		"id": "32875630b54438",
		"name": "red-sunset-4877",
		"state": "started",
		"region": "sjc",
		"instance_id": "01HV0G4DZZFNGX8A8AJB405CQK",
		"private_ip": "fdaa:2:3d6:a7b:247:f42a:8abf:2",
		"config": {
		"env": { "FLY_PROCESS_GROUP": "app", "PRIMARY_REGION": "iad" },
		"init": {},
		"guest": { "cpu_kind": "shared", "cpus": 4, "memory_mb": 2048 },
		"metadata": {
			"fly_flyctl_version": "0.2.23",
			"fly_platform_version": "v2",
			"fly_process_group": "app",
			"fly_release_id": "51azqDBYY72aRcNP8bOKKgDaP",
			"fly_release_version": "11"
		},
		"mounts": [
			{
			"encrypted": true,
			"path": "/data",
			"size_gb": 12,
			"volume": "vol_49k3088xz2zejxov",
			"name": "idev_tigris_ecache_data"
			}
		],
		"services": [
			{
			"internal_port": 6379,
			"autostop": false,
			"autostart": true,
			"force_instance_key": null
			}
		],
		"checks": {
			"kvrocks": {
			"port": 6379,
			"type": "tcp",
			"interval": "15s",
			"timeout": "5s",
			"grace_period": "30s"
			}
		},
		"image": "registry.fly.io/idev-tigris-ecache:deployment-01HSMPM95GC25XTPJDZS9RBZT9",
		"restart": {}
		},
		"image_ref": {
		"registry": "registry.fly.io",
		"repository": "idev-tigris-ecache",
		"tag": "deployment-01HSMPM95GC25XTPJDZS9RBZT9",
		"digest": "sha256:b05aa81491ec599c3e52835772d2ac5f1ddbdfbc073b70f2b430199f3a59faa5",
		"labels": {
			"org.opencontainers.image.created": "2023-07-21T23:34:15.259Z",
			"org.opencontainers.image.description": "Kvrocks is a distributed key value NoSQL database that uses RocksDB as storage engine and is compatible with Redis protocol.",
			"org.opencontainers.image.licenses": "Apache-2.0",
			"org.opencontainers.image.revision": "93d2f7e775f155dda9d388a86b89b6684fba1c9a",
			"org.opencontainers.image.source": "https://github.com/tigrisdata/kvrocks",
			"org.opencontainers.image.title": "kvrocks",
			"org.opencontainers.image.url": "https://github.com/tigrisdata/kvrocks",
			"org.opencontainers.image.version": "nightly"
		}
		},
		"created_at": "2024-04-09T03:58:51Z",
		"updated_at": "2024-04-09T03:59:07Z",
		"events": [
		{
			"id": "01HV0G4X6K88CYM5DDC6F1GWS9",
			"type": "start",
			"status": "started",
			"source": "flyd",
			"timestamp": 1712635147475
		},
		{
			"id": "01HV0G4E1JFD60PB9H3AKMSF4H",
			"type": "launch",
			"status": "created",
			"source": "user",
			"timestamp": 1712635131954
		}
		],
		"checks": [
		{
			"name": "kvrocks",
			"status": "passing",
			"output": "Success",
			"updated_at": "2024-04-09T03:59:13.789Z"
		}
		]
	},
	{
		"id": "2874d34a006098",
		"name": "dark-wildflower-599",
		"state": "started",
		"region": "sjc",
		"instance_id": "01HV0G4DW40ERSZ0Q0QDD6ZJCZ",
		"private_ip": "fdaa:2:3d6:a7b:167:eb85:7c48:2",
		"config": {
		"env": { "FLY_PROCESS_GROUP": "app", "PRIMARY_REGION": "iad" },
		"init": {},
		"guest": { "cpu_kind": "shared", "cpus": 4, "memory_mb": 2048 },
		"metadata": {
			"fly_flyctl_version": "0.2.23",
			"fly_platform_version": "v2",
			"fly_process_group": "app",
			"fly_release_id": "51azqDBYY72aRcNP8bOKKgDaP",
			"fly_release_version": "11"
		},
		"mounts": [
			{
			"encrypted": true,
			"path": "/data",
			"size_gb": 12,
			"volume": "vol_vzk5nee3ge893pzv",
			"name": "idev_tigris_ecache_data"
			}
		],
		"services": [
			{
			"internal_port": 6379,
			"autostop": false,
			"autostart": true,
			"force_instance_key": null
			}
		],
		"checks": {
			"kvrocks": {
			"port": 6379,
			"type": "tcp",
			"interval": "15s",
			"timeout": "5s",
			"grace_period": "30s"
			}
		},
		"image": "registry.fly.io/idev-tigris-ecache:deployment-01HSMPM95GC25XTPJDZS9RBZT9",
		"restart": {}
		},
		"image_ref": {
		"registry": "registry.fly.io",
		"repository": "idev-tigris-ecache",
		"tag": "deployment-01HSMPM95GC25XTPJDZS9RBZT9",
		"digest": "sha256:b05aa81491ec599c3e52835772d2ac5f1ddbdfbc073b70f2b430199f3a59faa5",
		"labels": {
			"org.opencontainers.image.created": "2023-07-21T23:34:15.259Z",
			"org.opencontainers.image.description": "Kvrocks is a distributed key value NoSQL database that uses RocksDB as storage engine and is compatible with Redis protocol.",
			"org.opencontainers.image.licenses": "Apache-2.0",
			"org.opencontainers.image.revision": "93d2f7e775f155dda9d388a86b89b6684fba1c9a",
			"org.opencontainers.image.source": "https://github.com/tigrisdata/kvrocks",
			"org.opencontainers.image.title": "kvrocks",
			"org.opencontainers.image.url": "https://github.com/tigrisdata/kvrocks",
			"org.opencontainers.image.version": "nightly"
		}
		},
		"created_at": "2024-04-09T03:58:51Z",
		"updated_at": "2024-04-09T03:59:06Z",
		"events": [
		{
			"id": "01HV0G4W9X6FHCJZ1Q7N6QKR8H",
			"type": "start",
			"status": "started",
			"source": "flyd",
			"timestamp": 1712635146557
		},
		{
			"id": "01HV0G4DWYNYMC2BBP8V6WF5TD",
			"type": "launch",
			"status": "created",
			"source": "user",
			"timestamp": 1712635131806
		}
		],
		"checks": [
		{
			"name": "kvrocks",
			"status": "passing",
			"output": "Success",
			"updated_at": "2024-04-09T03:59:20.373Z"
		}
		]
	}
	]
	`)
	object := []Machine{}
	err := json.Unmarshal(test, &object)
	assert.Equal(t, err, nil)
}
