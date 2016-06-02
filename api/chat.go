// This file is part of Camponotus
// Camponotus is free software: see LICENSE.txt for more details.

package api

import (
	"net/url"
	"strconv"
)

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

// GetChat maps to https://core.telegram.org/bots/api#getchat
func (a *API) GetChat(chat string) ([]byte, error) {
	params := url.Values{}

	params.Set("chat_id", chat)

	return a.call("getChat", params)
}

// GetChatAdministrators maps to https://core.telegram.org/bots/api#getchatadministrators
func (a *API) GetChatAdministrators(chat string) ([]byte, error) {
	params := url.Values{}

	params.Set("chat_id", chat)

	return a.call("getChatAdministrators", params)
}

// GetChatMembersCount maps to https://core.telegram.org/bots/api#getchatmemberscount
func (a *API) GetChatMembersCount(chat string) ([]byte, error) {
	params := url.Values{}

	params.Set("chat_id", chat)

	return a.call("getChatMembersCount", params)
}

// GetChatMember maps to https://core.telegram.org/bots/api#getchatmember
func (a *API) GetChatMember(chat string) ([]byte, error) {
	params := url.Values{}

	params.Set("chat_id", chat)

	return a.call("getChatMember", params)
}
