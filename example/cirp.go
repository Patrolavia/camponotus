package main

import (
	"fmt"

	"github.com/Patrolavia/telegram"
)

type chosenInlineResultProcessor struct {
	API telegram.API
	ME  *telegram.Victim
	CH  chan *telegram.ChosenInlineResult
}

func (p *chosenInlineResultProcessor) Run() {
	for r := range p.CH {
		fmt.Println(toString(r))
	}
}
