package toncenterzp

import (
	"fmt"
	"net/url"
	"strings"
)

// Error codes
const (
	// Client errors (1xxx)
	ErrInvalidParams    = 1001
	ErrNetworkError     = 1002
	ErrInvalidResponse  = 1003
	ErrAPIError         = 1004
	ErrInvalidAddress   = 1005
	ErrInvalidSignature = 1006
	ErrInvalidBlock     = 1007
	ErrInvalidMethod    = 1008
	ErrTimeout          = 1009
	ErrInvalidConfig    = 1010
	
	// Server errors (2xxx)
	ErrServerInternal   = 2001
	ErrServerOverload   = 2002
	ErrServerUnavailable = 2003
	
	// TON specific errors (3xxx)
	ErrInvalidTransaction = 3001
	ErrInvalidBoc         = 3002
	ErrInvalidCell        = 3003
	ErrContractExecution  = 3004
)

// ErrorWithCode represents an error with a code
type ErrorWithCode struct {
	Code    int
	Message string
	Err     error
}

// Error returns the error message
func (e *ErrorWithCode) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v (code: %d)", e.Message, e.Err, e.Code)
	}
	return fmt.Sprintf("%s (code: %d)", e.Message, e.Code)
}

// Unwrap returns the wrapped error
func (e *ErrorWithCode) Unwrap() error {
	return e.Err
}

// NewError creates a new error with a code
func NewError(code int, message string, err error) *ErrorWithCode {
	return &ErrorWithCode{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// IsValidAddress checks if an address is valid
func IsValidAddress(address string) bool {
	// Basic validation for TON addresses
	if len(address) == 0 {
		return false
	}
	
	// Check if it's a raw address (hex)
	if strings.HasPrefix(address, "0:") && len(address) >= 66 {
		return true
	}
	
	// Check if it's a user-friendly address
	if strings.Contains(address, ":") {
		parts := strings.Split(address, ":")
		if len(parts) != 2 {
			return false
		}
		
		// Check workchain
		if parts[0] != "0" && parts[0] != "-1" {
			return false
		}
		
		// Check hash part
		if len(parts[1]) != 64 {
			return false
		}
		
		return true
	}
	
	// Check if it's a base64 address
	if len(address) >= 48 && (strings.HasPrefix(address, "EQ") || strings.HasPrefix(address, "UQ") || 
	                          strings.HasPrefix(address, "kQ") || strings.HasPrefix(address, "Ef")) {
		return true
	}
	
	return false
}

// EncodeQueryParams encodes query parameters
func EncodeQueryParams(params map[string]string) string {
	if len(params) == 0 {
		return ""
	}
	
	values := url.Values{}
	for key, value := range params {
		values.Add(key, value)
	}
	
	return "?" + values.Encode()
}

// FormatNanoTON formats nano TON to TON
func FormatNanoTON(nanoTON string) (string, error) {
	// Remove leading zeros
	nanoTON = strings.TrimLeft(nanoTON, "0")
	if nanoTON == "" {
		nanoTON = "0"
	}
	
	// Ensure it's at least 10 digits (9 decimal places)
	for len(nanoTON) < 10 {
		nanoTON = "0" + nanoTON
	}
	
	// Insert decimal point
	decimalPos := len(nanoTON) - 9
	result := nanoTON[:decimalPos] + "." + nanoTON[decimalPos:]
	
	// Remove trailing zeros after decimal point
	result = strings.TrimRight(result, "0")
	if result[len(result)-1] == '.' {
		result = result[:len(result)-1]
	}
	
	return result, nil
}

// ValidateAPIKey checks if an API key is valid
func ValidateAPIKey(apiKey string) error {
	if apiKey == "" {
		return NewError(ErrInvalidConfig, "API key is required", nil)
	}
	
	// Basic validation for API key format
	if len(apiKey) < 8 {
		return NewError(ErrInvalidConfig, "API key is too short", nil)
	}
	
	return nil
}
