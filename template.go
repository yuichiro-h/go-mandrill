package mandrill

import "encoding/json"

type Template struct {
	Slug             string   `json:"slug"`
	Name             string   `json:"name"`
	Labels           []string `json:"labels"`
	Code             string   `json:"code"`
	Subject          string   `json:"subject"`
	FromEmail        string   `json:"from_email"`
	FromName         string   `json:"from_name"`
	Text             string   `json:"text"`
	PublishName      string   `json:"publish_name"`
	PublishCode      string   `json:"publish_code"`
	PublishSubject   string   `json:"publish_subject"`
	PublishFromEmail string   `json:"publish_from_email"`
	PublishFromName  string   `json:"publish_from_name"`
	PublishText      string   `json:"publish_text"`
	PublishedAt      string   `json:"published_at"`
	CreatedAt        string   `json:"created_at"`
	UpdatedAt        string   `json:"updated_at"`
}

type TemplateInfoRequest struct {
	Name string `json:"name"`
}

func (c *Client) TemplateInfo(req *TemplateInfoRequest) (*Template, error) {
	resp, body, errs := c.newRequest("/templates/info.json").Send(req).End()
	if len(errs) > 0 {
		return nil, errs[0]
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, c.parseError(body)
	}

	var res Template
	if err := json.Unmarshal([]byte(body), &res); err != nil {
		return nil, err
	}

	return &res, nil
}

type TemplateListRequest struct {
	Label *string `json:"label"`
}

func (c *Client) TemplateList(req *TemplateListRequest) (*[]Template, error) {
	resp, body, errs := c.newRequest("/templates/list.json").Send(req).End()
	if len(errs) > 0 {
		return nil, errs[0]
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, c.parseError(body)
	}

	var res []Template
	if err := json.Unmarshal([]byte(body), &res); err != nil {
		return nil, err
	}

	return &res, nil
}

type TemplateAddRequest struct {
	Name      string    `json:"name"`
	FromEmail *string   `json:"from_email"`
	FromName  *string   `json:"from_name"`
	Subject   *string   `json:"subject"`
	Code      *string   `json:"code"`
	Text      *string   `json:"text"`
	Publish   bool      `json:"publish"`
	Labels    *[]string `json:"labels"`
}

func (c *Client) TemplateAdd(req *TemplateAddRequest) (*Template, error) {
	resp, body, errs := c.newRequest("/templates/add.json").Send(req).End()
	if len(errs) > 0 {
		return nil, errs[0]
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, c.parseError(body)
	}

	var res Template
	if err := json.Unmarshal([]byte(body), &res); err != nil {
		return nil, err
	}

	return &res, nil
}

type TemplateUpdateRequest struct {
	Name      string    `json:"name"`
	FromEmail *string   `json:"from_email"`
	FromName  *string   `json:"from_name"`
	Subject   *string   `json:"subject"`
	Code      *string   `json:"code"`
	Text      *string   `json:"text"`
	Publish   bool      `json:"publish"`
	Labels    *[]string `json:"labels"`
}

func (c *Client) TemplateUpdate(req *TemplateUpdateRequest) (*Template, error) {
	resp, body, errs := c.newRequest("/templates/update.json").Send(req).End()
	if len(errs) > 0 {
		return nil, errs[0]
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, c.parseError(body)
	}

	var res Template
	if err := json.Unmarshal([]byte(body), &res); err != nil {
		return nil, err
	}

	return &res, nil
}

type TemplateDeleteRequest struct {
	Name string `json:"name"`
}

func (c *Client) TemplateDelete(req *TemplateDeleteRequest) (*Template, error) {
	resp, body, errs := c.newRequest("/templates/delete.json").Send(req).End()
	if len(errs) > 0 {
		return nil, errs[0]
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, c.parseError(body)
	}

	var res Template
	if err := json.Unmarshal([]byte(body), &res); err != nil {
		return nil, err
	}

	return &res, nil
}
