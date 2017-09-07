package mandrill

import "encoding/json"

type Var struct {
	Name    string      `json:"name"`
	Content interface{} `json:"content"`
}

type RecipientVar struct {
	RCPT string `json:"rcpt"`
	Var  []Var  `json:"vars"`
}

type RecipientMetadata struct {
	RCPT   string                 `json:"rcpt"`
	Values map[string]interface{} `json:"values"`
}

type ToType *string

var (
	ToTypeTo  ToType = String("to")
	ToTypeCC         = String("cc")
	ToTypeBCC        = String("bcc")
)

type To struct {
	Email string  `json:"email"`
	Name  *string `json:"name"`
	Type  ToType  `json:"type"`
}

type MergeLanguageType *string

var (
	MergeLanguageTypeMailchimp  MergeLanguageType = String("mailchimp")
	MergeLanguageTypeHandlebars                   = String("handlebars")
)

type Attachment struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

type Message struct {
	HTML                    *string              `json:"html"`
	Text                    *string              `json:"text"`
	Subject                 *string              `json:"subject"`
	FromEmail               *string              `json:"from_email"`
	FromName                *string              `json:"from_name"`
	To                      []To                 `json:"to"`
	Headers                 *map[string]string   `json:"headers"`
	Important               *bool                `json:"important"`
	TrackOpens              *bool                `json:"track_opens"`
	TrackClicks             *bool                `json:"track_clicks"`
	AutoText                *bool                `json:"auto_text"`
	AutoHTML                *bool                `json:"auto_html"`
	InlineCSS               *bool                `json:"inline_css"`
	URLStripQS              *bool                `json:"url_strip_qs"`
	PreserveRecipients      *bool                `json:"preserve_recipients"`
	ViewContentLink         *bool                `json:"view_content_link"`
	BCCAddress              *string              `json:"bcc_address"`
	TrackingDomain          *string              `json:"tracking_domain"`
	SigningDomain           *string              `json:"signing_domain"`
	ReturnPathDomain        *string              `json:"return_path_domain"`
	Merge                   *bool                `json:"merge"`
	MergeLanguage           MergeLanguageType    `json:"merge_language"`
	GlobalMergeVars         *[]Var               `json:"global_merge_vars"`
	MergeVars               *[]RecipientVar      `json:"merge_vars"`
	Tags                    *[]string            `json:"tags"`
	Subaccount              *string              `json:"subaccount"`
	GoogleAnalyticsDomains  *[]string            `json:"google_analytics_domains"`
	GoogleAnalyticsCampaign *string              `json:"google_analytics_campaign"`
	Metadata                *[]string            `json:"metadata"`
	RecipientMetadata       *[]RecipientMetadata `json:"recipient_metadata"`
	Attachments             *[]Attachment        `json:"attachments"`
	Images                  *[]Attachment        `json:"images"`
}

type MessageSendTemplateRequest struct {
	TemplateName    string  `json:"template_name"`
	TemplateContent *[]Var  `json:"template_content"`
	Message         Message `json:"message"`
	Async           *bool   `json:"async"`
	IPPool          *string `json:"ip_pool"`
	SendAt          *string `json:"send_at"`
}

type MessageStatusType string

const (
	MessageStatusTypeSent     MessageStatusType = "sent"
	MessageStatusTypeQueued                     = "queued"
	MessageStatusTypeRejected                   = "rejected"
	MessageStatusTypeInvalid                    = "invalid"
)

type RejectReasonType string

const (
	RejectReasonTypeHardBounce    RejectReasonType = "hard-bounce"
	RejectReasonTypeSoftBounce                     = "soft-bounce"
	RejectReasonTypeSpam                           = "spam"
	RejectReasonTypeUnsub                          = "unsub"
	RejectReasonTypeCustom                         = "custom"
	RejectReasonTypeInvalidSender                  = "invalid-sender"
	RejectReasonTypeInvalid                        = "invalid"
	RejectReasonTypeTestModeLimit                  = "test-mode-limit"
	RejectReasonTypeUnsigned                       = "unsigned"
	RejectReasonTypeRule                           = "rule"
)

type MessageSendReturn struct {
	Email        string            `json:"email"`
	Status       MessageStatusType `json:"status"`
	RejectReason RejectReasonType  `json:"reject_reason"`
	ID           string            `json:"_id"`
}

func (c *Client) MessageSendTemplate(req *MessageSendTemplateRequest) ([]MessageSendReturn, error) {
	resp, body, errs := c.newRequest("/messages/send-template.json").Send(req).End()
	if len(errs) > 0 {
		return nil, errs[0]
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, c.parseError(body)
	}

	var res []MessageSendReturn
	if err := json.Unmarshal([]byte(body), &res); err != nil {
		return nil, err
	}

	return res, nil
}
