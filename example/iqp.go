package main

import (
	"log"

	"github.com/Patrolavia/telegram"
)

type inlineQueryProcessor struct {
	API telegram.API
	ME  *telegram.Victim
	CH  chan *telegram.InlineQuery
}

func (p *inlineQueryProcessor) Run() {
	results := []telegram.InlineQueryResult{
		&telegram.InlineQueryResultArticle{
			IQR: telegram.IQR{
				ID:                  "article",
				InputMessageContent: telegram.InputTextMessageContent("the article", telegram.TextMode, false),
				Title:               "Article",
			},
		},
		&telegram.InlineQueryResultPhoto{
			IQR: telegram.IQR{
				ID:       "photo",
				ThumbURL: "https://patrolavia.com/logo64.jpg",
				ReplyMarkup: &telegram.InlineKeyboardMarkup{
					[][]telegram.InlineKeyboardButton{
						[]telegram.InlineKeyboardButton{
							telegram.InlineKeyboardButton{
								Text: "img",
								URL:  "https://patrolavia.com/logo64.jpg",
							},
							telegram.InlineKeyboardButton{
								Text: "web",
								URL:  "https://patrolavia.com",
							},
						},
					},
				},
			},
			URL: "https://patrolavia.com/logo64.jpg",
		},
	}
	for q := range p.CH {
		if err := p.API.AnswerInlineQuery(q.ID, results, nil); err != nil {
			log.Fatal(err)
		}
	}
}
