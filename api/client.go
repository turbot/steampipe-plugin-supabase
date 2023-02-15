package api

import (
	"context"
	"net/http"
	"time"
)

// import "context"

type Client struct {
	AccessToken *string
	ApiKey      *string
	HTTPClient  *http.Client
}

const (
	BaseUrl = "https://api.supabase.com/"
)

func CreateClient(ctx context.Context, config ClientConfig) (*Client, error) {
	return &Client{
		AccessToken: &config.AccessToken,
		ApiKey:      &config.ApiKey,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}, nil
}
