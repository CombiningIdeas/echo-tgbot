package services

import (
	"echo-tgbot/models/keyboard"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

// Here we describe the main methods (handlers) that will be used by the bot.
// processing the "start" command:

func (h *HandlerStruct) StartCommandHandler(bot *telego.Bot) th.Handler {
	return func(bot *telego.Bot, update telego.Update) {
		if update.Message == nil { //ignore "not messages"
			return
		}

		chatID := tu.ID(update.Message.Chat.ID) //get chatID

		//we write this way everywhere, because after we select our mode via the keyboard, we will just need to write the text,
		//and then the pressed button on the keyboard will always be the penultimate one, otherwise it will be overwritten,
		//and we will not be able to understand it in any way Thus the mode in which we keep:
		h.penultimateMessage = update.Message.Text

		//Connecting the main keyboard
		Keyboard := keyboard.MainKeyboard()

		message := tu.Message(chatID, "Hello friend🌟! Welcome to my telegram bot. Now you have a keyboard"+
			" below, depending on the buttons you press, certain functionality will be available to you.💪").WithReplyMarkup(Keyboard)

		_, _ = bot.SendMessage(message)

	}
}
