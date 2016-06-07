package main

import (
	"fmt"
	"log"

	"github.com/Patrolavia/telegram"
)

type callbackQueryProcessor struct {
	API telegram.API
	ME  *telegram.Victim
	CH  chan *telegram.CallbackQuery
}

func (p *callbackQueryProcessor) Run() {
	for c := range p.CH {
		msg := fmt.Sprintf("Received callback query:\n```\n%s\n```", toString(c))
		if _, err := p.API.EditText(
			c.Message.Chat.Identifier(),
			c.Message.ID,
			msg,
			telegram.MarkdownMode,
			false,
			nil,
		); err != nil {
			log.Fatalf("callbackQueryProcessor: failed to send message: %s", err)
		}
	}
}
