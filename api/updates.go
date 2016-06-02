// This file is part of Camponotus
// Camponotus is free software: see LICENSE.txt for more details.

package api

import (
	"io"
	"net/url"
)

// GetUpdates maps to https://core.telegram.org/bots/api#getupdates
func (a *API) GetUpdates(offset, limit, timeout int) ([]byte, error) {
	params := url.Values{}
	optInt(params, "offset", offset)
	optInt(params, "limit", limit)
	optInt(params, "timeout", timeout)

	return a.call("getUpdates", params)
}

// SetWebhook maps to https://core.telegram.org/bots/api#setwebhook
func (a *API) SetWebhook(cb string, certificate io.Reader) ([]byte, error) {
	params := url.Values{}
	optStr(params, "url", cb)
	if certificate != nil {
		return a.upload("setWebhook", params, "certificate", certificate)
	}

	return a.call("setWebhook", params)
}
