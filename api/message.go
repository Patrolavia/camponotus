// This file is part of Camponotus
// Camponotus is free software: see LICENSE.txt for more details.

package api

import (
	"io"
	"net/url"
	"strconv"
)

// these are valid parse modes
const (
	TextMode     = ""
	HTMLMode     = "HTML"
	MarkdownMode = "Markdown"
)

// SendMessage maps to https://core.telegram.org/bots/api#sendmessage
func (a *API) SendMessage(chat, text, mode string, noPreview, silent bool, reply int, markup []byte) ([]byte, error) {
	params := url.Values{}

	params.Set("chat_id", chat)
	params.Set("text", text)
	optStr(params, "parse_mode", mode)
	optBool(params, "disable_web_page_preview", noPreview)
	optBool(params, "disable_notification", silent)
	optInt(params, "reply_to_message_id", reply)
	optJSON(params, "reply_markup", markup)

	return a.call("sendMessage", params)
}

// ForwardMessage maps to https://core.telegram.org/bots/api#forwardmessage
func (a *API) ForwardMessage(to, from string, silent bool, message int) ([]byte, error) {
	params := url.Values{}

	params.Set("chat_id", to)
	params.Set("from_chat_id", from)
	optBool(params, "disable_notification", silent)
	params.Set("message_id", strconv.Itoa(message))

	return a.call("forwardMessage", params)
}

// SendPhoto sends cached photo, maps to https://core.telegram.org/bots/api#sendphoto
func (a *API) SendPhoto(chat, photo, caption string, silent bool, reply int, markup []byte) ([]byte, error) {
	params := url.Values{}

	params.Set("chat_id", chat)
	params.Set("photo", photo)
	optStr(params, "caption", caption)
	optBool(params, "disable_notification", silent)
	optInt(params, "reply_to_message_id", reply)
	optJSON(params, "reply_markup", markup)

	return a.call("sendPhoto", params)
}

// UploadPhoto uploads a photo and send it, maps to https://core.telegram.org/bots/api#sendphoto
func (a *API) UploadPhoto(chat string, photo io.Reader, caption string, silent bool, reply int, markup []byte) ([]byte, error) {
	params := url.Values{}

	params.Set("chat_id", chat)
	optStr(params, "caption", caption)
	optBool(params, "disable_notification", silent)
	optInt(params, "reply_to_message_id", reply)
	optJSON(params, "reply_markup", markup)

	return a.upload("sendPhoto", params, "photo", photo)
}

// SendAudio sends cached audio, maps to https://core.telegram.org/bots/api#sendaudio
func (a *API) SendAudio(chat, audio string, duration int, performer, title string, silent bool, reply int, markup []byte) ([]byte, error) {
	params := url.Values{}

	params.Set("chat_id", chat)
	params.Set("audio", audio)
	optInt(params, "duration", duration)
	optStr(params, "performer", performer)
	optStr(params, "title", title)
	optBool(params, "disable_notification", silent)
	optInt(params, "reply_to_message_id", reply)
	optJSON(params, "reply_markup", markup)

	return a.call("sendAudio", params)
}

// UploadAudio sends cached audio, maps to https://core.telegram.org/bots/api#sendaudio
func (a *API) UploadAudio(chat string, audio io.Reader, duration int, performer, title string, silent bool, reply int, markup []byte) ([]byte, error) {
	params := url.Values{}

	params.Set("chat_id", chat)
	optInt(params, "duration", duration)
	optStr(params, "performer", performer)
	optStr(params, "title", title)
	optBool(params, "disable_notification", silent)
	optInt(params, "reply_to_message_id", reply)
	optJSON(params, "reply_markup", markup)

	return a.upload("sendAudio", params, "audio", audio)
}

// SendDocument sends cached document, maps to https://core.telegram.org/bots/api#senddocument
func (a *API) SendDocument(chat, document, caption string, silent bool, reply int, markup []byte) ([]byte, error) {
	params := url.Values{}

	params.Set("chat_id", chat)
	params.Set("document", document)
	optStr(params, "caption", caption)
	optBool(params, "disable_notification", silent)
	optInt(params, "reply_to_message_id", reply)
	optJSON(params, "reply_markup", markup)

	return a.call("sendDocument", params)
}

// UploadDocument uploads a document and send it, maps to https://core.telegram.org/bots/api#senddocument
func (a *API) UploadDocument(chat string, document io.Reader, caption string, silent bool, reply int, markup []byte) ([]byte, error) {
	params := url.Values{}

	params.Set("chat_id", chat)
	optStr(params, "caption", caption)
	optBool(params, "disable_notification", silent)
	optInt(params, "reply_to_message_id", reply)
	optJSON(params, "reply_markup", markup)

	return a.upload("sendDocument", params, "document", document)
}

// SendSticker sends cached sticker, maps to https://core.telegram.org/bots/api#sendsticker
func (a *API) SendSticker(chat, sticker, caption string, silent bool, reply int, markup []byte) ([]byte, error) {
	params := url.Values{}

	params.Set("chat_id", chat)
	params.Set("sticker", sticker)
	optStr(params, "caption", caption)
	optBool(params, "disable_notification", silent)
	optInt(params, "reply_to_message_id", reply)
	optJSON(params, "reply_markup", markup)

	return a.call("sendSticker", params)
}

