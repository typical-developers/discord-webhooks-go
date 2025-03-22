package webhooks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/opensaucerer/goaxios"
)

func NewWebhook(args NewWebhookArgs) *WebhookClient {
	var webhookUrl string

	if args.URL != nil {
		parsedUrl, err := url.Parse(*args.URL)

		if err != nil {
			panic(err)
		}

		webhookUrl = parsedUrl.String()
	} else {
		webhookUrl = createWebhookURL(*args.ClientID, *args.Secret)
	}

	return &WebhookClient{
		URL:      webhookUrl,
		UserInfo: args.UserInfo,
	}
}

func (w *WebhookClient) Send(args *WebhookPayload) error {
	err := w.validatePayload(args)
	if err != nil {
		return err
	}

	var r goaxios.GoAxios
	if args.Files != nil {
		var files []goaxios.FormFile

		for _, file := range args.Files {
			files = append(files, goaxios.FormFile{
				Key:    file.Name,
				Name:   file.Name,
				Handle: file.Reader,
			})
		}

		payload := new(bytes.Buffer)
		err = json.NewEncoder(payload).Encode(args)
		if err != nil {
			return err
		}

		r = goaxios.GoAxios{
			Method: "POST",
			Url:    w.URL,
			Form: &goaxios.Form{
				Files: files,
				Data: []goaxios.FormData{
					{
						Key:   "payload_json",
						Value: payload.String(),
					},
				},
			},
		}
	} else {
		r = goaxios.GoAxios{
			Method: "POST",
			Url:    w.URL,
			Body:   args,
		}
	}

	res := r.RunRest()
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (client *WebhookClient) validatePayload(args *WebhookPayload) error {
	var content string
	var attachments []DiscordAttachments
	var embeds []DiscordEmbed

	if args.Content != nil {
		content = *args.Content
	}

	if len(args.Attachments) > 0 {
		for _, attachment := range args.Attachments {
			attachments = append(attachments, *attachment)
		}
	}

	if len(args.Embeds) > 0 {
		if len(args.Embeds) > 10 {
			return fmt.Errorf("DiscordEmbed must be less than 10.")
		}

		for i, embed := range args.Embeds {
			err := validateEmbed(*embed, i)

			if err != nil {
				return err
			}

			embeds = append(embeds, *embed)
		}
	}

	if len(args.Files) > 10 {
		return fmt.Errorf("WebhookFile must be less than 10.")
	}

	if content == "" && attachments == nil && embeds == nil && args.Files == nil {
		return fmt.Errorf("One of content, attachments, or embeds must be set.")
	}

	if client.UserInfo != nil {
		if args.Username == nil && client.UserInfo.Username != nil {
			args.Username = client.UserInfo.Username
		}

		if args.AvatarURL == nil && client.UserInfo.AvatarURL != nil {
			args.AvatarURL = client.UserInfo.AvatarURL
		}
	}

	return nil
}

func validateEmbed(embed DiscordEmbed, index int) error {
	if len(embed.Fields) > 25 {
		return fmt.Errorf("DiscordEmbed[%d] has too many fields", index)
	}

	return nil
}

func createWebhookURL(clientId string, token string) string {
	url := url.URL{
		Scheme: "https",
		Host:   "discord.com",
		Path:   "/api/webhooks/" + clientId + "/" + token,
	}

	return url.String()
}
