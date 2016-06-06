// This file is part of Camponotus
// Camponotus is free software: see LICENSE.txt for more details.

package telegram

// InlineQuery represents an incoming inline query.
// When the user sends an empty query, your bot could return some default or trending results.
type InlineQuery struct {
	ID       string    `json:"id"`
	From     *Victim   `json:"from"`
	Location *Location `json:"location,omitempty"`
	Query    string    `json:"query"`
	Offset   string    `json:"offset"`
}

// InlineQueryResult holds methods for all kinds of inline query results
type InlineQueryResult interface {
	ValidateType() // validates the result type
}

// AbstractInlineQueryResult holds common fields for inline query results
type AbstractInlineQueryResult struct {
	Type                string                `json:"type"`
	ID                  string                `json:"id"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
	Title               string                `json:"title,omitempty"`
	Description         string                `json:"description,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	ThumbURL            string                `json:"thumb_url,omitempty"`
}

// InlineQueryResultArticle represents a link to an article or web page.
type InlineQueryResultArticle struct {
	AbstractInlineQueryResult
	URL         string `json:"url,omitempty"`
	HideURL     bool   `json:"hide_url,omitempty"`
	ThumbWidth  int    `json:"thumb_width,omitempty"`
	ThumbHeight int    `json:"thumb_height,omitempty"`
}

// ValidateType validate and rewrite result type
func (r *InlineQueryResultArticle) ValidateType() {
	r.AbstractInlineQueryResult.Type = "article"
}

// InlineQueryResultPhoto represents a link to a photo.
type InlineQueryResultPhoto struct {
	AbstractInlineQueryResult
	URL     string `json:"photo_url,omitempty"`
	FileID  string `json:"photo_file_id,omitempty"`
	Width   int    `json:"photo_width,omitempty"`
	Height  int    `json:"photo_height,omitempty"`
	Caption string `json:"caption,omitempty"`
}

// ValidateType validate and rewrite result type
func (r *InlineQueryResultPhoto) ValidateType() {
	r.AbstractInlineQueryResult.Type = "photo"
}

// InlineQueryResultGif represents a link to an animated GIF file.
type InlineQueryResultGif struct {
	AbstractInlineQueryResult
	URL     string `json:"gif_url,omitempty"`
	FileID  string `json:"gif_file_id,omitempty"`
	Width   int    `json:"git_width,omitempty"`
	Height  int    `json:"gif_height,omitempty"`
	Caption string `json:"caption,omitempty"`
}

// ValidateType validate and rewrite result type
func (r *InlineQueryResultGif) ValidateType() {
	r.AbstractInlineQueryResult.Type = "gif"
}

// InlineQueryResultMpeg4Gif represents a link to a video animation.
type InlineQueryResultMpeg4Gif struct {
	AbstractInlineQueryResult
	URL     string `json:"mpeg4_url,omitempty"`
	FileID  string `json:"mpeg4_file_id,omitempty"`
	Width   int    `json:"mpeg4_width,omitempty"`
	Height  int    `json:"mpeg4_height,omitempty"`
	Caption string `json:"caption,omitempty"`
}

// ValidateType validate and rewrite result type
func (r *InlineQueryResultMpeg4Gif) ValidateType() {
	r.AbstractInlineQueryResult.Type = "mpeg4_gif"
}

// InlineQueryResultVideo represnets a link to a page containing an embedded video player or a video file.
type InlineQueryResultVideo struct {
	AbstractInlineQueryResult
	URL      string `json:"video_url,omitempty"`
	FileID   string `json:"video_file_id,omitempty"`
	MimeType string `json:"mime_type"`
	Caption  string `json:"caption,omitempty"`
	Width    int    `json:"video_width,omitempty"`
	Height   int    `json:"video_height,omitempty"`
}

// ValidateType validate and rewrite result type
func (r *InlineQueryResultVideo) ValidateType() {
	r.AbstractInlineQueryResult.Type = "video"
}

// InlineQueryResultAudio represents a link to an mp3 audio file.
type InlineQueryResultAudio struct {
	AbstractInlineQueryResult
	URL       string `json:"audio_url,omitempty"`
	FileID    string `json:"audio_file_id,omitempty"`
	Performer string `json:"performer,omitempty"`
	Duration  int    `json:"audio_duration,omitempty"`
}

// ValidateType validate and rewrite result type
func (r *InlineQueryResultAudio) ValidateType() {
	r.AbstractInlineQueryResult.Type = "audio"
}

// InlineQueryResultVoice represents a link to a voice recording in an .off container encoded with OPUS.
type InlineQueryResultVoice struct {
	AbstractInlineQueryResult
	URL      string `json:"voice_url,omitempty"`
	FileID   string `json:"voice_file_id,omitempty"`
	Duration int    `json:"voice_duration,omitempty"`
}

