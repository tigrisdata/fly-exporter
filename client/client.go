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

package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/tigrisdata/fly-exporter/models"
	"github.com/tigrisdata/fly-exporter/util"
)

const (
	machinesURL = "https://api.machines.dev"
)

type Client struct {
	client *http.Client
	url    string
	token  string
	slug   string
}

// TODO: getters and setters for above

func NewClient() *Client {
	return &Client{
		client: http.DefaultClient,
		url:    util.GetEnv("FLY_EXPORTER_URL", machinesURL),
		token:  util.GetEnv("FLY_API_TOKEN", ""),
		slug:   util.GetEnv("FLY_ORG_SLUG", ""),
	}
}

func (c *Client) GetApps(ctx context.Context) ([]models.Application, error) {
	apps := models.Applications{}

	req, err := c.NewRequest(ctx, http.MethodGet, fmt.Sprintf("%s/v1/apps?org_slug=%s", c.url, c.slug))
	if err != nil {
		return nil, err
	}

	if err := c.ParseDo(req, &apps); err != nil {
		return nil, err
	}

	return apps.Apps, nil
}

func (c *Client) GetMachines(ctx context.Context, app *models.Application) ([]models.Machine, error) {
	machines := []models.Machine{}

	req, err := c.NewRequest(ctx, http.MethodGet, fmt.Sprintf("%s/v1/apps/%s/machines", c.url, app.Name))
	if err != nil {
		return nil, err
	}

	if err := c.ParseDo(req, &machines); err != nil {
		return nil, err
	}

	return machines, nil
}

func (c *Client) NewRequest(ctx context.Context, method string, url string) (*http.Request, error) {
	if c.slug == "" {
		return nil, ErrSlugRequired
	}

	if c.token == "" {
		return nil, ErrTokenRequired
	}

	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+c.token)
	req.Header.Add("Content-Type", "application/json")

	return req, nil
}

func (c *Client) ParseDo(req *http.Request, out any) error {
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(body, &out); err != nil {
		return err
	}

	return nil
}
