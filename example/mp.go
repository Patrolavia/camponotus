package main

import (
	"fmt"
	"log"

	"github.com/Patrolavia/telegram"
)

type messageProcessor struct {
	API telegram.API
	ME  *telegram.Victim
	CH  chan *telegram.Message
}

func (p *messageProcessor) Run() {
	for m := range p.CH {
		msg := fmt.Sprintf("Received message:\n```\n%#s\n```\nEntities: %v", toString(m), m.EntityText())
		if _, err := p.API.SendMessage(m.From.Identifier(), msg, &telegram.Options{
			ReplyMarkup: &telegram.ReplyKeyboardMarkup{
				Keyboard: [][]telegram.KeyboardButton{
					[]telegram.KeyboardButton{
						telegram.KeyboardButton{
							Text: "one",
						},
						telegram.KeyboardButton{
							Text: "two",
						},
					},
					[]telegram.KeyboardButton{
						telegram.KeyboardButton{
							Text: "three",
						},
						telegram.KeyboardButton{
							Text: "four",
						},
					},
				},
				Once: true,
			},
			ParseMode: telegram.MarkdownMode,
		}); err != nil {
			log.Fatalf("messageProcessor: failed to send message: %s", err)
		}
	}
}
