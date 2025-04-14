package toncenterzp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	// DefaultBaseURL is the default base URL for the TON API
	DefaultBaseURL = "https://ton.getblock.io/mainnet/"
	
	// DefaultTimeout is the default timeout for HTTP requests
	DefaultTimeout = 30 * time.Second
)

// ErrorResponse represents an error response from the API
type ErrorResponse struct {
	OK     bool   `json:"ok"`
	Error  string `json:"error,omitempty"`
	Code   int    `json:"code,omitempty"`
	Status int    `json:"-"`
}

// Client represents a TON API client
type Client struct {
	BaseURL    string
	APIKey     string
	HTTPClient *http.Client
}

// NewClient creates a new TON API client with the given API key
func NewClient(apiKey string) *Client {
	return &Client{
		BaseURL:    DefaultBaseURL,
		APIKey:     apiKey,
		HTTPClient: &http.Client{Timeout: DefaultTimeout},
	}
}

// NewClientWithOptions creates a new TON API client with custom options
func NewClientWithOptions(apiKey, baseURL string, timeout time.Duration) *Client {
	return &Client{
		BaseURL:    baseURL,
		APIKey:     apiKey,
		HTTPClient: &http.Client{Timeout: timeout},
	}
}

// doRequest performs an HTTP request to the TON API
func (c *Client) doRequest(method, endpoint string, body interface{}) ([]byte, error) {
	var reqBody io.Reader
	
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("error marshaling request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}
	
	url := c.BaseURL + endpoint
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	
	req.Header.Set("x-api-key", c.APIKey)
	req.Header.Set("Content-Type", "application/json")
	
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()
	
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	
	if resp.StatusCode != http.StatusOK {
		var errResp ErrorResponse
		if err := json.Unmarshal(respBody, &errResp); err != nil {
			return nil, fmt.Errorf("error response with status code %d: %s", resp.StatusCode, string(respBody))
		}
		errResp.Status = resp.StatusCode
		return nil, fmt.Errorf("API error: %s (code: %d, status: %d)", errResp.Error, errResp.Code, errResp.Status)
	}
	
	return respBody, nil
}
