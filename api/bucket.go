package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Bucket struct {
	CreatedAt string `json:"created_at,omitempty"`
	Id        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Owner     string `json:"owner,omitempty"`
	ProjectId string `json:"-"`
	Public    bool   `json:"public,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

func ListBuckets(ctx context.Context, client *Client) ([]*Bucket, error) {
	projectUrl := *client.Url
	storageEndpoint := "/storage/v1/"
	apiEndpoint := "/bucket"

	// Build the request url
	requestURL, err := url.JoinPath(projectUrl, storageEndpoint, apiEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to construct the request URL: %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *client.ApiKey))

	resp, err := client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var buckets []*Bucket
	if resp.StatusCode >= 300 && resp.StatusCode < 500 {
		return nil, fmt.Errorf("the server returned a error: %s", resp.Status)
	}
	if resp.StatusCode != http.StatusNoContent {
		if err = json.NewDecoder(resp.Body).Decode(&buckets); err != nil {
			return nil, fmt.Errorf("failed to decode the response: %v", err)
		}
	}

	return buckets, nil
}

func GetBucket(ctx context.Context, client *Client, id string) (*Bucket, error) {
	projectUrl := *client.Url
	storageEndpoint := "/storage/v1/"
	apiEndpoint := fmt.Sprintf("/bucket/%s", id)

	// Build the request url
	requestURL, err := url.JoinPath(projectUrl, storageEndpoint, apiEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to construct the request URL: %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *client.ApiKey))

	resp, err := client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var bucket Bucket

	if resp.StatusCode >= 300 && resp.StatusCode < 500 {
		return nil, fmt.Errorf("the server returned a error: %s", resp.Status)
	}
	if resp.StatusCode != http.StatusNoContent {
		if err = json.NewDecoder(resp.Body).Decode(&bucket); err != nil {
			return nil, fmt.Errorf("failed to decode the response: %v", err)
		}
	}

	return &bucket, nil
}
