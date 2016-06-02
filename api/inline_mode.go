// This file is part of Camponotus
// Camponotus is free software: see LICENSE.txt for more details.

package api

import "net/url"

// AnswerInlineQuery maps to https://core.telegram.org/bots/api#answerinlinequery
func (a *API) AnswerInlineQuery(
	query string, results []byte, cache int, personal bool, next, pm, pmParam string,
) error {
	params := url.Values{}

	params.Set("inline_query_id", query)
	params.Set("results", string(results))
	optInt(params, "cache_time", cache)
	optBool(params, "is_personal", personal)
	optStr(params, "next_offset", next)
	optStr(params, "switch_pm_text", pm)
	optStr(params, "switch_pm_parameter", pmParam)

	return a.callAndSet("answerInlineQuery", params, nil)
}
