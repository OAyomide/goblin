package messaging

import "encoding/json"

//Buttons struct for the button "object" on the carousel
type Buttons struct {
	Type    string `json:"type"`
	URL     string `json:"url,omitempty"`
	Title   string `json:"title"`
	Payload string `json:"payload,omitempty"`
}

//Payload struct for the payload object
type Payload struct {
	TemplateType string     `json:"template_type"`
	Els          []Elements `json:"elements"`
}

//Elements struct for the elements object
type Elements struct {
	Title    string `json:"title"`
	Subt     string `json:"subtitle,omitempty"`
	ImageURL string `json:"image_url"`
	//DefAction DefaultAction `json:"default_action,omitempty"`
	Buttons []Buttons `json:"buttons,omitempty"`
}

//Attachment struct for the attachment object
type Attachment struct {
	Type string  `json:"type"`
	Payl Payload `json:"payload"`
}

//Recipient struct for the recipient object
type Recipient struct {
	ID string `json:"id"`
}

//Messages struct for the messages object
type Messages struct {
	Attach Attachment `json:"attachment"`
}

//Base struct for the whole carousel object
type Base struct {
	Rec Recipient `json:"recipient"`
	Mes Messages  `json:"message"`
}

//SendCarousel is a Reciever function to send carousel
func (c *Base) SendCarousel(userID string) {
	marsh, err := json.Marshal(c)

	callSendAPI(marsh)
}
