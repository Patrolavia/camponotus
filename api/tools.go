// This file is part of Camponotus
// Camponotus is free software: see LICENSE.txt for more details.

package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sort"
)

// LongPollFetcher fetches messages using long polling
type LongPollFetcher struct {
	Message            chan *Message
	EditedMessage      chan *Message
	InlineQuery        chan *InlineQuery
	ChosenInlineResult chan *ChosenInlineResult
	CallbackQuery      chan *CallbackQuery
	API                *API
}

type byUpdateID []Update

func (b byUpdateID) Len() int           { return len(b) }
func (b byUpdateID) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b byUpdateID) Less(i, j int) bool { return b[i].ID < b[j].ID }

// Fetch fetches messages by long polling in loops
func (l *LongPollFetcher) Fetch(limit, timeout int) error {
	offset := 0
	if limit < 1 {
		limit = 1
	}
	if timeout < 0 {
		timeout = 0
	}
	for {
		data, err := l.API.GetUpdates(offset, limit, timeout)
		if err != nil {
			return err
		}

		sort.Sort(byUpdateID(data))

		for _, u := range data {
			if u.ID >= offset {
				offset = u.ID + 1
			}
			switch {
			case u.Message != nil && l.Message != nil:
				l.Message <- u.Message
			case u.EditedMessage != nil && l.EditedMessage != nil:
				l.EditedMessage <- u.EditedMessage
			case u.InlineQuery != nil && l.InlineQuery != nil:
				l.InlineQuery <- u.InlineQuery
			case u.ChosenInlineResult != nil && l.ChosenInlineResult != nil:
				l.ChosenInlineResult <- u.ChosenInlineResult
			case u.CallbackQuery != nil && l.CallbackQuery != nil:
				l.CallbackQuery <- u.CallbackQuery
			}
		}
	}
}

// WebhookHandler eases your work to write handler
type WebhookHandler func(w http.ResponseWriter, r *http.Request, update *Update)

func (h WebhookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var u Update
	if err := json.Unmarshal(buf, &u); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	h(w, r, &u)
}
