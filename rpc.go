package toncenterzp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JSONRPCRequest represents a JSON-RPC request
type JSONRPCRequest struct {
	JSONRPC string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	ID      int         `json:"id"`
}

// JSONRPCResponse represents a JSON-RPC response
type JSONRPCResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   *JSONRPCError   `json:"error,omitempty"`
	ID      int             `json:"id"`
}

// JSONRPCError represents a JSON-RPC error
type JSONRPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data,omitempty"`
}

// JSONRPC sends a JSON-RPC request to the TON API
func (c *Client) JSONRPC(method string, params interface{}) (*JSONRPCResponse, error) {
	endpoint := "/jsonRPC"
	
	req := JSONRPCRequest{
		JSONRPC: "2.0",
		Method:  method,
		Params:  params,
		ID:      1,
	}
	
	respBody, err := c.doRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}
	
	var response JSONRPCResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w (code: %d)", err, ErrInvalidResponse)
	}
	
	if response.Error != nil {
		return nil, fmt.Errorf("JSON-RPC error: %s (code: %d)", response.Error.Message, response.Error.Code)
	}
	
	return &response, nil
}

// LookupBlockRequest represents the request for the /lookupBlock endpoint
type LookupBlockRequest struct {
	Workchain int    `json:"workchain"`
	Shard     string `json:"shard,omitempty"`
	SeqNo     int    `json:"seqno,omitempty"`
	Lt        string `json:"lt,omitempty"`
	UnixTime  int    `json:"unixtime,omitempty"`
}

// LookupBlockResponse represents the response from the /lookupBlock endpoint
type LookupBlockResponse struct {
	OK     bool    `json:"ok"`
	Result BlockID `json:"result"`
}

// LookupBlock looks up a block by various criteria
func (c *Client) LookupBlock(req LookupBlockRequest) (*LookupBlockResponse, error) {
	endpoint := "/lookupBlock"
	
	respBody, err := c.doRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}
	
	var response LookupBlockResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w (code: %d)", err, ErrInvalidResponse)
	}
	
	if !response.OK {
		return nil, fmt.Errorf("API returned non-OK status (code: %d)", ErrAPIError)
	}
	
	return &response, nil
}

// PackAddressResponse represents the response from the /packAddress endpoint
type PackAddressResponse struct {
	OK     bool `json:"ok"`
	Result struct {
		RawForm string `json:"raw_form"`
	} `json:"result"`
}

// PackAddress packs a TON address
func (c *Client) PackAddress(address string) (*PackAddressResponse, error) {
	endpoint := fmt.Sprintf("/packAddress?address=%s", address)
	
	respBody, err := c.doRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	
	var response PackAddressResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w (code: %d)", err, ErrInvalidResponse)
	}
	
	if !response.OK {
		return nil, fmt.Errorf("API returned non-OK status (code: %d)", ErrAPIError)
	}
	
	return &response, nil
}

// RunGetMethodRequest represents the request for the /runGetMethod endpoint
type RunGetMethodRequest struct {
	Address string        `json:"address"`
	Method  string        `json:"method"`
	Stack   []interface{} `json:"stack,omitempty"`
}

// RunGetMethodResponse represents the response from the /runGetMethod endpoint
type RunGetMethodResponse struct {
	OK     bool `json:"ok"`
	Result struct {
		GasUsed int             `json:"gas_used"`
		Stack   [][]interface{} `json:"stack"`
		ExitCode int            `json:"exit_code"`
	} `json:"result"`
}

// RunGetMethod runs a get method on a TON contract
func (c *Client) RunGetMethod(req RunGetMethodRequest) (*RunGetMethodResponse, error) {
	endpoint := "/runGetMethod"
	
	respBody, err := c.doRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}
	
	var response RunGetMethodResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w (code: %d)", err, ErrInvalidResponse)
	}
	
	if !response.OK {
		return nil, fmt.Errorf("API returned non-OK status (code: %d)", ErrAPIError)
	}
	
	return &response, nil
}
