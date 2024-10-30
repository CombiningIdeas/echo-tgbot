package keyboard

import (
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

//In your keyboard package, make sure the MainKeyboard method returns
//the keyboard so it can be used in other files:

func MainKeyboard() *telego.ReplyKeyboardMarkup {
	return &telego.ReplyKeyboardMarkup{
		Keyboard: [][]telego.KeyboardButton{
			{
				tu.KeyboardButton("Repeat mode"),
				tu.KeyboardButton("Link saving mode"),
			},
			{
				tu.KeyboardButton("Alarm mode using messages"),
			},
		},
		ResizeKeyboard: true, OneTimeKeyboard: false, InputFieldPlaceholder: "Choose of the options",
	}
}

//this is what our keyboard would look like, if we wanted to make an inline-keyboard and not a reply keyboard, you
//can try to rewrite the bot yourself using an inline-keyboard to better consolidate the material

//func MainKeyboard() *telego.InlineKeyboardMarkup {
//	return &telego.InlineKeyboardMarkup{
//		InlineKeyboard: [][]telego.InlineKeyboardButton{
//			{
//				tu.InlineKeyboardButton("Repeat mode").WithCallbackData("callback_1"),
//				tu.InlineKeyboardButton("Link saving mode").WithCallbackData("callback_2"),
//			},
//			{
//				tu.InlineKeyboardButton("Alarm mode using messages").WithCallbackData("callback_3"),
//			},
//		},
//	}
//}

//keyboard for repeating messages:

func RepeatModeKeyboard() *telego.ReplyKeyboardMarkup {
	return &telego.ReplyKeyboardMarkup{
		Keyboard: [][]telego.KeyboardButton{
			{
				tu.KeyboardButton("⤴️ Come back #1"),
			},
		},
		ResizeKeyboard: true, OneTimeKeyboard: false, InputFieldPlaceholder: "Please send message",
	}

}

//keyboards for saving links

func LinkSavingModeKeyboard() *telego.ReplyKeyboardMarkup {
	return &telego.ReplyKeyboardMarkup{
		Keyboard: [][]telego.KeyboardButton{
			//send message: "you can send a link to save"
			{
				tu.KeyboardButton("List of saved links"),
				tu.KeyboardButton("⤴️ Come back #1"),
			},
		},
		ResizeKeyboard: true, OneTimeKeyboard: false, InputFieldPlaceholder: "Choose of the options",
	}

}

func ListSavedLinksKeyboard() *telego.ReplyKeyboardMarkup {
	return &telego.ReplyKeyboardMarkup{
		Keyboard: [][]telego.KeyboardButton{
			{
				tu.KeyboardButton("link removal mode"),
				tu.KeyboardButton("⤴️ Come back #2"),
				// for ... (list of links in the database, with numbering)
				//When you click the button, the URL of the selected link will be copied to the clipboard.
			},
		},
		ResizeKeyboard: true, OneTimeKeyboard: false, InputFieldPlaceholder: "Choose of the options",
	}

}

func LinkRemovalModeKeyboard() *telego.ReplyKeyboardMarkup {
	return &telego.ReplyKeyboardMarkup{
		Keyboard: [][]telego.KeyboardButton{
			{
				//send message: please write in the chat the number of the link you want to delete
				tu.KeyboardButton("⤴️ Come back #3"),
			},
		},
		ResizeKeyboard: true, OneTimeKeyboard: false, InputFieldPlaceholder: "Choose of the options",
	}

}

//Keyboard for sending SMS at a specific time to the user himself

func AlarmModeUsingMessagesKeyboard() *telego.ReplyKeyboardMarkup {
	return &telego.ReplyKeyboardMarkup{
		Keyboard: [][]telego.KeyboardButton{
			{
				tu.KeyboardButton("clear last message creation with timer"),

				tu.KeyboardButton("clear all message creation with timer"),
			},
			{
				tu.KeyboardButton("⤴️ Come back #1"),
			},
		},
		ResizeKeyboard: true, OneTimeKeyboard: false, InputFieldPlaceholder: "Choose of the options",
	}

}
