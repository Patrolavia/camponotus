// This file is part of Camponotus
// Camponotus is free software: see LICENSE.txt for more details.

package telegram

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
func (a *api) SendMessage(chat, text string, opts *Options) (*Message, error) {
	params := url.Values{}

	params.Set("chat_id", chat)
	params.Set("text", text)
	if opts != nil {
		optStr(params, "parse_mode", opts.ParseMode)
		optBool(params, "disable_web_page_preview", opts.NoPreview)
		optBool(params, "disable_notification", opts.Silent)
		optInt(params, "reply_to_message_id", opts.ReplyID)

		if markup := opts.ReplyMarkup; markup != nil {
			m, err := markup.Bytes()
			if err != nil {
				return nil, err
			}
			optJSON(params, "reply_markup", m)
		}
	}

	return a.callAndSetMsg("sendMessage", params)
}

// ForwardMessage maps to https://core.telegram.org/bots/api#forwardmessage
func (a *api) ForwardMessage(to, from string, silent bool, message int) (*Message, error) {
	params := url.Values{}

	params.Set("chat_id", to)
	params.Set("from_chat_id", from)
	optBool(params, "disable_notification", silent)
	params.Set("message_id", strconv.Itoa(message))

	return a.callAndSetMsg("forwardMessage", params)
}

// SendPhoto sends cached photo, maps to https://core.telegram.org/bots/api#sendphoto
func (a *api) SendPhoto(chat, photo, caption string, opts *Options) (*Message, error) {
	params := url.Values{}

	params.Set("chat_id", chat)
	params.Set("photo", photo)
	optStr(params, "caption", caption)

	if opts != nil {
		optBool(params, "disable_notification", opts.Silent)
		optInt(params, "reply_to_message_id", opts.ReplyID)

		if markup := opts.ReplyMarkup; markup != nil {
			m, err := markup.Bytes()
			if err != nil {
				return nil, err
			}
			optJSON(params, "reply_markup", m)
		}
	}

	return a.callAndSetMsg("sendPhoto", params)
}

// UploadPhoto uploads a photo and send it, maps to https://core.telegram.org/bots/api#sendphoto
func (a *api) UploadPhoto(chat string, photo io.Reader, caption string, opts *Options) (*Message, error) {
	params := url.Values{}

	params.Set("chat_id", chat)
	optStr(params, "caption", caption)

	if opts != nil {
		optBool(params, "disable_notification", opts.Silent)
		optInt(params, "reply_to_message_id", opts.ReplyID)

		if markup := opts.ReplyMarkup; markup != nil {
			m, err := markup.Bytes()
			if err != nil {
				return nil, err
			}
			optJSON(params, "reply_markup", m)
		}
	}

	return a.uploadAndSetMsg("sendPhoto", params, "photo", photo)
}

// SendAudio sends cached audio, maps to https://core.telegram.org/bots/api#sendaudio
func (a *api) SendAudio(chat, audio string, duration int, performer, title string, opts *Options) (*Message, error) {
	params := url.Values{}

	params.Set("chat_id", chat)
	params.Set("audio", audio)
	optInt(params, "duration", duration)
	optStr(params, "performer", performer)
	optStr(params, "title", title)

	if opts != nil {
		optBool(params, "disable_notification", opts.Silent)
		optInt(params, "reply_to_message_id", opts.ReplyID)

		if markup := opts.ReplyMarkup; markup != nil {
			m, err := markup.Bytes()
			if err != nil {
				return nil, err
			}
			optJSON(params, "reply_markup", m)
		}
	}

	return a.callAndSetMsg("sendAudio", params)
}

// UploadAudio sends cached audio, maps to https://core.telegram.org/bots/api#sendaudio
func (a *api) UploadAudio(chat string, audio io.Reader, duration int, performer, title string, opts *Options) (*Message, error) {
	params := url.Values{}

	params.Set("chat_id", chat)
	optInt(params, "duration", duration)
	optStr(params, "performer", performer)
	optStr(params, "title", title)

	if opts != nil {
		optBool(params, "disable_notification", opts.Silent)
		optInt(params, "reply_to_message_id", opts.ReplyID)

		if markup := opts.ReplyMarkup; markup != nil {
			m, err := markup.Bytes()
			if err != nil {
				return nil, err
			}
			optJSON(params, "reply_markup", m)
		}
	}

	return a.uploadAndSetMsg("sendAudio", params, "audio", audio)
}

