package toncenterzp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetMasterchainBlockSignaturesRequest represents the request for the /getMasterchainBlockSignatures endpoint
type GetMasterchainBlockSignaturesRequest struct {
	SeqNo int `json:"seqno"`
}

// Signature represents a block signature
type Signature struct {
	NodeID      string `json:"node_id"`
	R           string `json:"r"`
	S           string `json:"s"`
	SignatureID int    `json:"signature_id"`
}

// GetMasterchainBlockSignaturesResponse represents the response from the /getMasterchainBlockSignatures endpoint
type GetMasterchainBlockSignaturesResponse struct {
	OK     bool `json:"ok"`
	Result struct {
		Signatures []Signature `json:"signatures"`
	} `json:"result"`
}

// GetMasterchainBlockSignatures gets the signatures of a masterchain block
func (c *Client) GetMasterchainBlockSignatures(req GetMasterchainBlockSignaturesRequest) (*GetMasterchainBlockSignaturesResponse, error) {
	endpoint := "/getMasterchainBlockSignatures"
	
	respBody, err := c.doRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}
	
	var response GetMasterchainBlockSignaturesResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w (code: %d)", err, ErrInvalidResponse)
	}
	
	if !response.OK {
		return nil, fmt.Errorf("API returned non-OK status (code: %d)", ErrAPIError)
	}
	
	return &response, nil
}

// GetMasterchainInfoResponse represents the response from the /getMasterchainInfo endpoint
type GetMasterchainInfoResponse struct {
	OK     bool `json:"ok"`
	Result struct {
		LastBlockID struct {
			Workchain int    `json:"workchain"`
			Shard     string `json:"shard"`
			SeqNo     int    `json:"seqno"`
			RootHash  string `json:"root_hash"`
			FileHash  string `json:"file_hash"`
		} `json:"last"`
		StateRootHash string `json:"state_root_hash"`
		InitSeqNo     int    `json:"init_seq_no"`
	} `json:"result"`
}

// GetMasterchainInfo gets information about the masterchain
func (c *Client) GetMasterchainInfo() (*GetMasterchainInfoResponse, error) {
	endpoint := "/getMasterchainInfo"
	
	respBody, err := c.doRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	
	var response GetMasterchainInfoResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w (code: %d)", err, ErrInvalidResponse)
	}
	
	if !response.OK {
		return nil, fmt.Errorf("API returned non-OK status (code: %d)", ErrAPIError)
	}
	
	return &response, nil
}

// GetShardBlockProofRequest represents the request for the /getShardBlockProof endpoint
type GetShardBlockProofRequest struct {
	Workchain int    `json:"workchain"`
	Shard     string `json:"shard"`
	SeqNo     int    `json:"seqno"`
	FromSeqNo int    `json:"from_seqno,omitempty"`
}

// GetShardBlockProofResponse represents the response from the /getShardBlockProof endpoint
type GetShardBlockProofResponse struct {
	OK     bool `json:"ok"`
	Result struct {
		MasterchainID struct {
			Workchain int    `json:"workchain"`
			Shard     string `json:"shard"`
			SeqNo     int    `json:"seqno"`
			RootHash  string `json:"root_hash"`
			FileHash  string `json:"file_hash"`
		} `json:"masterchain_id"`
		Links []struct {
			ToKeyBlock bool `json:"to_key_block"`
			From       struct {
				Workchain int    `json:"workchain"`
				Shard     string `json:"shard"`
				SeqNo     int    `json:"seqno"`
				RootHash  string `json:"root_hash"`
				FileHash  string `json:"file_hash"`
			} `json:"from"`
			To struct {
				Workchain int    `json:"workchain"`
				Shard     string `json:"shard"`
				SeqNo     int    `json:"seqno"`
				RootHash  string `json:"root_hash"`
				FileHash  string `json:"file_hash"`
			} `json:"to"`
			Dest_Proof string `json:"dest_proof"`
			Proof      string `json:"proof"`
			StateProof string `json:"state_proof"`
		} `json:"links"`
	} `json:"result"`
}

// GetShardBlockProof gets the proof of a shard block
func (c *Client) GetShardBlockProof(req GetShardBlockProofRequest) (*GetShardBlockProofResponse, error) {
	endpoint := "/getShardBlockProof"
	
	respBody, err := c.doRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}
	
	var response GetShardBlockProofResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w (code: %d)", err, ErrInvalidResponse)
	}
	
	if !response.OK {
		return nil, fmt.Errorf("API returned non-OK status (code: %d)", ErrAPIError)
	}
	
	return &response, nil
}

// GetTokenDataResponse represents the response from the /getTokenData endpoint
type GetTokenDataResponse struct {
	OK     bool `json:"ok"`
	Result struct {
		Name     string `json:"name"`
		Symbol   string `json:"symbol"`
		Decimals int    `json:"decimals"`
		Address  string `json:"address"`
	} `json:"result"`
}

// GetTokenData gets data about a token
func (c *Client) GetTokenData(address string) (*GetTokenDataResponse, error) {
	endpoint := fmt.Sprintf("/getTokenData?address=%s", address)
	
	respBody, err := c.doRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	
	var response GetTokenDataResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w (code: %d)", err, ErrInvalidResponse)
	}
	
	if !response.OK {
		return nil, fmt.Errorf("API returned non-OK status (code: %d)", ErrAPIError)
	}
	
	return &response, nil
}
