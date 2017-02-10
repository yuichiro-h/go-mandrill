package mandrill

import "encoding/json"

type RejectDeleteRequest struct {
	Email      string `json:"email"`
	SubAccount string `json:"subaccount"`
}

type RejectDeleteResponse struct {
	Email      string `json:"email"`
	Deleted    bool   `json:"deleted"`
	SubAccount string `json:"subaccount"`
}

func (c *Client) RejectDelete(req *RejectDeleteRequest) (*RejectDeleteResponse, error) {
	resp, body, errs := c.newRequest("/rejects/delete.json").Send(req).End()
	if len(errs) > 0 {
		return nil, errs[0]
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, c.parseError(body)
	}

	var res RejectDeleteResponse
	if err := json.Unmarshal([]byte(body), &res); err != nil {
		return nil, err
	}

	return &res, nil
}
