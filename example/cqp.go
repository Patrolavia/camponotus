package main

import (
	"fmt"

	"github.com/Patrolavia/telegram"
)

type callbackQueryProcessor struct {
	API telegram.API
	ME  *telegram.Victim
	CH  chan *telegram.CallbackQuery
}

func (p *callbackQueryProcessor) Run() {
	for c := range p.CH {
		fmt.Println(toString(c))
	}
}
