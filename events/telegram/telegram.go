package telegram

import (
	"github.com/Users/natza/telegaBot/clients/telegram"
	"github.com/Users/natza/telegaBot/events"
	"github.com/Users/natza/telegaBot/lib/e"
	"github.com/Users/natza/telegaBot/storage"
)

type Processor struct {
	tg *telegram.Client
	offset int
	storage storage.Storage
}

func New(client *telegram.Client, storage storage.Storage) *Processor{
	return &Processor{
		tg: client,
		storage: storage,
	}
}

func (p *Processor) Fetch(limit int) ([]events.Event, error){
	update, err := p.tg.Updates(p.offset, limit)
	if err != nil{
		return nil, e.Wrap("can't get events", err)
	}

	res := make([]events.Event, 0, len(update))

	for _, u := range update{
		res = append(res, event(u))
	}
}

func event(upd telegram.Update) events.Event{
	updType := fetchType(upd)
	
	res := events.Event{
		Type: updType,
		Text: fetchText(upd),
	}
}

func fetchText(upd telegram.Update) string{
	if upd.Message == nil{
		return ""
	}
	
	return upd.Message.Text
}

func fetchType(upd telegram.Update) events.Type {
	if upd.Message == nil{
		return events.Unknown
	}
	
	return events.Message
}
