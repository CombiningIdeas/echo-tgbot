package mode_logic_handlers

import (
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

//this is the logic for message repeat mode

func RepeatModeLogicHandler(bot *telego.Bot, chatID telego.ChatID, lastMessage string, penultimateMessage string) {
	if (lastMessage != "") && (penultimateMessage == "Repeat mode") {
		message := tu.Message(chatID, lastMessage)
		_, _ = bot.SendMessage(message)
	}
}
