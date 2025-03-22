package webhooks

import (
	"io"
	"time"
)

// Set the message
func (p *WebhookPayload) SetContent(content string) {
	p.Content = &content
}

// Set the username for the specific webhook message.
func (p *WebhookPayload) SetUsername(username string) {
	p.Username = &username
}

// Set the avatar url for the specific webhook message.
func (p *WebhookPayload) SetAvatarURL(avatarUrl string) {
	p.AvatarURL = &avatarUrl
}

// Set whether or not TTS (text-to-speech) should be played when the message is sent.
func (p *WebhookPayload) SetTTS(tts bool) {
	p.TTS = &tts
}

// Add an attachment to be uploaded to the message.
func (p *WebhookPayload) AddAttachment(name string, reader io.ReadCloser) {
	p.Files = append(p.Files, &WebhookFile{
		Name:   name,
		Reader: reader,
	})
}

// Add a new embed to the message. This returns a pointer to the embed so it can be modified.
func (p *WebhookPayload) AddEmbed() *DiscordEmbed {
	embed := &DiscordEmbed{}
	p.Embeds = append(p.Embeds, embed)
	return embed
}

// Set the title of the embed.
func (e *DiscordEmbed) SetTitle(title string) {
	e.Title = &title
}

// Set the description of the embed.
func (e *DiscordEmbed) SetDescription(description string) {
	e.Description = &description
}

// Set the URL of the embed.
func (e *DiscordEmbed) SetURL(url string) {
	e.URL = &url
}

// Set the timestamp of the embed.
//
// Instead of using a string, using time.Time ensures that the embed timestamp is formatted correctly.
func (e *DiscordEmbed) SetTimestamp(timestamp time.Time) {
	formatTimestamp := timestamp.Format(time.RFC3339)
	e.Timestamp = &formatTimestamp
}

// Set the color of the embed.
func (e *DiscordEmbed) SetColor(color int) {
	e.Color = &color
}

// Set the footer of the embed.
func (e *DiscordEmbed) SetFooter(text string, iconUrl string, proxyIconUrl string) {
	e.Footer = &DiscordEmbedFooter{
		Text:         &text,
		IconURL:      &iconUrl,
		ProxyIconURL: &proxyIconUrl,
	}
}

// Set the image of the embed.
func (e *DiscordEmbed) SetImage(url string) {
	e.Image = &DiscordEmbedThumbnail{
		URL: &url,
	}
}

// Set the author of the embed.
func (e *DiscordEmbed) SetAuthor(name string, url string, iconUrl string) {
	e.Author = &DiscordEmbedAuthor{
		Name:    &name,
		URL:     &url,
		IconURL: &iconUrl,
	}
}

// Add a field to the embed. This returns a pointer to the field so it can be modified.
func (e *DiscordEmbed) AddField() *DiscordEmbedField {
	field := &DiscordEmbedField{}
	e.Fields = append(e.Fields, field)
	return field
}

// Set an existing field.
func (e *DiscordEmbed) SetField(field *DiscordEmbedField) {
	e.Fields = append(e.Fields, field)
}

// Set the name of the field.
func (f *DiscordEmbedField) SetName(name string) {
	f.Name = name
}

// Set the value of the field.
func (f *DiscordEmbedField) SetValue(value string) {
	f.Value = value
}

// Set whether or not the field should be inline with other fields.
func (f *DiscordEmbedField) SetInline(inline bool) {
	f.Inline = &inline
}

// Set the fields of the embed.
func (e *DiscordEmbed) SetFields(fields []*DiscordEmbedField) {
	e.Fields = append(e.Fields, fields...)
}
