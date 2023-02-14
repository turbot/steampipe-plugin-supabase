package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Project struct {
	CreatedAt      string          `json:"created_at,omitempty"`
	Database       ProjectDatabase `json:"database,omitempty"`
	Id             string          `json:"id,omitempty"`
	Name           string          `json:"name,omitempty"`
	OrganizationId string          `json:"organization_id,omitempty"`
	Region         string          `json:"region,omitempty"`
}

type ProjectDatabase struct {
	Host    string `json:"host,omitempty"`
	Version string `json:"version,omitempty"`
}

func ListProjects(ctx context.Context, client *Client) ([]Project, error) {
	apiEndpoint := "/v1/projects"
	requestURL, err := url.JoinPath(BaseUrl, apiEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to construct the request URL: %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create requestL: %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *client.AccessToken))

	resp, err := client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var projects []Project
	if resp.StatusCode != http.StatusNoContent {
		if err = json.NewDecoder(resp.Body).Decode(&projects); err != nil {
			return nil, fmt.Errorf("failed to decode the response: %v", err)
		}
	}

	return projects, nil
}
