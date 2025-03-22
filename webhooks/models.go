package webhooks

import (
	"io"
)

type NewWebhookArgs struct {
	URL      *string
	ClientID *string
	Secret   *string
	UserInfo *WebhookUserInfo
}

type WebhookClient struct {
	URL      string
	UserInfo *WebhookUserInfo
}

type WebhookUserInfo struct {
	Username  *string
	AvatarURL *string
}

type WebhookFile struct {
	Name   string
	Reader io.ReadCloser
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
	DiscordEmbedGenericMedia
}

type DiscordEmbedThumbnail struct {
	DiscordEmbedGenericMedia
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
