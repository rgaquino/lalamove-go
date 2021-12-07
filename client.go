package lalamove

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
	"strings"
	"time"
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
	}
	if strings.TrimSpace(c.apiKey) == "" || strings.TrimSpace(c.secret) == "" {
		return nil, errCredentialsMissing
	}
	if strings.TrimSpace(c.baseURL) == "" {
		return nil, errBaseURLMissing
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

// WithSecret configures a Lalamove API client with a secret
func WithSecret(secret string) ClientOption {
	return func(c *Client) error {
		c.secret = secret
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

func (c *Client) get(ctx context.Context, city CityCode, path string, apiReq interface{}, apiResp interface{}) error {
	req, err := c.createRequest(city, http.MethodGet, path, apiReq)
	if err != nil {
		return err
	}
	return c.do(ctx, req, apiResp)
}

func (c *Client) post(ctx context.Context, city CityCode, path string, apiReq interface{}, apiResp interface{}) error {
	req, err := c.createRequest(city, http.MethodPost, path, apiReq)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	return c.do(ctx, req, apiResp)
}

func (c *Client) put(ctx context.Context, city CityCode, path string, apiReq interface{}, apiResp interface{}) error {
	req, err := c.createRequest(city, http.MethodPut, path, apiReq)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	return c.do(ctx, req, apiResp)
}

func (c *Client) createRequest(city CityCode, method, path string, apiReq interface{}) (*http.Request, error) {
	body, bodyBytes, err := marshalRequest(apiReq)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, c.baseURL+path, body)
	if err != nil {
		return nil, err
	}
	auth := c.generateAuth(method, path, bodyBytes)
	req.Header.Set("Authorization", auth)
	req.Header.Set("X-Request-ID", uuid.NewString())
	req.Header.Set("X-LLM-Country", string(city.GetLLMCountry()))
	return req, nil
}

func (c *Client) do(ctx context.Context, req *http.Request, apiResp interface{}) error {
	client := c.httpClient
	if client == nil {
		client = http.DefaultClient
	}
	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return decodeResponse(resp, apiResp)
}

func (c *Client) generateAuth(method, path string, body []byte) string {
	now := time.Now().UnixNano() / int64(time.Millisecond)
	rawSignature := fmt.Sprintf("%d\r\n%s\r\n%s\r\n\r\n%s", now, method, path, string(body))
	mac := hmac.New(sha256.New, []byte(c.secret))
	mac.Write([]byte(rawSignature))
	signature := hex.EncodeToString(mac.Sum(nil))
	return fmt.Sprintf("hmac %s:%d:%s", c.apiKey, now, signature)
}

func marshalRequest(apiReq interface{}) (io.Reader, []byte, error) {
	if apiReq == nil {
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
		return errTooManyRequests
	} else if resp.StatusCode == 401 {
		return errUnauthorized
	} else if resp.StatusCode == 402 || resp.StatusCode == 409 {
		errResp := &ErrorResponse{}
		if err := json.NewDecoder(resp.Body).Decode(errResp); err != nil {
			return err
		}
		return wrapAPIError(errResp)
	} else if resp.StatusCode >= 400 {
		return errUnknownError
	}
	if apiResp == nil {
		return nil
	}
	return json.NewDecoder(resp.Body).Decode(apiResp)
}