// SendDocument sends cached document, maps to https://core.telegram.org/bots/api#senddocument
func (a *api) SendDocument(chat, document, caption string, opts *Options) (*Message, error) {
	params := url.Values{}

	params.Set("chat_id", chat)
	params.Set("document", document)
	optStr(params, "caption", caption)

	if opts != nil {
		optBool(params, "disable_notification", opts.Silent)
		optInt(params, "reply_to_message_id", opts.ReplyID)

		if markup := opts.ReplyMarkup; markup != nil {
			m, err := markup.Bytes()
			if err != nil {
				return nil, err
			}
			optJSON(params, "reply_markup", m)
		}
	}

	return a.callAndSetMsg("sendDocument", params)
}

// UploadDocument uploads a document and send it, maps to https://core.telegram.org/bots/api#senddocument
func (a *api) UploadDocument(chat string, document io.Reader, caption string, opts *Options) (*Message, error) {
	params := url.Values{}

	params.Set("chat_id", chat)
	optStr(params, "caption", caption)

	if opts != nil {
		optBool(params, "disable_notification", opts.Silent)
		optInt(params, "reply_to_message_id", opts.ReplyID)

		if markup := opts.ReplyMarkup; markup != nil {
			m, err := markup.Bytes()
			if err != nil {
				return nil, err
			}
			optJSON(params, "reply_markup", m)
		}
	}

	return a.uploadAndSetMsg("sendDocument", params, "document", document)
}

// SendSticker sends cached sticker, maps to https://core.telegram.org/bots/api#sendsticker
func (a *api) SendSticker(chat, sticker, caption string, opts *Options) (*Message, error) {
	params := url.Values{}

	params.Set("chat_id", chat)
	params.Set("sticker", sticker)
	optStr(params, "caption", caption)

	if opts != nil {
		optBool(params, "disable_notification", opts.Silent)
		optInt(params, "reply_to_message_id", opts.ReplyID)

		if markup := opts.ReplyMarkup; markup != nil {
			m, err := markup.Bytes()
			if err != nil {
				return nil, err
			}
			optJSON(params, "reply_markup", m)
		}
	}

	return a.callAndSetMsg("sendSticker", params)
}

// UploadSticker uploads a sticker and send it, maps to https://core.telegram.org/bots/api#sendsticker
func (a *api) UploadSticker(chat string, sticker io.Reader, caption string, opts *Options) (*Message, error) {
	params := url.Values{}

	params.Set("chat_id", chat)
	optStr(params, "caption", caption)

	if opts != nil {
		optBool(params, "disable_notification", opts.Silent)
		optInt(params, "reply_to_message_id", opts.ReplyID)

		if markup := opts.ReplyMarkup; markup != nil {
			m, err := markup.Bytes()
			if err != nil {
				return nil, err
			}
			optJSON(params, "reply_markup", m)
		}
	}

	return a.uploadAndSetMsg("sendSticker", params, "sticker", sticker)
}

// SendVideo sends cached video, maps to https://core.telegram.org/bots/api#sendvideo
func (a *api) SendVideo(chat, video string, duration, width, height int, caption string, opts *Options) (*Message, error) {
	params := url.Values{}

	params.Set("chat_id", chat)
	params.Set("video", video)
	optInt(params, "duration", duration)
	optInt(params, "width", width)
	optInt(params, "height", height)
	optStr(params, "caption", caption)

	if opts != nil {
		optBool(params, "disable_notification", opts.Silent)
		optInt(params, "reply_to_message_id", opts.ReplyID)

		if markup := opts.ReplyMarkup; markup != nil {
			m, err := markup.Bytes()
			if err != nil {
				return nil, err
			}
			optJSON(params, "reply_markup", m)
		}
	}

	return a.callAndSetMsg("sendVideo", params)
}

// UploadVideo uploads a video and send it, maps to https://core.telegram.org/bots/api#sendvideo
func (a *api) UploadVideo(chat string, video io.Reader, duration, width, height int, caption string, opts *Options) (*Message, error) {
	params := url.Values{}

	params.Set("chat_id", chat)
	optInt(params, "duration", duration)
	optInt(params, "width", width)
	optInt(params, "height", height)
	optStr(params, "caption", caption)

	if opts != nil {
		optBool(params, "disable_notification", opts.Silent)
		optInt(params, "reply_to_message_id", opts.ReplyID)

		if markup := opts.ReplyMarkup; markup != nil {
			m, err := markup.Bytes()
			if err != nil {
				return nil, err
			}
			optJSON(params, "reply_markup", m)
		}
	}

	return a.uploadAndSetMsg("sendVideo", params, "video", video)
}

