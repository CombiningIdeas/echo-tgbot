package services

import (
	"echo-tgbot/internal/services/mode_logic_handlers"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func (h *HandlerStruct) DifferentCommandsForDifferentModesHandler(bot *telego.Bot) th.Handler {
	return func(bot *telego.Bot, update telego.Update) {
		if update.Message == nil {
			return
		}

		chatID := tu.ID(update.Message.Chat.ID)
		h.lastMessage = update.Message.Text

		//this is the logic for message repeat mode
		mode_logic_handlers.RepeatModeLogicHandler(bot, chatID, h.lastMessage, h.penultimateMessage)

		//This is the logic for messages with timers
		//We handle clearing all your SMS sent by time in timers  mode via message:
		mode_logic_handlers.TimerModeLogicHandler(bot, chatID, &h.lastMessage, &h.penultimateMessage, &h.timers, &h.flagAll, &h.flagOne,
			&h.messageWithTimer)

		// This is the logic for sending links to the database:
		mode_logic_handlers.LinkModeLogicHandler(update, bot, chatID, &h.lastMessage, &h.penultimateMessage)
	}
}