// UploadSticker uploads a sticker and send it, maps to https://core.telegram.org/bots/api#sendsticker
func (a *API) UploadSticker(chat string, sticker io.Reader, caption string, silent bool, reply int, markup []byte) ([]byte, error) {
	params := url.Values{}

	params.Set("chat_id", chat)
	optStr(params, "caption", caption)
	optBool(params, "disable_notification", silent)
	optInt(params, "reply_to_message_id", reply)
	optJSON(params, "reply_markup", markup)

	return a.upload("sendSticker", params, "sticker", sticker)
}

// SendVideo sends cached video, maps to https://core.telegram.org/bots/api#sendvideo
func (a *API) SendVideo(chat, video string, duration, width, height int, caption string, silent bool, reply int, markup []byte) ([]byte, error) {
	params := url.Values{}

	params.Set("chat_id", chat)
	params.Set("video", video)
	optInt(params, "duration", duration)
	optInt(params, "width", width)
	optInt(params, "height", height)
	optStr(params, "caption", caption)
	optBool(params, "disable_notification", silent)
	optInt(params, "reply_to_message_id", reply)
	optJSON(params, "reply_markup", markup)

	return a.call("sendVideo", params)
}

// UploadVideo uploads a video and send it, maps to https://core.telegram.org/bots/api#sendvideo
func (a *API) UploadVideo(chat string, video io.Reader, duration, width, height int, caption string, silent bool, reply int, markup []byte) ([]byte, error) {
	params := url.Values{}

	params.Set("chat_id", chat)
	optInt(params, "duration", duration)
	optInt(params, "width", width)
	optInt(params, "height", height)
	optStr(params, "caption", caption)
	optBool(params, "disable_notification", silent)
	optInt(params, "reply_to_message_id", reply)
	optJSON(params, "reply_markup", markup)

	return a.upload("sendVideo", params, "video", video)
}

// SendVoice sends cached voice, maps to https://core.telegram.org/bots/api#sendvoice
func (a *API) SendVoice(chat, voice string, duration int, silent bool, reply int, markup []byte) ([]byte, error) {
	params := url.Values{}

	params.Set("chat_id", chat)
	params.Set("voice", voice)
	optInt(params, "duration", duration)
	optBool(params, "disable_notification", silent)
	optInt(params, "reply_to_message_id", reply)
	optJSON(params, "reply_markup", markup)

	return a.call("sendVoice", params)
}

// UploadVoice uploads a voice and send it, maps to https://core.telegram.org/bots/api#sendvoice
func (a *API) UploadVoice(chat string, voice io.Reader, duration int, silent bool, reply int, markup []byte) ([]byte, error) {
	params := url.Values{}

	params.Set("chat_id", chat)
	optInt(params, "duration", duration)
	optBool(params, "disable_notification", silent)
	optInt(params, "reply_to_message_id", reply)
	optJSON(params, "reply_markup", markup)

	return a.upload("sendVoice", params, "voice", voice)
}

// SendLocation maps to https://core.telegram.org/bots/api#sendlocation
func (a *API) SendLocation(chat string, lat, lng float64, silent bool, reply int, markup []byte) ([]byte, error) {
	params := url.Values{}

	params.Set("chat_id", chat)
	params.Set("latitude", strconv.FormatFloat(lat, 'f', -1, 64))
	params.Set("longitude", strconv.FormatFloat(lng, 'f', -1, 64))
	optBool(params, "disable_notification", silent)
	optInt(params, "reply_to_message_id", reply)
	optJSON(params, "reply_markup", markup)

	return a.call("sendLocation", params)
}

// SendVenue maps to https://core.telegram.org/bots/api#sendvenue
func (a *API) SendVenue(chat string, lat, lng float64, title, addr, foursq string, silent bool, reply int, markup []byte) ([]byte, error) {
	params := url.Values{}

	params.Set("chat_id", chat)
	params.Set("latitude", strconv.FormatFloat(lat, 'f', -1, 64))
	params.Set("longitude", strconv.FormatFloat(lng, 'f', -1, 64))
	params.Set("title", title)
	params.Set("address", addr)
	optStr(params, "foursquare_id", foursq)
	optBool(params, "disable_notification", silent)
	optInt(params, "reply_to_message_id", reply)
	optJSON(params, "reply_markup", markup)

	return a.call("sendVenue", params)
}

// SendContact maps to https://core.telegram.org/bots/api#sendcontact
func (a *API) SendContact(chat, phone, firstName, lastName string, silent bool, reply int, markup []byte) ([]byte, error) {
	params := url.Values{}

	params.Set("chat_id", chat)
	params.Set("phone_number", phone)
	params.Set("first_name", firstName)
	optStr(params, "last_name", lastName)
	optBool(params, "disable_notification", silent)
	optInt(params, "reply_to_message_id", reply)
	optJSON(params, "reply_markup", markup)

	return a.call("sendContact", params)
}

// SendChatAction maps to https://core.telegram.org/bots/api#sendchataction
func (a *API) SendChatAction(chat, action string) ([]byte, error) {
	params := url.Values{}

	params.Set("chat_id", chat)
	params.Set("action", action)

	return a.call("sendChatAction", params)
}
