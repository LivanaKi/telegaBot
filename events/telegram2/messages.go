package telegram2

const msgHelp = `I can save and keep you pages. Also I can offer you them to read.
In order to save the page, just send me all link to it.
In order to get a random page from your list, send me command /rnd.
Caution! After that, this page will be removed from your list!`

const msgHello = "Hi there! \n" + msgHelp

const (	
	msgUnknownCommand = "Unknown command\n" + "Enter new pages\nComand:\n/start\n/rnd\n/help"
	msgNoSavedPages = "You have no saved pages\n" + "Enter new pages\nComand:\n/start\n/rnd\n/help"
	msgSaved = "Saved!\n" + "Enter new pages\nComand:\n/start\n/rnd\n/help"
	msgAlreadyExists = "You have already have this page in your list\n" + "Enter new pages\nComand:\n/start\n/rnd\n/help"
)