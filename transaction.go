package toncenterzp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetTransactionsRequest represents the request for the /getTransactions endpoint
type GetTransactionsRequest struct {
	Address string `json:"address"`
	Limit   int    `json:"limit,omitempty"`
	Lt      string `json:"lt,omitempty"`
	Hash    string `json:"hash,omitempty"`
	ToLt    string `json:"to_lt,omitempty"`
	ArchiveOnly bool `json:"archive_only,omitempty"`
}

// TransactionDetails represents the details of a transaction
type TransactionDetails struct {
	Data           string `json:"data"`
	Fee            string `json:"fee"`
	OtherFee       string `json:"other_fee"`
	StorageFee     string `json:"storage_fee"`
	GasFee         string `json:"gas_fee"`
	FwdFee         string `json:"fwd_fee"`
	TotalFees      string `json:"total_fees"`
	InMsg          Message `json:"in_msg"`
	OutMsgs        []Message `json:"out_msgs"`
	BlockID        BlockID `json:"block_id"`
	PrevTransHash  string `json:"prev_trans_hash"`
	PrevTransLt    string `json:"prev_trans_lt"`
	Now            int    `json:"now"`
	OutMsgsCount   int    `json:"outmsg_cnt"`
	OrigStatus     string `json:"orig_status"`
	EndStatus      string `json:"end_status"`
	AccountAddr    string `json:"account_addr"`
	Lt             string `json:"lt"`
	Hash           string `json:"hash"`
	Description    string `json:"description"`
	ComputePhase   ComputePhase `json:"compute_ph"`
	ActionPhase    ActionPhase `json:"action"`
	CreditPhase    CreditPhase `json:"credit_ph"`
	StoragePhase   StoragePhase `json:"storage_ph"`
	BouncePhase    BouncePhase `json:"bounce"`
}

// Message represents a message in a transaction
type Message struct {
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Value       string `json:"value"`
	FwdFee      string `json:"fwd_fee"`
	IhrFee      string `json:"ihr_fee"`
	CreatedLt   string `json:"created_lt"`
	BodyHash    string `json:"body_hash"`
	MsgType     string `json:"msg_type"`
	MsgData     MessageData `json:"msg_data"`
}

// MessageData represents the data of a message
type MessageData struct {
	Text     string `json:"text"`
	InitState string `json:"init_state"`
	Body     string `json:"body"`
}

// BlockID represents a block ID
type BlockID struct {
	Workchain int    `json:"workchain"`
	Shard     string `json:"shard"`
	SeqNo     int    `json:"seqno"`
	RootHash  string `json:"root_hash"`
	FileHash  string `json:"file_hash"`
}

// ComputePhase represents the compute phase of a transaction
type ComputePhase struct {
	SkippedReason string `json:"skipped_reason"`
	Success       bool   `json:"success"`
	GasUsed       string `json:"gas_used"`
	VmSteps       int    `json:"vm_steps"`
	ExitCode      int    `json:"exit_code"`
}

// ActionPhase represents the action phase of a transaction
type ActionPhase struct {
	Success        bool   `json:"success"`
	Valid          bool   `json:"valid"`
	NoFunds        bool   `json:"no_funds"`
	StatusChange   string `json:"status_change"`
	TotalFwdFees   string `json:"total_fwd_fees"`
	TotalActionFees string `json:"total_action_fees"`
	ResultCode     int    `json:"result_code"`
	TotalActions   int    `json:"tot_actions"`
}

// CreditPhase represents the credit phase of a transaction
type CreditPhase struct {
	DueFeesCollected string `json:"due_fees_collected"`
	Credit           string `json:"credit"`
}

// StoragePhase represents the storage phase of a transaction
type StoragePhase struct {
	StorageFeesCollected string `json:"storage_fees_collected"`
	StatusChange         string `json:"status_change"`
}

// BouncePhase represents the bounce phase of a transaction
type BouncePhase struct {
	BounceType string `json:"bounce_type"`
	FwdFees    string `json:"fwd_fees"`
	MsgFees    string `json:"msg_fees"`
	ReqFwdFees string `json:"req_fwd_fees"`
}

// GetTransactionsResponse represents the response from the /getTransactions endpoint
type GetTransactionsResponse struct {
	OK     bool `json:"ok"`
	Result struct {
		Transactions []TransactionDetails `json:"transactions"`
	} `json:"result"`
}

// GetTransactions gets transactions for a TON address
func (c *Client) GetTransactions(req GetTransactionsRequest) (*GetTransactionsResponse, error) {
	endpoint := "/getTransactions"
	
	respBody, err := c.doRequest(http.MethodPost, endpoint, req)
	if err != nil {
		return nil, err
	}
	
	var response GetTransactionsResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w (code: %d)", err, ErrInvalidResponse)
	}
	
	if !response.OK {
		return nil, fmt.Errorf("API returned non-OK status (code: %d)", ErrAPIError)
	}
	
	return &response, nil
}

// GetWalletInformationResponse represents the response from the /getWalletInformation endpoint
type GetWalletInformationResponse struct {
	OK     bool `json:"ok"`
	Result struct {
		Wallet        bool   `json:"wallet"`
		Balance       string `json:"balance"`
		Account       string `json:"account"`
		WalletType    string `json:"wallet_type"`
		SeqNo         int    `json:"seqno"`
		LastTransLt   string `json:"last_trans_lt"`
		LastTransHash string `json:"last_trans_hash"`
		WalletID      int    `json:"wallet_id"`
		PublicKey     string `json:"public_key"`
	} `json:"result"`
}

// GetWalletInformation gets information about a TON wallet
func (c *Client) GetWalletInformation(address string) (*GetWalletInformationResponse, error) {
	endpoint := fmt.Sprintf("/getWalletInformation?address=%s", address)
	
	respBody, err := c.doRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	
	var response GetWalletInformationResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w (code: %d)", err, ErrInvalidResponse)
	}
	
	if !response.OK {
		return nil, fmt.Errorf("API returned non-OK status (code: %d)", ErrAPIError)
	}
	
	return &response, nil
}
