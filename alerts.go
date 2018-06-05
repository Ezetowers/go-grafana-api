package gapi

import (
	"fmt"
)

// AlertSummary represents a Grafana alert summary
type AlertSummary struct {
	ID          int64  `json:"id,omitempty"`
	DashboardID int64  `json:"dashboardId"`
	PanelID     int64  `json:"panelId"`
	Name        string `json:"name"`
	State       string `json:"state"`
	URL         string `json:"url"`
}

// Alerts returns an array of AlertSummary objects
func (c *Client) Alerts() ([]AlertSummary, error) {
	result := make([]AlertSummary, 0)

	path := fmt.Sprintf("/api/alerts")
	res, err := c.doRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	if !res.OK() {
		return result, res.Error()
	}

	err = res.BindJSON(&result)
	return result, err
}
