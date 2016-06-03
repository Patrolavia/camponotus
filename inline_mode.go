// This file is part of Camponotus
// Camponotus is free software: see LICENSE.txt for more details.

package telegram

import (
	"encoding/json"
	"net/url"
)

// AnswerInlineQuery maps to https://core.telegram.org/bots/api#answerinlinequery
func (a *api) AnswerInlineQuery(
	query string, results []InlineQueryResult, cache int, personal bool, next, pm, pmParam string,
) error {
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

	optInt(params, "cache_time", cache)
	optBool(params, "is_personal", personal)
	optStr(params, "next_offset", next)
	optStr(params, "switch_pm_text", pm)
	optStr(params, "switch_pm_parameter", pmParam)

	return a.callAndSet("answerInlineQuery", params, nil)
}
