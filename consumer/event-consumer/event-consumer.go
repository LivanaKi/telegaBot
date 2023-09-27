package eventconsumer

import (
	"log"
	"time"

	"github.com/Users/natza/telegaBot/events"
)

type Consumer struct {
	fetcher events.Fetcher
	processor events.Processor
	batchSize int
}

func New(fetcher events.Fetcher, processor events.Processor, batchSize int) Consumer{
	return Consumer{
		fetcher: fetcher,
		processor: processor,
		batchSize: batchSize,
	}
}

func (c Consumer) Start() error {
	for {
		gotEvents, err := c.fetcher.Fetch(c.batchSize)
		if err != nil {
			log.Printf("[Err] consumer: %s", err.Error())

			continue
		}
		if len(gotEvents) == 0{
			time.Sleep(1 * time.Second)

			continue
		}

		if err := c.handleEvennts(gotEvents); err != nil{
			log.Print(err)

			continue
		}
	}
}

func (c *Consumer) handleEvennts(events []events.Event) error {
	for _, event := range events {
		log.Printf("got new event: %s", event.Text)

		if err := c.processor.Process(event); err != nil {
			log.Printf("can't handle event: %s", err.Error())

			continue
		}
	}
	return nil
}