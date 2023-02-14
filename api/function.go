package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Function struct {
	CreatedAt int64  `json:"created_at,omitempty"`
	Id        string `json:"id,omitempty"`
	ImportMap bool   `json:"import_map,omitempty"`
	Name      string `json:"name,omitempty"`
	ProjectId string `json:"-"`
	Slug      string `json:"slug,omitempty"`
	Status    string `json:"status,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
	VerifyJwt bool   `json:"verify_jwt,omitempty"`
	Version   int    `json:"version,omitempty"`
}

func ListFunctions(ctx context.Context, client *Client, projectID string) ([]*Function, error) {
	apiEndpoint := fmt.Sprintf("/v1/projects/%s/functions", projectID)
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

	var functions []*Function
	if resp.StatusCode != http.StatusNoContent {
		if err = json.NewDecoder(resp.Body).Decode(&functions); err != nil {
			return nil, fmt.Errorf("failed to decode the response: %v", err)
		}
	}

	return functions, nil
}

func GetFunction(ctx context.Context, client *Client, projectID string, slug string) (*Function, error) {
	apiEndpoint := fmt.Sprintf("/v1/projects/%s/functions/%s", projectID, slug)
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

	var function Function
	if resp.StatusCode != http.StatusNoContent {
		if err = json.NewDecoder(resp.Body).Decode(&function); err != nil {
			return nil, fmt.Errorf("failed to decode the response: %v", err)
		}
	}

	return &function, nil
}