// ValidateType validate and rewrite result type
func (r *InlineQueryResultVoice) ValidateType() {
	r.AbstractInlineQueryResult.Type = "voice"
}

// InlineQueryResultDocument represents a link to a file.
type InlineQueryResultDocument struct {
	AbstractInlineQueryResult
	Caption     string `json:"caption,omitempty"`
	URL         string `json:"document_url,omitempty"`
	FileID      string `json:"document_file_id,omitempty"`
	MimeType    string `json:"mime_type"`
	ThumbWidth  int    `json:"thumb_width,omitempty"`
	ThumbHeight int    `json:"thumb_height,omitempty"`
}

// ValidateType validate and rewrite result type
func (r *InlineQueryResultDocument) ValidateType() {
	r.AbstractInlineQueryResult.Type = "document"
}

// InlineQueryResultLocation represents a location on a map.
type InlineQueryResultLocation struct {
	AbstractInlineQueryResult
	Lat         float64 `json:"latitude"`
	Lng         float64 `json:"longitude"`
	ThumbWidth  int     `json:"thumb_width,omitempty"`
	ThumbHeight int     `json:"thumb_height,omitempty"`
}

// ValidateType validate and rewrite result type
func (r *InlineQueryResultLocation) ValidateType() {
	r.AbstractInlineQueryResult.Type = "location"
}

// InlineQueryResultVenue represents a venue.
type InlineQueryResultVenue struct {
	AbstractInlineQueryResult
	Lat         float64 `json:"latitude"`
	Lng         float64 `json:"longitude"`
	Address     string  `json:"address"`
	Foursquare  string  `json:"foursquare_id,omitempty"`
	ThumbWidth  int     `json:"thumb_width,omitempty"`
	ThumbHeight int     `json:"thumb_height,omitempty"`
}

// ValidateType validate and rewrite result type
func (r *InlineQueryResultVenue) ValidateType() {
	r.AbstractInlineQueryResult.Type = "venue"
}

// InlineQueryResultContact represents a contact with a phone number.
type InlineQueryResultContact struct {
	AbstractInlineQueryResult
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name,omitempty"`
	ThumbWidth  int    `json:"thumb_width,omitempty"`
	ThumbHeight int    `json:"thumb_height,omitempty"`
}

// ValidateType validate and rewrite result type
func (r *InlineQueryResultContact) ValidateType() {
	r.AbstractInlineQueryResult.Type = "contact"
}

// InlineQueryResultSticker represents a link to a sticker stored on the Telegram servers.
//
// This struct maps to https://core.telegram.org/bots/api#inlinequeryresultcachedsticker
type InlineQueryResultSticker struct {
	AbstractInlineQueryResult
	FileID string `json:"sticker_file_id"`
}

// ValidateType validate and rewrite result type
func (r *InlineQueryResultSticker) ValidateType() {
	r.AbstractInlineQueryResult.Type = "sticker"
}

// InputMessageContent represents the content of a message to be sent as an result of an inline query.
type InputMessageContent map[string]interface{}

// InputTextMessageContent creates an InputMessageContent for text message
func InputTextMessageContent(msg, mode string, noPreview bool) InputMessageContent {
	ret := map[string]interface{}{"message_text": msg}

	if mode != "" {
		ret["parse_mode"] = mode
	}
	if noPreview {
		ret["disable_web_page_preview"] = true
	}

	return InputMessageContent(ret)
}

// InputLocationMessageContent creates an InputMessageContent for location message
func InputLocationMessageContent(lat, lng float64) InputMessageContent {
	return InputMessageContent(map[string]interface{}{
		"latitude":  lat,
		"longitude": lng,
	})
}

// InputVenueMessageContent creates an InputMessageContent for venue message
func InputVenueMessageContent(lat, lng float64, title, addr, foursq string) InputMessageContent {
	ret := map[string]interface{}{
		"latitude":  lat,
		"longitude": lng,
		"title":     title,
		"address":   addr,
	}

	if foursq != "" {
		ret["foursquare_id"] = foursq
	}

	return InputMessageContent(ret)
}

// InputContactMessageContent creates an InputMessageContent for contact message
func InputContactMessageContent(phone, firstName, lastName string) InputMessageContent {
	ret := map[string]interface{}{
		"phone_number": phone,
		"first_name":   firstName,
	}

	if lastName != "" {
		ret["last_name"] = lastName
	}

	return InputMessageContent(ret)
}

// ChosenInlineResult represents a result of an inline query that was chosen by the user and sent to their chat partner.
type ChosenInlineResult struct {
	ID       string    `json:"result_id"`
	From     *Victim   `json:"from"`
	Location *Location `json:"location,omitempty"`
	InlineID string    `json:"inline_message_id,omitempty"`
	Query    string    `json:"query,omitempty"`
}
