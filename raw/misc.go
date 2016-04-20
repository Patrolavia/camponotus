// This file is part of Camponotus
// Camponotus is free software: see LICENSE.txt for more details.

package raw

import (
	"net/url"
	"strconv"
)

// GetMe maps to https://core.telegram.org/bots/api#getme
func (a *API) GetMe() ([]byte, error) {
	return a.call("getMe", url.Values{})
}

// GetUserProfilePhotos maps to https://core.telegram.org/bots/api#getuserprofilephotos
func (a *API) GetUserProfilePhotos(user, offset, limit int) ([]byte, error) {
	params := url.Values{}

	params.Set("user_id", strconv.Itoa(user))
	optInt(params, "offset", offset)
	optInt(params, "limit", limit)

	return a.call("getUserProfilePhotos", params)
}

// GetFile maps to https://core.telegram.org/bots/api#getfile
func (a *API) GetFile(file string) ([]byte, error) {
	params := url.Values{}

	params.Set("file_id", file)

	return a.call("getFile", params)
}

// KickChatMember maps to https://core.telegram.org/bots/api#kickchatmember
func (a *API) KickChatMember(chat string, user int) ([]byte, error) {
	params := url.Values{}

	params.Set("chat_id", chat)
	params.Set("user_id", strconv.Itoa(user))

	return a.call("kickChatMember", params)
}

// UnbanChatMember maps to https://core.telegram.org/bots/api#unbanchatmember
func (a *API) UnbanChatMember(chat string, user int) ([]byte, error) {
	params := url.Values{}

	params.Set("chat_id", chat)
	params.Set("user_id", strconv.Itoa(user))

	return a.call("unbanChatMember", params)
}

// AnswerCallbackQuery maps to https://core.telegram.org/bots/api#answercallbackquery
func (a *API) AnswerCallbackQuery(query, text string, alert bool) ([]byte, error) {
	params := url.Values{}

	params.Set("callback_query_id", query)
	optStr(params, "text", text)
	optBool(params, "show_alert", alert)

	return a.call("answerCallbackQuery", params)
}
