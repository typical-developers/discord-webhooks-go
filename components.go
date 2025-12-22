package webhooks

import "encoding/json"

// Not all components are supported by webhooks.
//
// Since this library is purely just for basic webhooks and not application ones,
// some components are not implemented to be used.

type ComponentType int

const (
	ComponentTypeActionRow ComponentType = iota + 1
	ComponentTypeButton
	_
	_
	_
	_
	_
	_
	ComponentTypeSection
	ComponentTypeTextDisplay
	ComponentTypeThumbnail
	ComponentTypeMediaGallery
	ComponentTypeFile
	ComponentTypeSeparator
	_
	_
	ComponentTypeContainer
	_
	_
)

type Component interface {
	json.Marshaler
	Type() ComponentType
}

type ActionRow struct {
	ID         int         `json:"id,omitempty"`
	Components []Component `json:"components"`
}

func (a *ActionRow) Type() ComponentType {
	return ComponentTypeActionRow
}

func (a *ActionRow) MarshalJSON() ([]byte, error) {
	type ar ActionRow

	return json.Marshal(struct {
		ar
		Type ComponentType `json:"type"`
	}{
		ar:   ar(*a),
		Type: a.Type(),
	})
}

type ButtonStyle int

const (
	_ ButtonStyle = iota + 1
	_
	_
	_
	ButtonStyleLink
	_
)

type Button struct {
	ID       int         `json:"id,omitempty"`
	Style    ButtonStyle `json:"style"`
	Label    string      `json:"label,omitempty"`
	CustomID string      `json:"custom_id,omitempty"`
	SKUID    string      `json:"sku_id,omitempty"`
	URL      string      `json:"url,omitempty"`
	Disabled bool        `json:"disabled,omitempty"`
}

func (b *Button) Type() ComponentType {
	return ComponentTypeButton
}

func (b *Button) MarshalJSON() ([]byte, error) {
	type bu Button

	return json.Marshal(struct {
		bu
		Type ComponentType `json:"type"`
	}{
		bu:   bu(*b),
		Type: b.Type(),
	})
}

type Section struct {
	ID         int         `json:"id,omitempty"`
	Components []Component `json:"components"`
	Accessory  Component   `json:"accessory"`
}

func (s *Section) Type() ComponentType {
	return ComponentTypeSection
}

func (s *Section) MarshalJSON() ([]byte, error) {
	type se Section

	return json.Marshal(struct {
		se
		Type ComponentType `json:"type"`
	}{
		se:   se(*s),
		Type: s.Type(),
	})
}

type TextDisplay struct {
	ID      int    `json:"id,omitempty"`
	Content string `json:"content"`
}

func (t *TextDisplay) Type() ComponentType {
	return ComponentTypeTextDisplay
}

func (t *TextDisplay) MarshalJSON() ([]byte, error) {
	type td TextDisplay

	return json.Marshal(struct {
		td
		Type ComponentType `json:"type"`
	}{
		td:   td(*t),
		Type: t.Type(),
	})
}

type UnfurledMediaItem struct {
	URL string `json:"url"`
}

type Thumbnail struct {
	ID          int               `json:"id,omitempty"`
	Description string            `json:"description,omitempty"`
	Media       UnfurledMediaItem `json:"media"`
	Spoiler     bool              `json:"spoiler,omitempty"`
}

func (t *Thumbnail) Type() ComponentType {
	return ComponentTypeThumbnail
}

func (t *Thumbnail) MarshalJSON() ([]byte, error) {
	type th Thumbnail

	return json.Marshal(struct {
		th
		Type ComponentType `json:"type"`
	}{
		th:   th(*t),
		Type: t.Type(),
	})
}

type MediaGalleryItem struct {
	Media       UnfurledMediaItem `json:"media"`
	Description string            `json:"description,omitempty"`
	Spoiler     bool              `json:"spoiler,omitempty"`
}

type MediaGallery struct {
	ID    int                `json:"id,omitempty"`
	Items []MediaGalleryItem `json:"items"`
}

func (m *MediaGallery) Type() ComponentType {
	return ComponentTypeMediaGallery
}

func (m *MediaGallery) MarshalJSON() ([]byte, error) {
	type mg MediaGallery

	return json.Marshal(struct {
		mg
		Type ComponentType `json:"type"`
	}{
		mg:   mg(*m),
		Type: m.Type(),
	})
}

type File struct {
	ID      int               `json:"id,omitempty"`
	File    UnfurledMediaItem `json:"file"`
	Spoiler bool              `json:"spoiler,omitempty"`
}

func (f *File) Type() ComponentType {
	return ComponentTypeFile
}

func (f *File) MarshalJSON() ([]byte, error) {
	type fi File

	return json.Marshal(struct {
		fi
		Type ComponentType `json:"type"`
	}{
		fi:   fi(*f),
		Type: f.Type(),
	})
}

type SeparatorSpacing int

const (
	SeparatorSpacingSmall SeparatorSpacing = iota + 1
	SeparatorSpacingLarge
)

type Separator struct {
	ID      int  `json:"id,omitempty"`
	Divider bool `json:"divider,omitempty"`
	Spacing SeparatorSpacing
}

func (s *Separator) Type() ComponentType {
	return ComponentTypeSeparator
}

func (s *Separator) MarshalJSON() ([]byte, error) {
	type se Separator

	return json.Marshal(struct {
		se
		Type ComponentType `json:"type"`
	}{
		se:   se(*s),
		Type: s.Type(),
	})
}

type Container struct {
	ID          int         `json:"id,omitempty"`
	Components  []Component `json:"components"`
	AccentColor *int        `json:"accent_color,omitempty"`
	Spoiler     bool        `json:"spoiler,omitempty"`
}

func (c *Container) Type() ComponentType {
	return ComponentTypeContainer
}

func (c *Container) MarshalJSON() ([]byte, error) {
	type co Container

	return json.Marshal(struct {
		co
		Type ComponentType `json:"type"`
	}{
		co:   co(*c),
		Type: c.Type(),
	})
}
