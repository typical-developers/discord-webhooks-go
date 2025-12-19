package webhooks

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
)

type WebhookClient struct {
	webhookUrl url.URL
	client     *http.Client
}

type Request struct {
	*http.Request
}

func NewWebhookClient(id, secret string) *WebhookClient {
	url := url.URL{
		Scheme: "https",
		Host:   "discord.com",
		Path:   "/api/webhooks/" + id + "/" + secret,
	}

	return &WebhookClient{
		webhookUrl: url,
		client:     http.DefaultClient,
	}
}

func NewWebhookClientFromURL(webhookUrl string) *WebhookClient {
	u, err := url.Parse(webhookUrl)
	if err != nil {
		panic(err)
	}

	return &WebhookClient{
		webhookUrl: *u,
		client:     http.DefaultClient,
	}
}

// NewRequest will create a new Request using the webhook url and provided parameters.
func (c *WebhookClient) NewRequest(method, url string, body any) (*Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	if body != nil {
		var buf bytes.Buffer
		encoder := json.NewEncoder(&buf)
		encoder.SetEscapeHTML(false)
		err := encoder.Encode(body)
		if err != nil {
			return nil, err
		}

		req.Body = io.NopCloser(&buf)
		req.Header.Set("Content-Type", "application/json")
	}

	return &Request{req}, nil
}

// NewMultipartRequest will create a new Request using the webhook url and provided parameters.
func (c *WebhookClient) NewMultipartRequest(method, url string, fields map[string]any) (*Request, error) {
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	for k, v := range fields {
		switch content := v.(type) {
		case WebhookFile:
			part, err := writer.CreateFormFile(k, content.FileName)
			if err != nil {
				return nil, err
			}

			if _, err := io.Copy(part, content.Reader); err != nil {
				return nil, err
			}
		case string:
			part, err := writer.CreateFormField(k)
			if err != nil {
				return nil, err
			}

			if _, err := part.Write([]byte(content)); err != nil {
				return nil, err
			}
		}
	}

	if err := writer.Close(); err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, &body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	return &Request{req}, nil
}

// Do will execute the http.Request.
func (c *WebhookClient) Do(ctx context.Context, req *Request) (*http.Response, error) {
	return c.client.Do(req.Request)
}

// Execute will execute the webhook with the provided message and params.
func (c *WebhookClient) Execute(ctx context.Context, content MessagePayload, params *url.Values) (*WebhookMessage, *http.Response, error) {
	var request *Request

	u := c.webhookUrl
	if params != nil {
		u.RawQuery = params.Encode()
	}

	url := u.String()
	if len(content.Files) > 0 {
		payload := make(map[string]any)
		payload["payload_json"] = content.PayloadJSON()
		for i, file := range content.Files {
			key := fmt.Sprintf("files[%d]", i)
			payload[key] = file
		}

		req, err := c.NewMultipartRequest(http.MethodPost, url, payload)
		if err != nil {
			return nil, nil, err
		}

		request = req
	} else {
		req, err := c.NewRequest(http.MethodPost, url, content)
		if err != nil {
			return nil, nil, err
		}

		request = req
	}

	res, err := c.Do(ctx, request)
	if err != nil {
		return nil, res, err
	}

	// 204 No Content = Wait query was not provided, no WebhookMessage should be returned.
	if res.StatusCode == http.StatusNoContent {
		return nil, res, nil
	}

	message := new(WebhookMessage)
	message.c = c

	if err := json.NewDecoder(res.Body).Decode(message); err != nil {
		return nil, res, err
	}

	return message, res, nil
}

func (c *WebhookClient) EditMessage(ctx context.Context, messageID string, content EditMessagePayload, params *url.Values) (*WebhookMessage, *http.Response, error) {
	var request *Request

	u := c.webhookUrl.
		JoinPath("messages").
		JoinPath(messageID)
	if params != nil {
		u.RawQuery = params.Encode()
	}

	url := u.String()
	if len(content.Files) > 0 {
		payload := make(map[string]any)

		payload["payload_json"] = content.PayloadJSON()
		for i, file := range content.Files {
			key := fmt.Sprintf("files[%d]", i)
			payload[key] = file
		}

		req, err := c.NewMultipartRequest(http.MethodPatch, url, payload)
		if err != nil {
			return nil, nil, err
		}

		request = req
	} else {
		req, err := c.NewRequest(http.MethodPatch, url, content)
		if err != nil {
			return nil, nil, err
		}

		request = req
	}

	res, err := c.Do(ctx, request)
	if err != nil {
		return nil, res, err
	}

	message := new(WebhookMessage)
	message.c = c

	if err := json.NewDecoder(res.Body).Decode(message); err != nil {
		return nil, res, err
	}

	return message, res, nil
}
