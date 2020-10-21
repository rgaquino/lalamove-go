package lalamove

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/twinj/uuid"
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

func (c *Client) get(ctx context.Context, region UNLOCODE, path string, apiReq interface{}, apiResp interface{}) error {
	body, bodyBytes, err := marshalRequest(apiReq)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodGet, c.baseURL+path, body)
	if err != nil {
		return err
	}
	auth := c.generateAuth(http.MethodGet, path, bodyBytes)
	req.Header.Set("Authorization", auth)
	req.Header.Set("X-LLM-Country", region.getLLMCountry())
	req.Header.Set("X-Request-ID", uuid.NewV4().String())

	resp, err := c.do(ctx, req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	return decodeResponse(resp, apiResp)
}

func (c *Client) post(ctx context.Context, region UNLOCODE, path string, apiReq interface{}, apiResp interface{}) error {
	body, bodyBytes, err := marshalRequest(apiReq)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, c.baseURL+path, body)
	if err != nil {
		return err
	}
	auth := c.generateAuth(http.MethodPost, path, bodyBytes)
	req.Header.Set("Authorization", auth)
	req.Header.Set("X-LLM-Country", region.getLLMCountry())
	req.Header.Set("X-Request-ID", uuid.NewV4().String())
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.do(ctx, req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	return decodeResponse(resp, apiResp)
}

func (c *Client) do(ctx context.Context, req *http.Request) (*http.Response, error) {
	client := c.httpClient
	if client == nil {
		client = http.DefaultClient
	}
	return client.Do(req.WithContext(ctx))
}

func (c *Client) generateAuth(method, path string, body []byte) string {
	now := time.Now().Unix()
	rawSignature := fmt.Sprintf("%d\r\n%s\r\n%s\r\n\r\n%s", now, method, path, body)
	mac := hmac.New(sha256.New, []byte(rawSignature))
	mac.Write([]byte(c.secret))
	signature := mac.Sum(nil)
	return fmt.Sprintf("hmac %s:%d:%s", c.apiKey, now, signature)
}

func marshalRequest(apiReq interface{}) (io.Reader, []byte, error) {
	if apiReq != nil {
		return nil, nil, nil
	}
	body, err := json.Marshal(apiReq)
	if err != nil {
		return nil, nil, err
	}
	return bytes.NewBuffer(body), body, nil
}

func decodeResponse(resp *http.Response, apiResp interface{}) error {
	if resp.StatusCode == 429 {
		return apiErrTooManyRequests
	} else if resp.StatusCode >= 400 {
		errResp := &ErrorResponse{}
		if err := json.NewDecoder(resp.Body).Decode(errResp); err != nil {
			return err
		}
		return wrapAPIError(errResp)
	}
	return json.NewDecoder(resp.Body).Decode(apiResp)
}
