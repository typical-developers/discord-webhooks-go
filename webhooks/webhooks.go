package webhooks

import (
	"fmt"
	"net/url"

	"github.com/opensaucerer/goaxios"
)

// Create a new WebhookClient using a clientId and secret.
func NewWebhookClient(clientId string, secret string) *WebhookClient {
	return &WebhookClient{
		BaseURL: url.URL{
			Scheme: "https",
			Host:   "discord.com",
			Path:   "/api/webhooks/" + clientId + "/" + secret,
		},
		DefaultUserInfo: &WebhookUserInfo{},
	}
}

// Create a new WebhookClient using the direct webhook url.
func NewWebhookClientFromURL(webhookUrl string) *WebhookClient {
	parsedUrl, err := url.Parse(webhookUrl)

	if err != nil {
		panic(err)
	}

	return &WebhookClient{
		BaseURL:         *parsedUrl,
		DefaultUserInfo: &WebhookUserInfo{},
	}
}

// Set a default username for the webhook.
// This prevents you from having to provide it in WebhookPayload.Username constantly.
//
// If you provide WebhookPayload.Username, this value will be ignored.
func (c *WebhookClient) SetDefaultUsername(username string) {
	c.DefaultUserInfo.Username = &username
}

// Set a default avatar url for the webhook.
// This prevents you from having to provide it in WebhookPayload.AvatarURL constantly.
//
// If you provide WebhookPayload.AvatarURL, this value will be ignored.
func (c *WebhookClient) SetDefaultAvatarURL(avatarUrl string) {
	c.DefaultUserInfo.AvatarURL = &avatarUrl
}

// Delete the webhook.
func (c *WebhookClient) DeleteWebhook() (bool, error) {
	request := goaxios.GoAxios{
		Method: "DELETE",
		Url:    c.BaseURL.String(),
	}

	res := request.RunRest()
	if res.Error != nil {
		return false, res.Error
	}

	return true, nil
}

// Send a message from the webhook.
func (w *WebhookClient) SendMessage(args *WebhookPayload) (response *WebhookPayloadResponse, err error) {
	err = w.validatePayload(args)
	if err != nil {
		return nil, err
	}

	requestUrl := w.BaseURL
	query := requestUrl.Query()
	query.Add("wait", "true")
	requestUrl.RawQuery = query.Encode()

	var r goaxios.GoAxios
	if args.Files != nil {
		form, err := args.parseFormData()
		if err != nil {
			return nil, err
		}

		r = goaxios.GoAxios{
			Method: "POST",
			Url:    requestUrl.String(),
			Form:   form,
		}
	} else {
		r = goaxios.GoAxios{
			Method: "POST",
			Url:    requestUrl.String(),
			Body:   args,
		}
	}

	r.ResponseStruct = &WebhookPayloadResponse{}

	res := r.RunRest()
	if res.Error != nil {
		return nil, res.Error
	}

	response, ok := res.Body.(*WebhookPayloadResponse)
	if !ok {
		return nil, fmt.Errorf("An invalid WebhookPayloadResponse was returned.")
	}

	return response, nil
}

// Edit an existing message that the webhook has sent.
func (w *WebhookClient) EditMessage(messageId string, args *WebhookPayload) (response *WebhookPayloadResponse, err error) {
	err = w.validatePayload(args)
	if err != nil {
		return nil, err
	}

	requestUrl := w.BaseURL
	requestUrl.Path = requestUrl.Path + "/messages/" + messageId

	query := requestUrl.Query()
	query.Add("wait", "true")
	requestUrl.RawQuery = query.Encode()

	var r goaxios.GoAxios
	if args.Files != nil {
		form, err := args.parseFormData()
		if err != nil {
			return nil, err
		}

		r = goaxios.GoAxios{
			Method: "PATCH",
			Url:    requestUrl.String(),
			Form:   form,
		}
	} else {
		r = goaxios.GoAxios{
			Method: "PATCH",
			Url:    requestUrl.String(),
			Body:   args,
		}
	}

	r.ResponseStruct = &WebhookPayloadResponse{}

	res := r.RunRest()
	if res.Error != nil {
		return nil, res.Error
	}

	response, ok := res.Body.(*WebhookPayloadResponse)
	if !ok {
		return nil, fmt.Errorf("An invalid WebhookPayloadResponse was returned.")
	}

	return response, nil
}

// Delete an existing message that the webhook has sent.
func (w *WebhookClient) DeleteMessage(messageId string) (err error) {
	requestUrl := w.BaseURL
	requestUrl.Path = requestUrl.Path + "/messages/" + messageId

	r := goaxios.GoAxios{
		Method: "DELETE",
		Url:    requestUrl.String(),
	}

	res := r.RunRest()
	if res.Error != nil {
		return res.Error
	}

	return nil
}
