package slack

import (
	"fmt"
	"log"

	"github.com/parnurzeal/gorequest"
)

type Field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

type Attachment struct {
	Fallback   *string  `json:"fallback"`
	Color      *string  `json:"color"`
	PreText    *string  `json:"pretext"`
	AuthorName *string  `json:"author_name"`
	AuthorLink *string  `json:"author_link"`
	AuthorIcon *string  `json:"author_icon"`
	Title      *string  `json:"title"`
	TitleLink  *string  `json:"title_link"`
	Text       *string  `json:"text"`
	ImageUrl   *string  `json:"image_url"`
	Fields     []*Field `json:"fields"`
}

type Payload struct {
	Parse				string				`json:"parse,omitempty"`
	Username		string				`json:"username,omitempty"`
	IconUrl			string				`json:"icon_url,omitempty"`
	IconEmoji		string				`json:"icon_emoji,omitempty"`
	Channel			string				`json:"channel,omitempty"`
	Text				string				`json:"text,omitempty"`
	Attachments	[]Attachment	`json:"attachments,omitempty"`
}

func (attachment *Attachment) AddField(field Field) *Attachment {
	attachment.Fields = append(attachment.Fields, &field)
	return attachment
}

func redirectPolicyFunc(req gorequest.Request, via []gorequest.Request) error {
      return fmt.Errorf("Incorrect token (redirection)")
}

func Send(token string, proxy string, payload Payload) []error {
  webhookUrl := fmt.Sprintf("https://hooks.slack.com/services/%v", token)

	request := gorequest.New().Proxy(proxy)
	resp, _, err := request.
		Post(webhookUrl).
		RedirectPolicy(redirectPolicyFunc).
		Send(payload).
		End()

	if err != nil {
		log.Fatal(err)
		return err
	}
	if resp.StatusCode >= 400 {
		return []error{fmt.Errorf("Error sending msg. Status: %v", resp.Status)}
	}

	return nil
}
