// This file is part of Camponotus
// Camponotus is free software: see LICENSE.txt for more details.

package raw

import (
	"net/url"
	"strconv"
)

// EditText edits text message, maps to https://core.telegram.org/bots/api#editmessagetext
func (a *API) EditText(chat string, msg int, text, mode string, noPreview bool, markup []byte) ([]byte, error) {
	params := url.Values{}

	params.Set("chat_id", chat)
	params.Set("message_id", strconv.Itoa(msg))
	params.Set("text", text)
	optStr(params, "parse_mode", mode)
	optBool(params, "disable_Web_page_preview", noPreview)
	optJSON(params, "reply_markup", markup)

	return a.call("editMessageText", params)
}

// EditInlineText edits inline text message, maps to https://core.telegram.org/bots/api#editmessagetext
func (a *API) EditInlineText(msg, text, mode string, noPreview bool, markup []byte) ([]byte, error) {
	params := url.Values{}

	params.Set("inline_message_id", msg)
	params.Set("text", text)
	optStr(params, "parse_mode", mode)
	optBool(params, "disable_Web_page_preview", noPreview)
	optJSON(params, "reply_markup", markup)

	return a.call("editMessageText", params)
}

// EditCaption edits caption message, maps to https://core.telegram.org/bots/api#editmessagecaption
func (a *API) EditCaption(chat string, msg int, caption, mode string, noPreview bool, markup []byte) ([]byte, error) {
	params := url.Values{}

	params.Set("chat_id", chat)
	params.Set("message_id", strconv.Itoa(msg))
	params.Set("caption", caption)
	optStr(params, "parse_mode", mode)
	optBool(params, "disable_Web_page_preview", noPreview)
	optJSON(params, "reply_markup", markup)

	return a.call("editMessageCaption", params)
}

// EditInlineCaption edits inline caption message, maps to https://core.telegram.org/bots/api#editmessagecaption
func (a *API) EditInlineCaption(msg, caption, mode string, noPreview bool, markup []byte) ([]byte, error) {
	params := url.Values{}

	params.Set("inline_message_id", msg)
	params.Set("caption", caption)
	optStr(params, "parse_mode", mode)
	optBool(params, "disable_Web_page_preview", noPreview)
	optJSON(params, "reply_markup", markup)

	return a.call("editMessageCaption", params)
}

// EditMarkup edits text message, maps to https://core.telegram.org/bots/api#editmessagetext
func (a *API) EditMarkup(chat string, msg int, markup []byte) ([]byte, error) {
	params := url.Values{}

	params.Set("chat_id", chat)
	params.Set("message_id", strconv.Itoa(msg))
	params.Set("reply_markup", string(markup))

	return a.call("editMessageReplyMarkup", params)
}

// EditInlineMarkup edits text message, maps to https://core.telegram.org/bots/api#editmessagetext
func (a *API) EditInlineMarkup(msg string, markup []byte) ([]byte, error) {
	params := url.Values{}

	params.Set("inline_message_id", msg)
	params.Set("reply_markup", string(markup))

	return a.call("editMessageReplyMarkup", params)
}
