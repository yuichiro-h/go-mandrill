package mandrill

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
)

const baseURL = "https://mandrillapp.com/api/1.0"

type Client struct {
	apiKey string
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey,
	}
}

func (c *Client) newRequest(URI string) *gorequest.SuperAgent {
	return gorequest.New().Post(baseURL+URI).
		Set("User-Agent", fmt.Sprintf("go-mandrill/%s", Version)).
		Send(fmt.Sprintf(`{"key":"%s"}`, c.apiKey))
}
