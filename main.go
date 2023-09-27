package main

import (
	"flag"
	"log"

	"github.com/Users/natza/telegaBot/clients/telegram"
	eventconsumer "github.com/Users/natza/telegaBot/consumer/event-consumer"
	"github.com/Users/natza/telegaBot/events/telegram2"
	"github.com/Users/natza/telegaBot/storage/files"
)

//6336153425:AAEmtiGS6iZead1F_EK7ZPo3c3taeccSZuc

const(
	tgBotHost = "api.telegram.org"
	storagePath = "storage"
	batchSize = 100
)

func main() {

	evetsProcessor := telegram2.New(
		telegram.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Print("service started")

	consumer := eventconsumer.New(evetsProcessor, evetsProcessor, batchSize)
	if err := consumer.Start(); err != nil{
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	token := flag.String("token-bot-token", "", "token for access to telegram bot",)

	flag.Parse()

	if *token == ""{
		log.Fatal("token is not specified")
	}

	return *token
}