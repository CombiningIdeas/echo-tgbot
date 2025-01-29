package services

import (
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	"time"
)

//we can use interfaces to avoid being tied to package names when importing functions (once we create an interface with a
//structure that will implement that interface, it is already methods, not a function). We also become accessible to the
//fields of the created structure, where we can store the fields we need.

//For example, flags, arrays, strings (for remembering the answers we need from the user or other cases),
//numeric variables, for creating restrictions on something in the structure itself.

//Because using all these variables directly through the structure is much more convenient than simply creating variables,
//we ourselves can decide where to use pointers and where to copy copies to the fields of the structure

//create an interface:

type HandlerInterface interface {
	StartCommandHandler(bot *telego.Bot) th.Handler
	MainKeyboardHandler(bot *telego.Bot) th.Handler
	MainReturnKeyboardHandler(bot *telego.Bot) th.Handler

	//these methods are needed for keyboard navigation that saves links:
	LinkSavingModeReturnKeyboardHandler(bot *telego.Bot) th.Handler
	ListSavedLinksReturnKeyboardHandler(bot *telego.Bot) th.Handler
	LinkSavingModeKeyboardHandler(bot *telego.Bot) th.Handler
	ListSavedLinksKeyboardHandler(bot *telego.Bot) th.Handler

	//general bot logic
	DifferentCommandsForDifferentModesHandler(bot *telego.Bot) th.Handler
}

// struct:
type HandlerStruct struct {
	lastMessage        string
	penultimateMessage string
	messageWithTimer   string
	flagOne            bool
	flagAll            bool
	timers             []*time.Timer
}

// constructor:
func NewHandlerStruct(lastMessage string, penultimateMessage string, messageWithTimer string,
	flagOne bool, flagAll bool, timers []*time.Timer) *HandlerStruct {
	return &HandlerStruct{
		lastMessage:        lastMessage,
		penultimateMessage: penultimateMessage,
		messageWithTimer:   messageWithTimer,
		flagOne:            flagOne,
		flagAll:            flagAll,
		timers:             timers,
	}
}
