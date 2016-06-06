// This file is part of Camponotus
// Camponotus is free software: see LICENSE.txt for more details.

package telegram

import (
	"encoding/json"
	"net/url"
)

// InlineQueryOptions represents optional parameters for api method answerInlineQuery
type InlineQueryOptions struct {
	CacheTime   int
	Personal    bool
	NextOffset  string
	SwitchPM    string
	SwitchParam string
}

// AnswerInlineQuery maps to https://core.telegram.org/bots/api#answerinlinequery
func (a *api) AnswerInlineQuery(query string, results []InlineQueryResult, opts *InlineQueryOptions) error {
	params := url.Values{}

	params.Set("inline_query_id", query)

	for idx := range results {
		results[idx].ValidateType()
	}

	res, err := json.Marshal(results)
	if err != nil {
		return err
	}
	params.Set("results", string(res))

	if opts != nil {
		optInt(params, "cache_time", opts.CacheTime)
		optBool(params, "is_personal", opts.Personal)
		optStr(params, "next_offset", opts.NextOffset)
		optStr(params, "switch_pm_text", opts.SwitchPM)
		optStr(params, "switch_pm_parameter", opts.SwitchParam)
	}

	return a.callAndSet("answerInlineQuery", params, nil)
}
