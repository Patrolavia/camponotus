// This program demonstrates the usage of package telegram.
// It DOES NOT follow go idioms.
package main

import (
	"log"
	"os"

	"github.com/Patrolavia/telegram"
)

func main() {
	token := os.Getenv("TOKEN")
	if token == "" {
		log.Fatal("Please set Telegram api token in environment variable TOKEN.")
	}
	api := telegram.New(token, nil)
	me, err := api.GetMe()
	if err != nil {
		log.Fatalf("Cannot get user info of this bot: %s", err)
	}

	mp := &messageProcessor{api, me, make(chan *telegram.Message)}
	ep := &editedProcessor{api, me, make(chan *telegram.Message)}
	iqp := &inlineQueryProcessor{api, me, make(chan *telegram.InlineQuery)}
	cirp := &chosenInlineResultProcessor{api, me, make(chan *telegram.ChosenInlineResult)}
	cqp := &callbackQueryProcessor{api, me, make(chan *telegram.CallbackQuery)}

	go mp.Run()
	go ep.Run()
	go iqp.Run()
	go cirp.Run()
	go cqp.Run()

	fetcher := &telegram.LongPollFetcher{
		mp.CH,
		ep.CH,
		iqp.CH,
		cirp.CH,
		cqp.CH,
		api,
	}

	log.Fatal(fetcher.Fetch(10, 30))
}
