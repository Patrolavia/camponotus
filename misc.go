// This file is part of Camponotus
// Camponotus is free software: see LICENSE.txt for more details.

package telegram

import (
	"net/url"
	"strconv"
)

// GetMe maps to https://core.telegram.org/bots/api#getme
func (a *api) GetMe() (*Victim, error) {
	var v userResult
	err := a.callAndSet("getMe", url.Values{}, &v)
	return v.User, err
}

// GetUserProfilePhotos maps to https://core.telegram.org/bots/api#getuserprofilephotos
func (a *api) GetUserProfilePhotos(user, offset, limit int) (*UserProfilePhotos, error) {
	params := url.Values{}

	params.Set("user_id", strconv.Itoa(user))
	optInt(params, "offset", offset)
	optInt(params, "limit", limit)

	var r profilePhotoResult
	err := a.callAndSet("getUserProfilePhotos", params, &r)
	return r.UserProfilePhotos, err
}

// GetFile maps to https://core.telegram.org/bots/api#getfile
func (a *api) GetFile(file string) (*File, error) {
	params := url.Values{}

	params.Set("file_id", file)

	var r fileResult
	err := a.callAndSet("getFile", params, &r)
	return r.File, err
}

// AnswerCallbackQuery maps to https://core.telegram.org/bots/api#answercallbackquery
func (a *api) AnswerCallbackQuery(query, text string, alert bool) error {
	params := url.Values{}

	params.Set("callback_query_id", query)
	optStr(params, "text", text)
	optBool(params, "show_alert", alert)

	return a.callAndSet("answerCallbackQuery", params, nil)
}
