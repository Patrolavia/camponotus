package main

import (
	"fmt"
	"log"

	"github.com/Patrolavia/telegram"
)

type editedProcessor struct {
	API telegram.API
	ME  *telegram.Victim
	CH  chan *telegram.Message
}

func (p *editedProcessor) Run() {
	for m := range p.CH {
		msg := fmt.Sprintf("Received edited message:\n```\n%#s\n```", toString(m))
		if _, err := p.API.SendMessage(m.From.Identifier(), msg, &telegram.Options{
			ReplyMarkup: &telegram.InlineKeyboardMarkup{
				[][]telegram.InlineKeyboardButton{
					[]telegram.InlineKeyboardButton{
						telegram.InlineKeyboardButton{
							Text: "1",
							Data: "one",
						},
						telegram.InlineKeyboardButton{
							Text: "2",
							Data: "two",
						},
					},
				},
			},
			ParseMode: telegram.MarkdownMode,
		}); err != nil {
			log.Fatalf("editedProcessor: failed to send message: %s", err)
		}
	}
}
