package webhooks

import (
	"io"
	"time"
)

func (p *WebhookPayload) SetContent(content string) {
	p.Content = &content
}

func (p *WebhookPayload) SetUsername(username string) {
	p.Username = &username
}

func (p *WebhookPayload) SetAvatarURL(avatarUrl string) {
	p.AvatarURL = &avatarUrl
}

func (p *WebhookPayload) SetTTS(tts bool) {
	p.TTS = &tts
}

func (p *WebhookPayload) AddAttachment(name string, reader io.ReadCloser) {
	p.Files = append(p.Files, &WebhookFile{
		Name:   name,
		Reader: reader,
	})
}

func (p *WebhookPayload) AddEmbed() *DiscordEmbed {
	embed := &DiscordEmbed{}
	p.Embeds = append(p.Embeds, embed)
	return embed
}

func (e *DiscordEmbed) SetTitle(title string) {
	e.Title = &title
}

func (e *DiscordEmbed) SetDescription(description string) {
	e.Description = &description
}

func (e *DiscordEmbed) SetURL(url string) {
	e.URL = &url
}

func (e *DiscordEmbed) SetTimestamp(timestamp time.Time) {
	formatTimestamp := timestamp.Format(time.RFC3339)
	e.Timestamp = &formatTimestamp
}

func (e *DiscordEmbed) SetColor(color int) {
	e.Color = &color
}

func (e *DiscordEmbed) SetFooter(text string, iconUrl string, proxyIconUrl string) {
	e.Footer = &DiscordEmbedFooter{
		Text:         &text,
		IconURL:      &iconUrl,
		ProxyIconURL: &proxyIconUrl,
	}
}

func (e *DiscordEmbed) SetImage(url string) {
	e.Image = &DiscordEmbedThumbnail{
		URL: &url,
	}
}

func (e *DiscordEmbed) SetAuthor(name string, url string, iconUrl string) {
	e.Author = &DiscordEmbedAuthor{
		Name:    &name,
		URL:     &url,
		IconURL: &iconUrl,
	}
}

func (e *DiscordEmbed) AddField() *DiscordEmbedField {
	field := &DiscordEmbedField{}
	e.Fields = append(e.Fields, field)
	return field
}

func (e *DiscordEmbed) SetField(field *DiscordEmbedField) {
	e.Fields = append(e.Fields, field)
}

func (f *DiscordEmbedField) SetName(name string) {
	f.Name = name
}

func (f *DiscordEmbedField) SetValue(value string) {
	f.Value = value
}

func (f *DiscordEmbedField) SetInline(inline bool) {
	f.Inline = &inline
}

func (e *DiscordEmbed) SetFields(fields []*DiscordEmbedField) {
	e.Fields = append(e.Fields, fields...)
}
