package toncenterzp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetBlockHeaderRequest represents the request for the /getBlockHeader endpoint
type GetBlockHeaderRequest struct {
	Workchain int    `json:"workchain"`
	Shard     string `json:"shard"`
	SeqNo     int    `json:"seqno"`
}

// GetBlockHeaderResponse represents the response from the /getBlockHeader endpoint
type GetBlockHeaderResponse struct {
	OK     bool `json:"ok"`
	Result struct {
		GlobalID         int    `json:"global_id"`
		Version          int    `json:"version"`
		AfterMerge       bool   `json:"after_merge"`
		BeforeSplit      bool   `json:"before_split"`
		WantMerge        bool   `json:"want_merge"`
		WantSplit        bool   `json:"want_split"`
		ValidatorListHashShort  int    `json:"validator_list_hash_short"`
		CatchainSeqno    int    `json:"catchain_seqno"`
		MinRefMcSeqno    int    `json:"min_ref_mc_seqno"`
		IsKeyBlock       bool   `json:"is_key_block"`
		PrevKeyBlockSeqno int   `json:"prev_key_block_seqno"`
		StartLt          string `json:"start_lt"`
		EndLt            string `json:"end_lt"`
		GenUtime         int    `json:"gen_utime"`
		VertSeqno        int    `json:"vert_seqno"`
		GenSoftwareVersion int  `json:"gen_software_version"`
		GenSoftwareCapabilities string `json:"gen_software_capabilities"`
	} `json:"result"`
}

// GetBlockHeader gets the header of a block
func (c *Client) GetBlockHeader(req GetBlockHeaderRequest) (*GetBlockHeaderResponse, error) {
	endpoint := "/getBlockHeader"
	
	respBody, err := c.doRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}
	
	var response GetBlockHeaderResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w (code: %d)", err, ErrInvalidResponse)
	}
	
	if !response.OK {
		return nil, fmt.Errorf("API returned non-OK status (code: %d)", ErrAPIError)
	}
	
	return &response, nil
}

// GetBlockTransactionsRequest represents the request for the /getBlockTransactions endpoint
type GetBlockTransactionsRequest struct {
	Workchain int    `json:"workchain"`
	Shard     string `json:"shard"`
	SeqNo     int    `json:"seqno"`
	Count     int    `json:"count,omitempty"`
	RootHash  string `json:"root_hash,omitempty"`
	FileHash  string `json:"file_hash,omitempty"`
	AfterLt   string `json:"after_lt,omitempty"`
	AfterHash string `json:"after_hash,omitempty"`
}

// Transaction represents a transaction in the block
type Transaction struct {
	Account   string `json:"account"`
	Hash      string `json:"hash"`
	Lt        string `json:"lt"`
	PrevTrans struct {
		Hash string `json:"hash"`
		Lt   string `json:"lt"`
	} `json:"prev_trans"`
}

// GetBlockTransactionsResponse represents the response from the /getBlockTransactions endpoint
type GetBlockTransactionsResponse struct {
	OK     bool `json:"ok"`
	Result struct {
		Incomplete bool          `json:"incomplete"`
		Transactions []Transaction `json:"transactions"`
	} `json:"result"`
}

// GetBlockTransactions gets the transactions in a block
func (c *Client) GetBlockTransactions(req GetBlockTransactionsRequest) (*GetBlockTransactionsResponse, error) {
	endpoint := "/getBlockTransactions"
	
	respBody, err := c.doRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}
	
	var response GetBlockTransactionsResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w (code: %d)", err, ErrInvalidResponse)
	}
	
	if !response.OK {
		return nil, fmt.Errorf("API returned non-OK status (code: %d)", ErrAPIError)
	}
	
	return &response, nil
}

// GetConsensusBlockRequest represents the request for the /getConsensusBlock endpoint
type GetConsensusBlockRequest struct {
	BlockID int `json:"block_id,omitempty"`
}

// GetConsensusBlockResponse represents the response from the /getConsensusBlock endpoint
type GetConsensusBlockResponse struct {
	OK     bool `json:"ok"`
	Result struct {
		Consensus struct {
			SeqNo     int    `json:"seq_no"`
			RootHash  string `json:"root_hash"`
			FileHash  string `json:"file_hash"`
			Timestamp int    `json:"timestamp"`
		} `json:"consensus"`
		Pending struct {
			SeqNo     int    `json:"seq_no"`
			RootHash  string `json:"root_hash"`
			FileHash  string `json:"file_hash"`
			Timestamp int    `json:"timestamp"`
		} `json:"pending"`
	} `json:"result"`
}

// GetConsensusBlock gets the consensus block
func (c *Client) GetConsensusBlock(req *GetConsensusBlockRequest) (*GetConsensusBlockResponse, error) {
	endpoint := "/getConsensusBlock"
	
	respBody, err := c.doRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}
	
	var response GetConsensusBlockResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w (code: %d)", err, ErrInvalidResponse)
	}
	
	if !response.OK {
		return nil, fmt.Errorf("API returned non-OK status (code: %d)", ErrAPIError)
	}
	
	return &response, nil
}

// GetExtendedAddressInformationResponse represents the response from the /getExtendedAddressInformation endpoint
type GetExtendedAddressInformationResponse struct {
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
		BlockID         struct {
			Workchain int    `json:"workchain"`
			Shard     string `json:"shard"`
			SeqNo     int    `json:"seqno"`
			RootHash  string `json:"root_hash"`
			FileHash  string `json:"file_hash"`
		} `json:"block_id"`
		Parsed struct {
			Status    string `json:"status"`
			Timestamp int    `json:"timestamp"`
			IsWallet  bool   `json:"is_wallet"`
			WalletType string `json:"wallet_type"`
			SeqNo     int    `json:"seqno"`
			PublicKey string `json:"public_key"`
			WalletID  int    `json:"wallet_id"`
		} `json:"parsed"`
	} `json:"result"`
}

// GetExtendedAddressInformation gets extended information about a TON address
func (c *Client) GetExtendedAddressInformation(address string) (*GetExtendedAddressInformationResponse, error) {
	endpoint := fmt.Sprintf("/getExtendedAddressInformation?address=%s", address)
	
	respBody, err := c.doRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	
	var response GetExtendedAddressInformationResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w (code: %d)", err, ErrInvalidResponse)
	}
	
	if !response.OK {
		return nil, fmt.Errorf("API returned non-OK status (code: %d)", ErrAPIError)
	}
	
	return &response, nil
}
