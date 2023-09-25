package telegram

import (
	"errors"
	"log"
	"net/url"
	"strings"

	//"github.com/Users/natza/telegaBot/clients/telegram"
	"github.com/Users/natza/telegaBot/lib/e"
	"github.com/Users/natza/telegaBot/storage"
	//"github.com/Users/natza/telegaBot/storage/files"
)

const (
	RndCmd = "/rnd"
	HelpCmd = "/help"
	StartCmd = "/start"
)

func (p *Processor) doCmd(text string, chatID int, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("got new command '%s' from '%s'", text, username)

	if isAddCmd(text) {
		return p.savePage(chatID, text, username)
	}

	switch text {
	case RndCmd:
		return p.SendRandom()
	case HelpCmd:
	case StartCmd:
	default:
	}
}

func (p *Processor) savePage(chatID int, pageURL string, username string) (err error){
	defer func() { err = e.Wrap("can't do command: save page", err)}()

		page := &storage.Page{
		URL: pageURL,
		UserName: username,
	}

	IsExists, err := p.storage.IsExists(page)
	if err != nil {
		return err
	}
	if IsExists{
		return p.tg.SendMessage(chatID, msgAlreadyExists)
	}

	if err := p.storage.Save(page); err != nil{
		return err
	}

	if err := p.tg.SendMessage(chatID, msgSaved); err != nil {

	}
	return nil
}

func (p *Processor) SendMessage (chatID int, username string) (err error) {
	defer func() { err = e.Wrap("can't do command: can't send random", err)}()

	page, err := p.storage.PickRandom(username)
	if err != nil && !errors.Is(err, storage.ErrNoSavePages){
		return err 
	}
	if errors.Is(err, storage.ErrNoSavePages){
		return p.tg.SendMessage(chatID, msgNoSavedPages)
	}
	if err := p.tg.SendMessage(chatID, page.URL); err != nil {
		return err
	}

	return p.storage.Remove(page)
}

func (p *Processor) SendHelp (chatID int) error {
	return p.tg.SendMessage(chatID, msgHelp)
}

func (p *Processor) SendHello (chatID int) error {
	return p.tg.SendMessage(chatID, msgHello)
}

func isAddCmd(text string) bool{
	return isURL(text)
}

func isURL(text string) bool {
	u, err := url.Parse(text)

	return err == nil && u.Host != ""
}