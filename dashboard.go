package gapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

// DashboardMeta represents Grafana dashboard meta.
type DashboardMeta struct {
	IsStarred bool   `json:"isStarred"`
	Slug      string `json:"slug"`
	Folder    int64  `json:"folderId"`
}

type DashboardSaveResponse struct {
	Slug    string `json:"slug"`
	Id      int64  `json:"id"`
	Uid     string `json:"uid"`
	Status  string `json:"status"`
	Version int64  `json:"version"`
}

type DashboardSearchResponse struct {
	Id          uint     `json:"id"`
	Uid         string   `json:"uid"`
	Title       string   `json:"title"`
	Uri         string   `json:"uri"`
	Url         string   `json:"url"`
	Slug        string   `json:"slug"`
	Type        string   `json:"type"`
	Tags        []string `json:"tags"`
	IsStarred   bool     `json:"isStarred"`
	FolderId    uint     `json:"folderId"`
	FolderUid   string   `json:"folderUid"`
	FolderTitle string   `json:"folderTitle"`
	FolderUrl   string   `json:"folderUrl"`
}

type Dashboard struct {
	Meta      DashboardMeta          `json:"meta"`
	Model     map[string]interface{} `json:"dashboard"`
	Folder    int64                  `json:"folderId"`
	Overwrite bool                   `json:"overwrite"`
}

// SaveDashboard is a deprecated method for saving a Grafana dashboard. Use NewDashboard.
// Deprecated: Use NewDashboard instead.
func (c *Client) SaveDashboard(model map[string]interface{}, overwrite bool) (*DashboardSaveResponse, error) {
	wrapper := map[string]interface{}{
		"dashboard": model,
		"overwrite": overwrite,
	}
	data, err := json.Marshal(wrapper)
	if err != nil {
		return nil, err
	}
	resp, err := c.request("POST", "/api/dashboards/db", nil, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := &DashboardSaveResponse{}
	err = json.Unmarshal(data, &result)
	return result, err
}

// NewDashboard creates a new Grafana dashboard.
func (c *Client) NewDashboard(dashboard Dashboard) (*DashboardSaveResponse, error) {
	data, err := json.Marshal(dashboard)
	if err != nil {
		return nil, err
	}
	resp, err := c.request("POST", "/api/dashboards/db", nil, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := &DashboardSaveResponse{}
	err = json.Unmarshal(data, &result)
	return result, err
}

// Dashboards fetches and returns Grafana dashboards.
func (c *Client) Dashboards() ([]DashboardSearchResponse, error) {
	dashboards := make([]DashboardSearchResponse, 0)
	query := url.Values{}
	// search only dashboards
	query.Add("type", "dash-db")
	resp, err := c.request("GET", "/api/search", query, nil)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dashboards, err
	}

	err = json.Unmarshal(data, &dashboards)
	return dashboards, err
}

// DashboardByUid fetches and returns the dashboard whose UID is passed.
func (c *Client) DashboardByUid(uid string) (*Dashboard, error) {
	return c.dashboard(fmt.Sprintf("/api/dashboards/uid/%s", uid))
}

// Dashboard will be removed.
// Deprecated: Starting from Grafana v5.0. Use DashboardByUid instead.
func (c *Client) Dashboard(slug string) (*Dashboard, error) {
	return c.dashboard(fmt.Sprintf("/api/dashboards/db/%s", slug))
}

// DashboardByUID will be removed.
// Deprecated: Interface typo. Use DashboardByUid instead.
func (c *Client) DashboardByUID(uid string) (*Dashboard, error) {
	return c.dashboard(fmt.Sprintf("/api/dashboards/uid/%s", uid))
}

func (c *Client) dashboard(path string) (*Dashboard, error) {
	resp, err := c.request("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := &Dashboard{}
	err = json.Unmarshal(data, &result)
	result.Folder = result.Meta.Folder

	return result, err
}

// DeleteDashboardByUid deletes the dashboard whose UID it's passed.
func (c *Client) DeleteDashboardByUid(uid string) error {
	return c.deleteDashboard(fmt.Sprintf("/api/dashboards/uid/%s", uid))
}

// DeleteDashboard will be removed.
// Deprecated: Starting from Grafana v5.0. Use DeleteDashboardByUid instead.
func (c *Client) DeleteDashboard(slug string) error {
	return c.deleteDashboard(fmt.Sprintf("/api/dashboards/db/%s", slug))
}

// DeleteDashboardByUID will be removed.
// Deprecated: Interface typo. Use DeleteDashboardByUid instead.
func (c *Client) DeleteDashboardByUID(uid string) error {
	return c.deleteDashboard(fmt.Sprintf("/api/dashboards/uid/%s", uid))
}

func (c *Client) deleteDashboard(path string) error {
	_, err := c.request("DELETE", path, nil, nil)

	return err
}
