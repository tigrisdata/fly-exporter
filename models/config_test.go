package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyCheckHeaderUnmarshal(t *testing.T) {
	test := []byte(`{}`)
	object := CheckHeader{}
	err := json.Unmarshal(test, &object)
	assert.Equal(t, err, nil)
}

// TODO: obtain samples
func TestCheckHeaderUnmarshal(t *testing.T) {
	test := []byte(`{}`)
	object := CheckHeader{}
	err := json.Unmarshal(test, &object)
	assert.Equal(t, err, nil)
}

func TestEmptyCheckConfigUnmarshal(t *testing.T) {
	test := []byte(`{}`)
	object := CheckConfig{}
	err := json.Unmarshal(test, &object)
	assert.Equal(t, err, nil)
}

func TestCheckConfigUnmarshal(t *testing.T) {
	test := []byte(`
	{
        "kvrocks": {
          "port": 6379,
          "type": "tcp",
          "interval": "15s",
          "timeout": "5s",
          "grace_period": "30s"
        }
	}
	`)
	object := CheckConfig{}
	err := json.Unmarshal(test, &object)
	assert.Equal(t, err, nil)
}

func TestEmptyEnvUnmarshal(t *testing.T) {
	test := []byte(`{}`)
	object := map[string]string{}
	err := json.Unmarshal(test, &object)
	assert.Equal(t, err, nil)
}

func TestEnvUnmarshal(t *testing.T) {
	test := []byte(`
	{
		"FLY_PROCESS_GROUP": "app", 
		"PRIMARY_REGION": "iad"
	}
	`)
	object := map[string]string{}
	err := json.Unmarshal(test, &object)
	assert.Equal(t, err, nil)
}

func TestEmptyInitConfigUnmarshal(t *testing.T) {
	test := []byte(`{}`)
	object := InitConfig{}
	err := json.Unmarshal(test, &object)
	assert.Equal(t, err, nil)
}

// TODO: need sample
func TestInitConfigUnmarshal(t *testing.T) {
	test := []byte(`{}`)
	object := InitConfig{}
	err := json.Unmarshal(test, &object)
	assert.Equal(t, err, nil)
}

func TestEmptyGuestConfigUnmarshal(t *testing.T) {
	test := []byte(`{}`)
	object := RestartConfig{}
	err := json.Unmarshal(test, &object)
	assert.Equal(t, err, nil)
}

func TestGuestConfigUnmarshal(t *testing.T) {
	test := []byte(`{ "cpu_kind": "shared", "cpus": 4, "memory_mb": 2048 }`)
	object := GuestConfig{}
	err := json.Unmarshal(test, &object)
	assert.Equal(t, err, nil)
}

func TestEmptyImageUnmarshal(t *testing.T) {
	test := []byte(`""`)
	var object string
	err := json.Unmarshal(test, &object)
	assert.Equal(t, err, nil)
}

func TestImageUnmarshal(t *testing.T) {
	test := []byte(`"registry.fly.io/idev-tigris-ecache:deployment-01HSMPM95GC25XTPJDZS9RBZT9"`)
	var object string
	err := json.Unmarshal(test, &object)
	assert.Equal(t, err, nil)
}

func TestEmptyMetadataConfigUnmarshal(t *testing.T) {
	test := []byte(`{}`)
	object := MetadataConfig{}
	err := json.Unmarshal(test, &object)
	assert.Equal(t, err, nil)
}

func TestMetadataConfigUnmarshal(t *testing.T) {
	test := []byte(`
	{
        "fly_flyctl_version": "0.2.23",
        "fly_platform_version": "v2",
        "fly_process_group": "app",
        "fly_release_id": "51azqDBYY72aRcNP8bOKKgDaP",
        "fly_release_version": "11"
    }
	`)
	object := MetadataConfig{}
	err := json.Unmarshal(test, &object)
	assert.Equal(t, err, nil)
}

func TestEmptyMountConfigUnmarshal(t *testing.T) {
	test := []byte(`{}`)
	object := MountConfig{}
	err := json.Unmarshal(test, &object)
	assert.Equal(t, err, nil)
}

func TestMountConfigUnmarshal(t *testing.T) {
	test := []byte(`
	{
		"encrypted": true,
		"path": "/data",
		"size_gb": 12,
		"volume": "vol_vzk5nee3ge893pzv",
		"name": "idev_tigris_ecache_data"
	}
	`)
	object := MountConfig{}
	err := json.Unmarshal(test, &object)
	assert.Equal(t, err, nil)
}

func TestEmptyRestartConfigUnmarshal(t *testing.T) {
	test := []byte(`{}`)
	object := RestartConfig{}
	err := json.Unmarshal(test, &object)
	assert.Equal(t, err, nil)
}

// TODO: need sample
func TestRestartConfigUnmarshal(t *testing.T) {
	test := []byte(`{}`)
	object := RestartConfig{}
	err := json.Unmarshal(test, &object)
	assert.Equal(t, err, nil)
}

func TestEmptServiceConfigUnmarshal(t *testing.T) {
	test := []byte(`{}`)
	object := ServiceConfig{}
	err := json.Unmarshal(test, &object)
	assert.Equal(t, err, nil)
}

func TestServicesConfigUnmarshal(t *testing.T) {
	test := []byte(`
	{
		"internal_port": 6379,
		"autostop": false,
		"autostart": true,
		"force_instance_key": null
	}
	`)
	object := ServiceConfig{}
	err := json.Unmarshal(test, &object)
	assert.Equal(t, err, nil)
}

func TestConfigUnmarshal(t *testing.T) {
	test := []byte(`
	{
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
	}
	`)
	object := Config{}
	err := json.Unmarshal(test, &object)
	assert.Equal(t, err, nil)
}
