package toncenterzp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// SendBocRequest represents the request for the /sendBoc endpoint
type SendBocRequest struct {
	Boc string `json:"boc"`
}

// SendBocResponse represents the response from the /sendBoc endpoint
type SendBocResponse struct {
	OK     bool `json:"ok"`
	Result struct {
		Status int    `json:"status"`
		Hash   string `json:"hash"`
	} `json:"result"`
}

// SendBoc sends a bag of cells to the TON network
func (c *Client) SendBoc(req SendBocRequest) (*SendBocResponse, error) {
	endpoint := "/sendBoc"
	
	respBody, err := c.doRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}
	
	var response SendBocResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w (code: %d)", err, ErrInvalidResponse)
	}
	
	if !response.OK {
		return nil, fmt.Errorf("API returned non-OK status (code: %d)", ErrAPIError)
	}
	
	return &response, nil
}

// SendBocReturnHashRequest represents the request for the /sendBocReturnHash endpoint
type SendBocReturnHashRequest struct {
	Boc string `json:"boc"`
}

// SendBocReturnHashResponse represents the response from the /sendBocReturnHash endpoint
type SendBocReturnHashResponse struct {
	OK     bool   `json:"ok"`
	Result string `json:"result"`
}

// SendBocReturnHash sends a bag of cells to the TON network and returns the hash
func (c *Client) SendBocReturnHash(req SendBocReturnHashRequest) (*SendBocReturnHashResponse, error) {
	endpoint := "/sendBocReturnHash"
	
	respBody, err := c.doRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}
	
	var response SendBocReturnHashResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w (code: %d)", err, ErrInvalidResponse)
	}
	
	if !response.OK {
		return nil, fmt.Errorf("API returned non-OK status (code: %d)", ErrAPIError)
	}
	
	return &response, nil
}

// SendQueryRequest represents the request for the /sendQuery endpoint
type SendQueryRequest struct {
	Address string `json:"address"`
	Body    string `json:"body"`
	Init    string `json:"init,omitempty"`
}

// SendQueryResponse represents the response from the /sendQuery endpoint
type SendQueryResponse struct {
	OK     bool `json:"ok"`
	Result struct {
		Status int    `json:"status"`
		Hash   string `json:"hash"`
	} `json:"result"`
}

// SendQuery sends a query to the TON network
func (c *Client) SendQuery(req SendQueryRequest) (*SendQueryResponse, error) {
	endpoint := "/sendQuery"
	
	respBody, err := c.doRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}
	
	var response SendQueryResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w (code: %d)", err, ErrInvalidResponse)
	}
	
	if !response.OK {
		return nil, fmt.Errorf("API returned non-OK status (code: %d)", ErrAPIError)
	}
	
	return &response, nil
}

// ShardsResponse represents the response from the /shards endpoint
type ShardsResponse struct {
	OK     bool `json:"ok"`
	Result struct {
		Shards []struct {
			Workchain int    `json:"workchain"`
			Shard     string `json:"shard"`
			SeqNo     int    `json:"seqno"`
			RootHash  string `json:"root_hash"`
			FileHash  string `json:"file_hash"`
		} `json:"shards"`
	} `json:"result"`
}

// Shards gets the list of shards
func (c *Client) Shards(seqNo int) (*ShardsResponse, error) {
	endpoint := fmt.Sprintf("/shards?seqno=%d", seqNo)
	
	respBody, err := c.doRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	
	var response ShardsResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w (code: %d)", err, ErrInvalidResponse)
	}
	
	if !response.OK {
		return nil, fmt.Errorf("API returned non-OK status (code: %d)", ErrAPIError)
	}
	
	return &response, nil
}

// UnpackAddressResponse represents the response from the /unpackAddress endpoint
type UnpackAddressResponse struct {
	OK     bool `json:"ok"`
	Result struct {
		RawForm   string `json:"raw_form"`
		TestOnly  bool   `json:"test_only"`
		Bounceable bool   `json:"bounceable"`
		WorkChain int    `json:"workchain"`
		Hash      string `json:"hash"`
	} `json:"result"`
}

// UnpackAddress unpacks a TON address
func (c *Client) UnpackAddress(address string) (*UnpackAddressResponse, error) {
	endpoint := fmt.Sprintf("/unpackAddress?address=%s", address)
	
	respBody, err := c.doRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	
	var response UnpackAddressResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w (code: %d)", err, ErrInvalidResponse)
	}
	
	if !response.OK {
		return nil, fmt.Errorf("API returned non-OK status (code: %d)", ErrAPIError)
	}
	
	return &response, nil
}
