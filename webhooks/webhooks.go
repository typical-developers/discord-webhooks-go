package webhooks

import (
	"fmt"
	"net/url"

	"github.com/opensaucerer/goaxios"
)

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

func (c *WebhookClient) SetDefaultUsername(username string) {
	c.DefaultUserInfo.Username = &username
}

func (c *WebhookClient) SetDefaultAvatarURL(avatarUrl string) {
	c.DefaultUserInfo.AvatarURL = &avatarUrl
}

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
