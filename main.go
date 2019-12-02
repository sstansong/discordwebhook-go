package dcwebhook

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
)

type Thumbnail struct {
	URL string `json:"thumbnail,omitempty"`
}

type Image struct {
	URL string `json:"url,omitempty"`
}


type Footer struct {
	Text string `json:"text,omitempty"`
	IconUrl string `json:"icon_url,omitempty"`
}


type Author struct {
	Name string `json:"name,omitempty"`
	URL string `json:"url,omitempty"`
	IconUrl string `json:"icon_url,omitempty"`
}

type Field struct {
	Name string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
	Inline bool `json:"inline,omitempty"`
}


type Embed struct{
	Title string `json:"title,omitempty"`
	URL string `json:"url,omitempty"`
	Description string `json:"description,omitempty"`
	Color int `json:"color,omitempty"`
	Author Author `json:"author,omitempty"`
	Fields []Field `json:"fields,omitempty"`
	Image Image `json:"image,omitempty"`
	Thumbnail Thumbnail `thumbnail:"image,omitempty"`
	Footer Footer `json:"footer,omitempty"`
}

type Payload struct {
	Username    string       `json:"username,omitempty"`
	TTS			bool		 `json:"tts,omitempty"`
	IconUrl     string       `json:"avatar_url,omitempty"`
	Content     string       `json:"content,omitempty"`
	Embeds      []Embed      `json:"embeds,omitempty"`
}

func (embed *Embed)AddFieldtoEmbed(field Field) *Embed{
	embed.Fields = append(embed.Fields, field)
	return embed
}

func (payload *Payload)AddEmbed(embed Embed) *Payload{
	payload.Embeds = append(payload.Embeds, embed)
	return payload
}

func redirectPolicyFunc(req gorequest.Request, via []gorequest.Request) error {
	return fmt.Errorf("Incorrect token (redirection)")
}

func Send(webhookUrl string, proxy string, payload Payload) []error {
	request := gorequest.New().Proxy(proxy)
	resp, _, err := request.
		Post(webhookUrl).
		RedirectPolicy(redirectPolicyFunc).
		Send(payload).
		End()

	if err != nil {
		return err
	}
	if resp.StatusCode >= 400 {
		return []error{fmt.Errorf("Error sending msg. Status: %v", resp.Status)}
	}

	return nil
}