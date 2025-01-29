package services

import (
	"echo-tgbot/models/keyboard"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

//Processing clicks on the "Main keyboard" buttons:

func (h *HandlerStruct) MainKeyboardHandler(bot *telego.Bot) th.Handler {
	//Here we will process the main keyboard keys
	//"Repeat mode", "Link saving mode", "Alarm mode using messages"
	return func(bot *telego.Bot, update telego.Update) {
		if update.Message == nil {
			return
		}

		chatID := tu.ID(update.Message.Chat.ID)

		h.penultimateMessage = update.Message.Text

		switch update.Message.Text {
		case "Repeat mode":
			message := tu.Message(chatID, "You are in Repeat mode.").WithReplyMarkup(keyboard.RepeatModeKeyboard())
			_, _ = bot.SendMessage(message)
		case "Link saving mode":
			message := tu.Message(chatID, "You are in Link saving mode. Any link you send will be "+
				"saved to the database").WithReplyMarkup(keyboard.LinkSavingModeKeyboard())
			_, _ = bot.SendMessage(message)
		case "Timer mode using messages":
			message := tu.Message(chatID, "You are in message with timer mode😊. To ensure that the message arrives at your "+
				"designated time, first enter the message itself, then “:”, and then indicate how many minutes later the message "+
				"will be sent to you, otherwise there will be an error. An example of such a message: “Don’t forget to walk the dog:851.” "+
				"This message will arrive in 851 minutes, that is, in 14 hours and 11 minutes.You can also indicate the following message:"+
				" “Don’t forget to wash the dishes in 5 and a half minutes: 5.5”").WithReplyMarkup(keyboard.AlarmModeUsingMessagesKeyboard())
			_, _ = bot.SendMessage(message)
		}
	}
}

// the keyboard handler "main-keyboard" for the command back command:

func (h *HandlerStruct) MainReturnKeyboardHandler(bot *telego.Bot) th.Handler {
	return func(bot *telego.Bot, update telego.Update) {
		if update.Message == nil {
			return
		}

		chatID := tu.ID(update.Message.Chat.ID)

		h.penultimateMessage = update.Message.Text

		if update.Message.Text == "⤴️ Come back #1" {
			message := tu.Message(chatID, "Returning to the main menu.").WithReplyMarkup(keyboard.MainKeyboard())
			_, _ = bot.SendMessage(message)
		}
	}
}

//the keyboard handler "LinkSavingMode-keyboard" for the command back command:

func (h *HandlerStruct) LinkSavingModeReturnKeyboardHandler(bot *telego.Bot) th.Handler {
	return func(bot *telego.Bot, update telego.Update) {
		if update.Message == nil {
			return
		}

		chatID := tu.ID(update.Message.Chat.ID)

		h.penultimateMessage = update.Message.Text

		if update.Message.Text == "⤴️ Come back #2" {
			message := tu.Message(chatID, "Returning to the Link Saving Mode Keyboard.").WithReplyMarkup(keyboard.LinkSavingModeKeyboard())
			_, _ = bot.SendMessage(message)
		}
	}
}

//the keyboard handler "ListSavedLinks-keyboard" for the command back command:

func (h *HandlerStruct) ListSavedLinksReturnKeyboardHandler(bot *telego.Bot) th.Handler {
	return func(bot *telego.Bot, update telego.Update) {
		if update.Message == nil {
			return
		}

		chatID := tu.ID(update.Message.Chat.ID)

		h.penultimateMessage = update.Message.Text

		if update.Message.Text == "⤴️ Come back #3" {
			message := tu.Message(chatID, "Returning to the List Saved Links Keyboard. "+
				"You will get a list of saved links if you click the \"Show list of my links ℹ\"button:").WithReplyMarkup(keyboard.ListSavedLinksKeyboard())
			_, _ = bot.SendMessage(message)
		}
	}
}

//special handlers for "saving links":

func (h *HandlerStruct) LinkSavingModeKeyboardHandler(bot *telego.Bot) th.Handler {
	return func(bot *telego.Bot, update telego.Update) {
		if update.Message == nil {
			return
		}

		chatID := tu.ID(update.Message.Chat.ID)

		h.penultimateMessage = update.Message.Text

		if update.Message.Text == "List of saved links" {
			// for ... displaying a list of saved links on the screen will be limited to 1000 links:
			message := tu.Message(chatID, "You will get a list of saved links if you click the"+
				" \"Show list of my links ℹ\"button:").WithReplyMarkup(keyboard.ListSavedLinksKeyboard())
			_, _ = bot.SendMessage(message)
		}
	}
}

func (h *HandlerStruct) ListSavedLinksKeyboardHandler(bot *telego.Bot) th.Handler {
	return func(bot *telego.Bot, update telego.Update) {
		if update.Message == nil {
			return
		}

		chatID := tu.ID(update.Message.Chat.ID)

		h.penultimateMessage = update.Message.Text

		if update.Message.Text == "link removal mode" {
			message := tu.Message(chatID, "you are in link removal mode, you can specify the link number from "+
				"the list of links you sent, and then this link will be deleted from the database, this can be useful "+
				"in order not to clutter up your list with all the links or to delete those that are no longer needed").WithReplyMarkup(keyboard.LinkRemovalModeKeyboard())
			_, _ = bot.SendMessage(message)
		}
	}
}
