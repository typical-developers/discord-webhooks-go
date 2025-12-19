package webhooks

import (
	"context"
	"net/http"
	"net/url"
)

// TODO: Add support for selecting commented out fields.
// Focusing on webhook functionality for now. It's rather
// tedious to add some of these fields.

type WebhookMessage struct {
	ID        string `json:"id"`
	ChannelID string `json:"channel_id"`
	// Author any       `json:"author"`
	Timestamp       string `json:"timestamp"`
	EditedTimestamp string `json:"edited_timestamp"`
	TTS             bool   `json:"tts"`
	MentionEveryone bool   `json:"mention_everyone"`
	// Mentions []any  `json:"mentions"`
	// MentionRoles []any `json:"mention_roles"`
	// MentionChannels []any `json:"mention_channels"`
	// Attachments []any `json:"attachments"`
	Embeds []Embed `json:"embeds"`
	// Reactions []any `json:"reactions"`
	Pinned    bool        `json:"pinned"`
	WebhookID *string     `json:"webhook_id"`
	Flags     MessageFlag `json:"flags"`

	c *WebhookClient
}

// Edit will edit the current message.
func (m *WebhookMessage) Edit(ctx context.Context, message EditMessagePayload, params *url.Values) (*WebhookMessage, *http.Response, error) {
	return m.c.EditMessage(ctx, m.ID, message, params)
}

// Delete will delete the current message.
func (m *WebhookMessage) Delete(ctx context.Context, params *url.Values) (*http.Response, error) {
	return m.c.DeleteMessage(ctx, m.ID, params)
}
