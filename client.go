package lalamove

import (
	"errors"
	"net/http"
)

var errCredentialsMissing = errors.New("API Key credentials missing")

// Client may be used to make requests to the Lalamove APIs
type Client struct {
	httpClient *http.Client
	apiKey     string
	secret     string
	baseURL    string
}

// ClientOption is the type of constructor options for NewClient(...).
type ClientOption func(*Client) error

// NewClient constructs a new Client which can make requests to the Lalamove APIs.
func NewClient(options ...ClientOption) (*Client, error) {
	c := &Client{}
	for _, option := range options {
		err := option(c)
		if err != nil {
			return nil, err
		}
		if c.apiKey == "" {
			return nil, errCredentialsMissing
		}
	}
	return c, nil
}

// WithHTTPClient configures a Lalamove API client with a http.Client to make requests over.
func WithHTTPClient(c *http.Client) ClientOption {
	return func(client *Client) error {
		if c.Transport == nil {
			c.Transport = http.DefaultTransport
		}
		client.httpClient = c
		return nil
	}
}

// WithAPIKey configures a Lalamove API client with an API Key
func WithAPIKey(apiKey string) ClientOption {
	return func(c *Client) error {
		c.apiKey = apiKey
		return nil
	}
}

// WithBaseURL configures a Lalamove API client with a custom base url
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		c.baseURL = baseURL
		return nil
	}
}