// SendVoice sends cached voice, maps to https://core.telegram.org/bots/api#sendvoice
func (a *api) SendVoice(chat, voice string, duration int, opts *Options) (*Message, error) {
	params := url.Values{}

	params.Set("chat_id", chat)
	params.Set("voice", voice)
	optInt(params, "duration", duration)

	if opts != nil {
		optBool(params, "disable_notification", opts.Silent)
		optInt(params, "reply_to_message_id", opts.ReplyID)

		if markup := opts.ReplyMarkup; markup != nil {
			m, err := markup.Bytes()
			if err != nil {
				return nil, err
			}
			optJSON(params, "reply_markup", m)
		}
	}

	return a.callAndSetMsg("sendVoice", params)
}

// UploadVoice uploads a voice and send it, maps to https://core.telegram.org/bots/api#sendvoice
func (a *api) UploadVoice(chat string, voice io.Reader, duration int, opts *Options) (*Message, error) {
	params := url.Values{}

	params.Set("chat_id", chat)
	optInt(params, "duration", duration)

	if opts != nil {
		optBool(params, "disable_notification", opts.Silent)
		optInt(params, "reply_to_message_id", opts.ReplyID)

		if markup := opts.ReplyMarkup; markup != nil {
			m, err := markup.Bytes()
			if err != nil {
				return nil, err
			}
			optJSON(params, "reply_markup", m)
		}
	}

	return a.uploadAndSetMsg("sendVoice", params, "voice", voice)
}

// SendLocation maps to https://core.telegram.org/bots/api#sendlocation
func (a *api) SendLocation(chat string, lat, lng float64, opts *Options) (*Message, error) {
	params := url.Values{}

	params.Set("chat_id", chat)
	params.Set("latitude", strconv.FormatFloat(lat, 'f', -1, 64))
	params.Set("longitude", strconv.FormatFloat(lng, 'f', -1, 64))

	if opts != nil {
		optBool(params, "disable_notification", opts.Silent)
		optInt(params, "reply_to_message_id", opts.ReplyID)

		if markup := opts.ReplyMarkup; markup != nil {
			m, err := markup.Bytes()
			if err != nil {
				return nil, err
			}
			optJSON(params, "reply_markup", m)
		}
	}

	return a.callAndSetMsg("sendLocation", params)
}

// SendVenue maps to https://core.telegram.org/bots/api#sendvenue
func (a *api) SendVenue(chat string, lat, lng float64, title, addr, foursq string, opts *Options) (*Message, error) {
	params := url.Values{}

	params.Set("chat_id", chat)
	params.Set("latitude", strconv.FormatFloat(lat, 'f', -1, 64))
	params.Set("longitude", strconv.FormatFloat(lng, 'f', -1, 64))
	params.Set("title", title)
	params.Set("address", addr)
	optStr(params, "foursquare_id", foursq)

	if opts != nil {
		optBool(params, "disable_notification", opts.Silent)
		optInt(params, "reply_to_message_id", opts.ReplyID)

		if markup := opts.ReplyMarkup; markup != nil {
			m, err := markup.Bytes()
			if err != nil {
				return nil, err
			}
			optJSON(params, "reply_markup", m)
		}
	}

	return a.callAndSetMsg("sendVenue", params)
}

// SendContact maps to https://core.telegram.org/bots/api#sendcontact
func (a *api) SendContact(chat, phone, firstName, lastName string, opts *Options) (*Message, error) {
	params := url.Values{}

	params.Set("chat_id", chat)
	params.Set("phone_number", phone)
	params.Set("first_name", firstName)
	optStr(params, "last_name", lastName)

	if opts != nil {
		optBool(params, "disable_notification", opts.Silent)
		optInt(params, "reply_to_message_id", opts.ReplyID)

		if markup := opts.ReplyMarkup; markup != nil {
			m, err := markup.Bytes()
			if err != nil {
				return nil, err
			}
			optJSON(params, "reply_markup", m)
		}
	}

	return a.callAndSetMsg("sendContact", params)
}

// SendChatAction maps to https://core.telegram.org/bots/api#sendchataction
func (a *api) SendChatAction(chat, action string) error {
	params := url.Values{}

	params.Set("chat_id", chat)
	params.Set("action", action)

	return a.callAndSet("sendChatAction", params, nil)
}
