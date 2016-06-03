// This file is part of Camponotus
// Camponotus is free software: see LICENSE.txt for more details.

package telegram

import (
	"fmt"
	"io"
	"net/url"
	"strconv"
)

// types for api method result

type result interface {
	OK() bool
}

type boolResult struct {
	Ok bool `json:"ok"`
}

func (r boolResult) OK() bool {
	return r.Ok
}

type intResult struct {
	boolResult
	Result int64 `json:"result"`
}

type updateResult struct {
	boolResult
	Updates []Update `json:"result"`
}

type userResult struct {
	boolResult
	User *Victim `json:"result"`
}

type msgResult struct {
	boolResult
	Message *Message `json:"result"`
}

type profilePhotoResult struct {
	boolResult
	UserProfilePhotos *UserProfilePhotos `json:"result"`
}

type fileResult struct {
	boolResult
	File *File `json:"result"`
}

type memberResult struct {
	boolResult
	Member *ChatMember `json:"result"`
}

type membersResult struct {
	boolResult
	Members []ChatMember `json:"result"`
}

// ErrNotOK means Telegram server returns fail on your method call
type ErrNotOK struct {
	Method string
	Params url.Values
	Field  string
	Data   io.Reader
}

func (e ErrNotOK) Error() string {
	msg := "is"
	if e.Data != nil {
		msg = "is not"
	}
	return fmt.Sprintf(
		"Telegram server returns fail on method %s, param %v, field %s, reader %s nil",
		e.Method,
		e.Params,
		e.Field,
		msg,
	)
}

// helpers

func optStr(params url.Values, key, val string) {
	if val != "" {
		params.Set(key, val)
	}
}

func optInt(params url.Values, key string, val int) {
	if val != 0 {
		params.Set(key, strconv.Itoa(val))
	}
}

func optBool(params url.Values, key string, val bool) {
	if val {
		params.Set(key, "true")
	}
}

func optFloat(params url.Values, key string, val float64) {
	if val != 0.0 {
		params.Set(key, strconv.FormatFloat(val, 'f', -1, 64))
	}
}

func optJSON(params url.Values, key string, val []byte) {
	if val != nil && len(val) > 0 {
		params.Set(key, string(val))
	}
}

// Options abstracts some commonly used options for api methods
//
// Not every option supported by every api method.
type Options struct {
	ParseMode   string
	NoPreview   bool
	Silent      bool
	ReplyID     int
	ReplyMarkup ReplyMarkup
}
