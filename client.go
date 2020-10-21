package lalamove

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

var (
	errCredentialsMissing = errors.New("API Key credentials missing")
	errBaseURLMissing     = errors.New("base URL missing")
)

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
		if strings.TrimSpace(c.apiKey) == "" {
			return nil, errCredentialsMissing
		}
		if strings.TrimSpace(c.baseURL) == "" {
			return nil, errBaseURLMissing
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

func (c *Client) get(ctx context.Context, path string, body interface{}, apiResp interface{}) error {
	var bodyReader io.Reader
	if body != nil {
		bodyStr, err := json.Marshal(body)
		if err != nil {
			return err
		}
		bytes.NewBuffer(bodyStr)
	}
	req, err := http.NewRequest(http.MethodGet, c.baseURL+path, bodyReader)
	if err != nil {
		return err
	}

	resp, err := c.do(ctx, req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		errResp := &ErrorResponse{}
		if err := json.NewDecoder(resp.Body).Decode(errResp); err != nil {
			return err
		}
		return c.wrapAPIError(errResp)
	}
	return json.NewDecoder(resp.Body).Decode(apiResp)
}

func (c *Client) post(ctx context.Context, path string, body interface{}, apiResp interface{}) error {
	var bodyReader io.Reader
	if body != nil {
		bodyStr, err := json.Marshal(body)
		if err != nil {
			return err
		}
		bytes.NewBuffer(bodyStr)
	}
	req, err := http.NewRequest(http.MethodPost, c.baseURL+path, bodyReader)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.do(ctx, req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		errResp := &ErrorResponse{}
		if err := json.NewDecoder(resp.Body).Decode(errResp); err != nil {
			return err
		}
		return c.wrapAPIError(errResp)
	}
	return json.NewDecoder(resp.Body).Decode(apiResp)
}

func (c *Client) do(ctx context.Context, req *http.Request) (*http.Response, error) {
	client := c.httpClient
	if client == nil {
		client = http.DefaultClient
	}
	return client.Do(req.WithContext(ctx))
}

var (
	apiErrUnknownError       = errors.New("ERR_UNKNOWN")
	apiErrInsufficientCredit = errors.New("ERR_INSUFFICIENT_CREDIT")
	apiErrInvalidCurrency    = errors.New("ERR_INVALID_CURRENCY")
	apiErrPriceMismatch      = errors.New("ERR_PRICE_MISMATCH")
)

func (c *Client) wrapAPIError(errResp *ErrorResponse) error {
	switch errResp.Error {
	case "ERR_INSUFFICIENT_CREDIT":
		return apiErrInsufficientCredit
	case "ERR_INVALID_CURRENCY":
		return apiErrInvalidCurrency
	case "ERR_PRICE_MISMATCH":
		return apiErrPriceMismatch
	}
	return apiErrUnknownError
}
