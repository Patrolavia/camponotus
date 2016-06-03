// This file is part of Camponotus
// Camponotus is free software: see LICENSE.txt for more details.

package telegram

import (
	"io"
	"net/url"
)

// GetUpdates maps to https://core.telegram.org/bots/api#getupdates
func (a *api) GetUpdates(offset, limit, timeout int) ([]Update, error) {
	params := url.Values{}
	optInt(params, "offset", offset)
	optInt(params, "limit", limit)
	optInt(params, "timeout", timeout)

	var u updateResult
	err := a.callAndSet("getUpdates", params, &u)
	return u.Updates, err
}

// SetWebhook maps to https://core.telegram.org/bots/api#setwebhook
func (a *api) SetWebhook(cb string, certificate io.Reader) error {
	params := url.Values{}
	optStr(params, "url", cb)
	var r boolResult
	var err error
	if certificate != nil {
		err = a.uploadAndSet("setWebhook", params, "certificate", certificate, &r)
	} else {
		err = a.callAndSet("setWebhook", params, &r)
	}

	if err == nil && !r.Ok {
		err = &ErrNotOK{"setWebhook", params, "certificate", certificate}
	}
	return err
}
