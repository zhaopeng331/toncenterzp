package toncenterzp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Error codes
const (
	ErrInvalidParams    = 1001
	ErrNetworkError     = 1002
	ErrInvalidResponse  = 1003
	ErrAPIError         = 1004
	ErrInvalidAddress   = 1005
	ErrInvalidSignature = 1006
	ErrInvalidBlock     = 1007
	ErrInvalidMethod    = 1008
)

// DetectAddressResponse represents the response from the /detectAddress endpoint
type DetectAddressResponse struct {
	OK     bool `json:"ok"`
	Result struct {
		Bounceable struct {
			B64    string `json:"b64"`
			B64Url string `json:"b64url"`
		} `json:"bounceable"`
		GivenType      string `json:"given_type"`
		NonBounceable  struct {
			B64    string `json:"b64"`
			B64Url string `json:"b64url"`
		} `json:"non_bounceable"`
		RawForm   string `json:"raw_form"`
		TestOnly  bool   `json:"test_only"`
		WorkChain int    `json:"workchain"`
	} `json:"result"`
}

// DetectAddress detects the type of a TON address
func (c *Client) DetectAddress(address string) (*DetectAddressResponse, error) {
	endpoint := fmt.Sprintf("/detectAddress?address=%s", address)
	
	respBody, err := c.doRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	
	var response DetectAddressResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w (code: %d)", err, ErrInvalidResponse)
	}
	
	if !response.OK {
		return nil, fmt.Errorf("API returned non-OK status (code: %d)", ErrAPIError)
	}
	
	return &response, nil
}

// EstimateFeeRequest represents the request for the /estimateFee endpoint
type EstimateFeeRequest struct {
	Address      string `json:"address"`
	Body         string `json:"body,omitempty"`
	InitCode     string `json:"init_code,omitempty"`
	InitData     string `json:"init_data,omitempty"`
	IgnoreChksig bool   `json:"ignore_chksig,omitempty"`
}

// EstimateFeeResponse represents the response from the /estimateFee endpoint
type EstimateFeeResponse struct {
	OK     bool `json:"ok"`
	Result struct {
		DestFee   string `json:"dest_fee"`
		FwdFee    string `json:"fwd_fee"`
		GasFee    string `json:"gas_fee"`
		InFwdFee  string `json:"in_fwd_fee"`
		Source    struct {
			Address string `json:"address"`
			WC      int    `json:"wc"`
		} `json:"source"`
		StorageFee string `json:"storage_fee"`
	} `json:"result"`
}

// EstimateFee estimates the fee for a transaction
func (c *Client) EstimateFee(req EstimateFeeRequest) (*EstimateFeeResponse, error) {
	endpoint := "/estimateFee"
	
	respBody, err := c.doRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}
	
	var response EstimateFeeResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w (code: %d)", err, ErrInvalidResponse)
	}
	
	if !response.OK {
		return nil, fmt.Errorf("API returned non-OK status (code: %d)", ErrAPIError)
	}
	
	return &response, nil
}

// GetAddressBalanceResponse represents the response from the /getAddressBalance endpoint
type GetAddressBalanceResponse struct {
	OK     bool   `json:"ok"`
	Result string `json:"result"`
}

// GetAddressBalance gets the balance of a TON address
func (c *Client) GetAddressBalance(address string) (*GetAddressBalanceResponse, error) {
	endpoint := fmt.Sprintf("/getAddressBalance?address=%s", address)
	
	respBody, err := c.doRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	
	var response GetAddressBalanceResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w (code: %d)", err, ErrInvalidResponse)
	}
	
	if !response.OK {
		return nil, fmt.Errorf("API returned non-OK status (code: %d)", ErrAPIError)
	}
	
	return &response, nil
}

// GetAddressInformationResponse represents the response from the /getAddressInformation endpoint
type GetAddressInformationResponse struct {
	OK     bool `json:"ok"`
	Result struct {
		Address         string `json:"address"`
		Balance         string `json:"balance"`
		Code            string `json:"code"`
		Data            string `json:"data"`
		LastTransLT     string `json:"last_trans_lt"`
		LastTransHash   string `json:"last_trans_hash"`
		FrozenHash      string `json:"frozen_hash"`
		SyncUtime       int    `json:"sync_utime"`
		State           string `json:"state"`
		AccountStatus   string `json:"account_status"`
		AccountStorage  string `json:"account_storage"`
		AccountCode     string `json:"account_code"`
		AccountData     string `json:"account_data"`
		ProofOfStateVal string `json:"proof_of_state_val"`
	} `json:"result"`
}

// GetAddressInformation gets detailed information about a TON address
func (c *Client) GetAddressInformation(address string) (*GetAddressInformationResponse, error) {
	endpoint := fmt.Sprintf("/getAddressInformation?address=%s", address)
	
	respBody, err := c.doRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	
	var response GetAddressInformationResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w (code: %d)", err, ErrInvalidResponse)
	}
	
	if !response.OK {
		return nil, fmt.Errorf("API returned non-OK status (code: %d)", ErrAPIError)
	}
	
	return &response, nil
}

// GetAddressStateResponse represents the response from the /getAddressState endpoint
type GetAddressStateResponse struct {
	OK     bool   `json:"ok"`
	Result string `json:"result"`
}

// GetAddressState gets the state of a TON address
func (c *Client) GetAddressState(address string) (*GetAddressStateResponse, error) {
	endpoint := fmt.Sprintf("/getAddressState?address=%s", address)
	
	respBody, err := c.doRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	
	var response GetAddressStateResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w (code: %d)", err, ErrInvalidResponse)
	}
	
	if !response.OK {
		return nil, fmt.Errorf("API returned non-OK status (code: %d)", ErrAPIError)
	}
	
	return &response, nil
}
