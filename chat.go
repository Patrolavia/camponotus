// This file is part of Camponotus
// Camponotus is free software: see LICENSE.txt for more details.

package telegram

import (
	"net/url"
	"strconv"
)

// KickChatMember maps to https://core.telegram.org/bots/api#kickchatmember
func (a *API) KickChatMember(chat string, user int) error {
	params := url.Values{}

	params.Set("chat_id", chat)
	params.Set("user_id", strconv.Itoa(user))

	return a.callAndSet("kickChatMember", params, nil)
}

// LeaveChat maps to https://core.telegram.org/bots/api#leavechat
func (a *API) LeaveChat(chat string) error {
	params := url.Values{}

	params.Set("chat_id", chat)

	return a.callAndSet("leaveChat", params, nil)
}

// UnbanChatMember maps to https://core.telegram.org/bots/api#unbanchatmember
func (a *API) UnbanChatMember(chat string, user int) error {
	params := url.Values{}

	params.Set("chat_id", chat)
	params.Set("user_id", strconv.Itoa(user))

	return a.callAndSet("unbanChatMember", params, nil)
}

// GetChat maps to https://core.telegram.org/bots/api#getchat
func (a *API) GetChat(chat string) (*Victim, error) {
	params := url.Values{}

	params.Set("chat_id", chat)

	var r userResult
	err := a.callAndSet("getChat", params, &r)
	return r.User, err
}

// GetChatAdministrators maps to https://core.telegram.org/bots/api#getchatadministrators
func (a *API) GetChatAdministrators(chat string) ([]ChatMember, error) {
	params := url.Values{}

	params.Set("chat_id", chat)

	var r membersResult
	err := a.callAndSet("getChatAdministrators", params, &r)
	return r.Members, err
}

// GetChatMembersCount maps to https://core.telegram.org/bots/api#getchatmemberscount
func (a *API) GetChatMembersCount(chat string) (int, error) {
	params := url.Values{}

	params.Set("chat_id", chat)

	var r intResult
	err := a.callAndSet("getChatMembersCount", params, &r)
	return int(r.Result), err
}

// GetChatMember maps to https://core.telegram.org/bots/api#getchatmember
func (a *API) GetChatMember(chat string) (*ChatMember, error) {
	params := url.Values{}

	params.Set("chat_id", chat)

	var r memberResult
	err := a.callAndSet("getChatMember", params, &r)
	return r.Member, err
}
