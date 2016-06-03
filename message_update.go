// This file is part of Camponotus
// Camponotus is free software: see LICENSE.txt for more details.

package telegram

import (
	"net/url"
	"strconv"
)

// EditText edits text message, maps to https://core.telegram.org/bots/api#editmessagetext
//
// By official documentations, server will return boolean true when editing the message sent by others.
// This method will report json parse error when such situation.
func (a *api) EditText(chat string, msg int, text, mode string, noPreview bool, markup ReplyMarkup) (*Message, error) {
	params := url.Values{}

	params.Set("chat_id", chat)
	params.Set("message_id", strconv.Itoa(msg))
	params.Set("text", text)
	optStr(params, "parse_mode", mode)
	optBool(params, "disable_Web_page_preview", noPreview)

	if markup != nil {
		m, err := markup.Bytes()
		if err != nil {
			return nil, err
		}
		optJSON(params, "reply_markup", m)
	}

	return a.callAndSetMsg("editMessageText", params)
}

// EditInlineText edits inline text message, maps to https://core.telegram.org/bots/api#editmessagetext
//
// By official documentations, server will return boolean true when editing the message sent by others.
// This method will report json parse error when such situation.
func (a *api) EditInlineText(msg, text, mode string, noPreview bool, markup ReplyMarkup) (*Message, error) {
	params := url.Values{}

	params.Set("inline_message_id", msg)
	params.Set("text", text)
	optStr(params, "parse_mode", mode)
	optBool(params, "disable_Web_page_preview", noPreview)

	if markup != nil {
		m, err := markup.Bytes()
		if err != nil {
			return nil, err
		}
		optJSON(params, "reply_markup", m)
	}

	return a.callAndSetMsg("editMessageText", params)
}

// EditCaption edits caption message, maps to https://core.telegram.org/bots/api#editmessagecaption
//
// By official documentations, server will return boolean true when editing the message sent by others.
// This method will report json parse error when such situation.
func (a *api) EditCaption(chat string, msg int, caption, mode string, noPreview bool, markup ReplyMarkup) (*Message, error) {
	params := url.Values{}

	params.Set("chat_id", chat)
	params.Set("message_id", strconv.Itoa(msg))
	params.Set("caption", caption)
	optStr(params, "parse_mode", mode)
	optBool(params, "disable_Web_page_preview", noPreview)

	if markup != nil {
		m, err := markup.Bytes()
		if err != nil {
			return nil, err
		}
		optJSON(params, "reply_markup", m)
	}

	return a.callAndSetMsg("editMessageCaption", params)
}

// EditInlineCaption edits inline caption message, maps to https://core.telegram.org/bots/api#editmessagecaption
//
// By official documentations, server will return boolean true when editing the message sent by others.
// This method will report json parse error when such situation.
func (a *api) EditInlineCaption(msg, caption, mode string, noPreview bool, markup ReplyMarkup) (*Message, error) {
	params := url.Values{}

	params.Set("inline_message_id", msg)
	params.Set("caption", caption)
	optStr(params, "parse_mode", mode)
	optBool(params, "disable_Web_page_preview", noPreview)

	if markup != nil {
		m, err := markup.Bytes()
		if err != nil {
			return nil, err
		}
		optJSON(params, "reply_markup", m)
	}

	return a.callAndSetMsg("editMessageCaption", params)
}

// EditMarkup edits text message, maps to https://core.telegram.org/bots/api#editmessagetext
//
// By official documentations, server will return boolean true when editing the message sent by others.
// This method will report json parse error when such situation.
func (a *api) EditMarkup(chat string, msg int, markup ReplyMarkup) (*Message, error) {
	params := url.Values{}

	params.Set("chat_id", chat)
	params.Set("message_id", strconv.Itoa(msg))

	m, err := markup.Bytes()
	if err != nil {
		return nil, err
	}
	params.Set("reply_markup", string(m))

	return a.callAndSetMsg("editMessageReplyMarkup", params)
}

// EditInlineMarkup edits text message, maps to https://core.telegram.org/bots/api#editmessagetext
//
// By official documentations, server will return boolean true when editing the message sent by others.
// This method will report json parse error when such situation.
func (a *api) EditInlineMarkup(msg string, markup ReplyMarkup) (*Message, error) {
	params := url.Values{}

	params.Set("inline_message_id", msg)

	m, err := markup.Bytes()
	if err != nil {
		return nil, err
	}
	params.Set("reply_markup", string(m))

	return a.callAndSetMsg("editMessageReplyMarkup", params)
}
