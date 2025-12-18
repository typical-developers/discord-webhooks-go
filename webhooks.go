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
func (c *WebhookClient) NewRequest(method string, body any) (*Request, error) {
	req, err := http.NewRequest(method, c.webhookUrl.String(), nil)
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
func (c *WebhookClient) NewMultipartRequest(method string, fields map[string]any) (*Request, error) {
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

	req, err := http.NewRequest(method, c.webhookUrl.String(), &body)
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

type ExecuteParams struct {
	Wait           bool   `url:"wait,omitempty"`
	ThreadID       string `url:"thread_id,omitempty"`
	WithComponents bool   `url:"with_components,omitempty"`
}

// Execute will execute the webhook with the provided message and params.
func (c *WebhookClient) Execute(ctx context.Context, content MessagePayload, params *ExecuteParams) (*http.Response, error) {
	var request *Request

	if len(content.Files) > 0 {
		payload := make(map[string]any)
		payload["payload_json"] = content.PayloadJSON()
		for i, file := range content.Files {
			key := fmt.Sprintf("files[%d]", i)
			payload[key] = file
		}

		req, err := c.NewMultipartRequest(http.MethodPost, payload)
		if err != nil {
			return nil, err
		}

		request = req
	} else {
		req, err := c.NewRequest(http.MethodPost, content)
		if err != nil {
			return nil, err
		}

		request = req
	}

	return c.Do(ctx, request)
}
