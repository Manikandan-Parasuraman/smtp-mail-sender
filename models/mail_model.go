package models

type MailInput struct {
	Email            string `json:"Email"`
	Username         string `json:"Username"`
	BodyTemplateName string `json:"BodyTemplateName"`
	AttachementPath  string `json:"AttachementPath,omitempty"`
	PanicTrigger     bool   `json:"PanicTrigger,omitempty"`
}