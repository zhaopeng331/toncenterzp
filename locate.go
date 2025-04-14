package toncenterzp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// TryLocateResultTxRequest represents the request for the /tryLocateResultTx endpoint
type TryLocateResultTxRequest struct {
	Source    string `json:"source"`
	Destination string `json:"destination"`
	CreatedLt string `json:"created_lt"`
}

// TryLocateResultTxResponse represents the response from the /tryLocateResultTx endpoint
type TryLocateResultTxResponse struct {
	OK     bool `json:"ok"`
	Result struct {
		Transaction TransactionDetails `json:"transaction"`
	} `json:"result"`
}

// TryLocateResultTx tries to locate a result transaction
func (c *Client) TryLocateResultTx(req TryLocateResultTxRequest) (*TryLocateResultTxResponse, error) {
	endpoint := "/tryLocateResultTx"
	
	respBody, err := c.doRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}
	
	var response TryLocateResultTxResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w (code: %d)", err, ErrInvalidResponse)
	}
	
	if !response.OK {
		return nil, fmt.Errorf("API returned non-OK status (code: %d)", ErrAPIError)
	}
	
	return &response, nil
}

// TryLocateSourceTxRequest represents the request for the /tryLocateSourceTx endpoint
type TryLocateSourceTxRequest struct {
	Source    string `json:"source"`
	Destination string `json:"destination"`
	CreatedLt string `json:"created_lt"`
}

// TryLocateSourceTxResponse represents the response from the /tryLocateSourceTx endpoint
type TryLocateSourceTxResponse struct {
	OK     bool `json:"ok"`
	Result struct {
		Transaction TransactionDetails `json:"transaction"`
	} `json:"result"`
}

// TryLocateSourceTx tries to locate a source transaction
func (c *Client) TryLocateSourceTx(req TryLocateSourceTxRequest) (*TryLocateSourceTxResponse, error) {
	endpoint := "/tryLocateSourceTx"
	
	respBody, err := c.doRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}
	
	var response TryLocateSourceTxResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w (code: %d)", err, ErrInvalidResponse)
	}
	
	if !response.OK {
		return nil, fmt.Errorf("API returned non-OK status (code: %d)", ErrAPIError)
	}
	
	return &response, nil
}

// TryLocateTxRequest represents the request for the /tryLocateTx endpoint
type TryLocateTxRequest struct {
	Hash string `json:"hash"`
}

// TryLocateTxResponse represents the response from the /tryLocateTx endpoint
type TryLocateTxResponse struct {
	OK     bool `json:"ok"`
	Result struct {
		Transaction TransactionDetails `json:"transaction"`
	} `json:"result"`
}

// TryLocateTx tries to locate a transaction by hash
func (c *Client) TryLocateTx(req TryLocateTxRequest) (*TryLocateTxResponse, error) {
	endpoint := "/tryLocateTx"
	
	respBody, err := c.doRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}
	
	var response TryLocateTxResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w (code: %d)", err, ErrInvalidResponse)
	}
	
	if !response.OK {
		return nil, fmt.Errorf("API returned non-OK status (code: %d)", ErrAPIError)
	}
	
	return &response, nil
}
