package mandrill

import (
	"encoding/json"
	"fmt"
)

type APIError struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("Mandrill API Error. status=%s, code=%d, name=%s, message=%s",
		e.Status, e.Code, e.Name, e.Message)
}

func (c *Client) parseError(response string) error {
	var apiErr APIError
	if err := json.Unmarshal([]byte(response), &apiErr); err != nil {
		return err
	}

	return &apiErr
}

const (
	ErrNameInvalidKey        = "Invalid_Key"
	ErrNameValidationError   = "ValidationError"
	ErrNameGeneralError      = "GeneralError"
	ErrNameInvalidReject     = "Invalid_Reject"
	ErrNameUnknownSubaccount = "Unknown_Subaccount"
	ErrNameInvalidTemplate   = "Invalid_Template"
	ErrNameUnknownTemplate   = "Unknown_Template"
)
