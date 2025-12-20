package webhooks

import (
	"encoding/json"
	"io"
)

type Webhook struct {
	ID            string  `json:"id"`
	Type          int     `json:"type"`
	GuildID       *string `json:"guild_id"`
	ChannelID     string  `json:"channel_id"`
	Name          string  `json:"name"`
	Avatar        string  `json:"avatar"`
	ApplicationID *string `json:"application_id"`
}

type ModifyWebhook struct {
	Name      string `json:"name,omitempty"`
	Avatar    string `json:"avatar,omitempty"`
	ChannelID string `json:"channel_id,omitempty"`
}

type WebhookFile struct {
	FileName string
	Reader   io.Reader
}

type EmbedFooter struct {
	Text         string `json:"text,omitempty"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

type EmbedImage struct {
	URL      string `json:"url,omitempty"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
}

type EmbedThumbnail EmbedImage

type EmbedAuthor struct {
	Name         string `json:"name,omitempty"`
	URL          string `json:"url,omitempty"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

type EmbedField struct {
	Name   string `json:"name,omitempty"`
	Value  string `json:"value,omitempty"`
	Inline bool   `json:"inline,omitempty"`
}

type Embed struct {
	Title string `json:"title,omitempty"`
	// Type        string      `json:"type,omitempty"` -- Always rich for Webhooks, field not necessary.
	Description string          `json:"description,omitempty"`
	URL         string          `json:"url,omitempty"`
	Timestamp   string          `json:"timestamp,omitempty"`
	Color       int             `json:"color,omitempty"`
	Footer      *EmbedFooter    `json:"footer,omitempty"`
	Image       *EmbedImage     `json:"image,omitempty"`
	Thumbnail   *EmbedThumbnail `json:"thumbnail,omitempty"`
	Author      *EmbedAuthor    `json:"author,omitempty"`
	Fields      []EmbedField    `json:"fields,omitempty"`
}

type AllowedMentionsParse string

var (
	AllowedMentionsParseUsers    AllowedMentionsParse = "users"
	AllowedMentionsParseRoles    AllowedMentionsParse = "roles"
	AllowedMentionsParseEveryone AllowedMentionsParse = "everyone"
)

type AllowedMentions struct {
	Parse       []AllowedMentionsParse `json:"parse,omitempty"`
	Roles       []string               `json:"roles,omitempty"`
	Users       []string               `json:"users,omitempty"`
	RepliedUser bool                   `json:"replied_user,omitempty"`
}

// Supported message flags for Webhooks.
type MessageFlag int

var (
	MessaegFlagSuppressEmbeds        MessageFlag = 1 << 2
	MessageFlagSuppressNotifications MessageFlag = 1 << 12
	MessageFlagIsComponentsV2        MessageFlag = 1 << 15
)

type PollQuestion struct {
	Text string `json:"text,omitempty"`
	// Emoji interface{} `json:"emoji,omitempty"` -- Webhooks only support the "Text" field in this structure.
}

type PollMedia struct {
	Text string `json:"text,omitempty"`
	// Emoji interface{} `json:"emoji,omitempty"` -- Webhooks only support the "Text" field in this structure.
}

type PollAnswer struct {
	AnswerID string `json:"answer_id,omitempty"`
}

type Poll struct {
	Question         *PollQuestion `json:"question,omitempty"`
	Answers          []PollAnswer  `json:"answers,omitempty"`
	Expiry           string        `json:"expiry,omitempty"`
	AllowMultiselect bool          `json:"allow_multiselect,omitempty"`
}

type MessagePayload struct {
	Content         string           `json:"content,omitempty"`
	Username        string           `json:"username,omitempty"`
	AvatarURL       string           `json:"avatar_url,omitempty"`
	TTS             bool             `json:"tts,omitempty"`
	Embeds          []Embed          `json:"embeds,omitempty"`
	AllowedMentions *AllowedMentions `json:"allowed_mentions,omitempty"`
	Flags           *MessageFlag     `json:"flags,omitempty"`
	ThreadName      string           `json:"thread_name,omitempty"`
	AppliedTags     []string         `json:"applied_tags,omitempty"`
	// Components      []interface{}   `json:"components,omitempty"` -- TODO: Add support for components

	Files []WebhookFile `json:"-"`
}

// PayloadJSON will return JSON encoded of non-file params
func (p *MessagePayload) PayloadJSON() string {
	json, _ := json.Marshal(p)
	return string(json)
}

type EditMessagePayload struct {
	Content         string           `json:"content,omitempty"`
	Embeds          []Embed          `json:"embeds,omitempty"`
	Flags           *MessageFlag     `json:"flags,omitempty"`
	AllowedMentions *AllowedMentions `json:"allowed_mentions,omitempty"`
	// Components      []interface{}   `json:"components,omitempty"` -- TODO: Add support for components

	Files []WebhookFile `json:"-"`
}

// PayloadJSON will return JSON encoded of non-file params
func (p *EditMessagePayload) PayloadJSON() string {
	json, _ := json.Marshal(p)
	return string(json)
}
