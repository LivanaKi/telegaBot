package telegram

import "github.com/Users/natza/telegaBot/clients/telegram"

type Processor struct {
	tg *telegram.Client
	offset int
}

func New(client *telegram.Client) {
	
}