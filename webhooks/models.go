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
