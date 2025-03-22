package webhooks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/url"

	"github.com/opensaucerer/goaxios"
)

type WebhookUserInfo struct {
	Username  *string
	AvatarURL *string
}

type WebhookClient struct {
	BaseURL         url.URL
	DefaultUserInfo *WebhookUserInfo
}

// Make sure that the WebhookPayload is valid.
// This should help prevent the request from erroring out due to an invalid request body.
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
			err := embed.validateEmbed(i)

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

	if client.DefaultUserInfo != nil {
		if args.Username == nil && client.DefaultUserInfo.Username != nil {
			args.Username = client.DefaultUserInfo.Username
		}

		if args.AvatarURL == nil && client.DefaultUserInfo.AvatarURL != nil {
			args.AvatarURL = client.DefaultUserInfo.AvatarURL
		}
	}

	return nil
}

type WebhookFile struct {
	Name   string
	Reader io.ReadCloser
}

type WebhookModify struct {
	Name      *string `json:"name"`
	Avatar    *string `json:"avatar"`
	ChannelId string  `json:"channel_id"`
}

type WebhookPayload struct {
	Content         *string                 `json:"content,omitempty"`
	Username        *string                 `json:"username,omitempty"`
	AvatarURL       *string                 `json:"avatar_url,omitempty"`
	TTS             *bool                   `json:"tts,omitempty"`
	Embeds          []*DiscordEmbed         `json:"embeds,omitempty"`
	AllowedMentions *DiscordAllowedMentions `json:"allowed_mentions,omitempty"`
	Attachments     []*DiscordAttachments   `json:"attachments,omitempty"`
	Flags           *int                    `json:"flags,omitempty"`
	ThreadName      *string                 `json:"thread_name,omitempty"`
	AppliedTags     *string                 `json:"applied_tags,omitempty"`
	Files           []*WebhookFile          `json:"-"`
}

// If there are Files attached to the payload, the form data has to be parsed.
//
// Under the hood, this uses goaxios to parse the form data automatically.
func (args *WebhookPayload) parseFormData() (*goaxios.Form, error) {
	var files []goaxios.FormFile

	for _, file := range args.Files {
		files = append(files, goaxios.FormFile{
			Key:    file.Name,
			Name:   file.Name,
			Handle: file.Reader,
		})
	}

	payload := new(bytes.Buffer)
	err := json.NewEncoder(payload).Encode(args)

	if err != nil {
		return nil, err
	}

	return &goaxios.Form{
		Files: files,
		Data: []goaxios.FormData{
			{
				Key:   "payload_json",
				Value: payload.String(),
			},
		},
	}, nil
}

type WebhookPayloadResponse struct {
	Type            int                   `json:"type"`
	Content         string                `json:"content"`
	Mentions        []string              `json:"mentions"`
	MentionRoles    []string              `json:"mention_roles"`
	Attachments     []DiscordAttachments  `json:"attachments"`
	Embeds          []DiscordEmbed        `json:"embeds"`
	Timestamp       string                `json:"timestamp"`
	EditedTimestamp *string               `json:"edited_timestamp"`
	Flags           int                   `json:"flags"`
	Components      []any                 `json:"components"` // TODO: Figure out the type of this
	MessageID       string                `json:"id"`
	ChannelID       string                `json:"channel_id"`
	Author          *DiscordMessageAuthor `json:"author"`
	Pinned          bool                  `json:"pinned"`
	MentionEveryone bool                  `json:"mention_everyone"`
	TTS             bool                  `json:"tts"`
	WebhookID       string                `json:"webhook_id"`
}

type DiscordMessageAuthor struct {
	UserID        string  `json:"id"`
	Username      string  `json:"username"`
	Avatar        *string `json:"avatar"`
	Discriminator string  `json:"discriminator"`
	PublicFlags   int     `json:"public_flags"`
	Bot           bool    `json:"bot"`
	GlobalName    *string `json:"global_name"`
	Clan          *string `json:"clan"`
	PrimaryGuild  *string `json:"primary_guild"`
}

type DiscordEmbed struct {
	Title       *string                `json:"title,omitempty"`
	Description *string                `json:"description,omitempty"`
	URL         *string                `json:"url,omitempty"`
	Timestamp   *string                `json:"timestamp,omitempty"`
	Color       *int                   `json:"color,omitempty"`
	Footer      *DiscordEmbedFooter    `json:"footer,omitempty"`
	Image       *DiscordEmbedThumbnail `json:"image,omitempty"`
	Author      *DiscordEmbedAuthor    `json:"author,omitempty"`
	Fields      []*DiscordEmbedField   `json:"fields,omitempty"`
}

// Validate the structure of the Discord embed based on documentation.
//
// https://discord.com/developers/docs/resources/message#embed-object
func (e DiscordEmbed) validateEmbed(index int) error {
	if len(e.Fields) > 25 {
		return fmt.Errorf("DiscordEmbed[%d] has too many fields", index)
	}

	return nil
}

type DiscordEmbedFooter struct {
	Text         *string `json:"text,omitempty"`
	IconURL      *string `json:"icon_url,omitempty"`
	ProxyIconURL *string `json:"proxy_icon_url,omitempty"`
}

type DiscordEmbedGenericMedia struct {
	URL      *string `json:"url,omitempty"`
	ProxyURL *string `json:"proxy_url,omitempty"`
	Height   *int    `json:"height,omitempty"`
	Width    *int    `json:"width,omitempty"`
}

type DiscordEmbedImage struct {
	URL      *string `json:"url,omitempty"`
	ProxyURL *string `json:"proxy_url,omitempty"`
	Height   *int    `json:"height,omitempty"`
	Width    *int    `json:"width,omitempty"`
}

type DiscordEmbedThumbnail struct {
	URL      *string `json:"url,omitempty"`
	ProxyURL *string `json:"proxy_url,omitempty"`
	Height   *int    `json:"height,omitempty"`
	Width    *int    `json:"width,omitempty"`
}

type DiscordEmbedAuthor struct {
	Name         *string `json:"name,omitempty"`
	URL          *string `json:"url,omitempty"`
	IconURL      *string `json:"icon_url,omitempty"`
	ProxyIconURL *string `json:"proxy_icon_url,omitempty"`
}

type DiscordEmbedField struct {
	Name   string `json:"name,omitempty"`
	Value  string `json:"value,omitempty"`
	Inline *bool  `json:"inline,omitempty"`
}

type DiscordAllowedMentions struct {
	Parse       []*string `json:"parse,omitempty"`
	Roles       []*string `json:"roles,omitempty"`
	Users       []*string `json:"users,omitempty"`
	RepliedUser *bool     `json:"replied_user,omitempty"`
}

type DiscordAttachments struct {
	ID          *string `json:"id,omitempty"`
	Filename    *string `json:"filename,omitempty"`
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	ContentType *string `json:"content_type,omitempty"`
	Size        *int    `json:"size,omitempty"`
	URL         *string `json:"url,omitempty"`
	ProxyURL    *string `json:"proxy_url,omitempty"`
	Height      *int    `json:"height,omitempty"`
	Width       *int    `json:"width,omitempty"`
	Duration    *int    `json:"duration_secs,omitempty"`
	Waveform    *string `json:"waveform,omitempty"`
	Flags       *int    `json:"flags,omitempty"`
}
