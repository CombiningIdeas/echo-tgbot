package app

import (
	"echo-tgbot/internal/services"
	"fmt"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	"log"
	"os"
)

func Startup() {

	//our token:
	var botToken string = "botToken"

	//create a bot instance using the telego library
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
		os.Exit(1)
	}

	//In our case, we will use Long Polling as an executor;
	//why we will use this particular method,
	//look in the documentation -> docs/project_logic.md -> API Part

	//Since we are using "Long Polling", we will name the variable accordingly
	updates, _ := bot.UpdatesViaLongPolling(nil)
	//The Updates Via Long Polling function has 2 return values "<-chan Update" and "error".
	//"<-chan Update" - This is a channel from which you can receive updates (objects of type Update).
	//It will provide new updates as they become available from Telegram. This allows you to process
	//updates asynchronously, receiving them as they arrive.

	BotHandler, _ := th.NewBotHandler(bot, updates)

	//Closing the channel through which data is transmitted will be
	//called when your Startup function is closed, and not before
	defer bot.StopLongPolling()
	defer BotHandler.Stop()

	//create a new instance of the structure:
	handler := services.NewHandlerStruct("", "", "", true, true, nil)

	//start command handler:
	BotHandler.Handle(handler.StartCommandHandler(bot), th.CommandEqual("start")) // indicate the type of function handler inside Handle,
	// and pass the "start" command there

	// Connect the keyboard handler "main-keyboard" for the command back command
	BotHandler.Handle(handler.MainReturnKeyboardHandler(bot), th.TextEqual("⤴️ Come back #1"))

	// Connect the keyboard handler "LinkSavingMode-keyboard" for the command back command
	BotHandler.Handle(handler.LinkSavingModeReturnKeyboardHandler(bot), th.TextEqual("⤴️ Come back #2"))

	// Connect the keyboard handler "ListSavedLinks-keyboard" for the command back command
	BotHandler.Handle(handler.ListSavedLinksReturnKeyboardHandler(bot), th.TextEqual("⤴️ Come back #3"))

	//in our case, several nested keyboards are found only in the mode of saving links, since most of the logic
	//will be implemented there, and there will be interaction with the database:
	BotHandler.Handle(handler.LinkSavingModeKeyboardHandler(bot), th.TextEqual("List of saved links"))
	BotHandler.Handle(handler.ListSavedLinksKeyboardHandler(bot), th.TextEqual("link removal mode"))

	// Connect handlers for the main-keyboard
	BotHandler.Handle(handler.MainKeyboardHandler(bot), th.TextEqual("Repeat mode"))
	BotHandler.Handle(handler.MainKeyboardHandler(bot), th.TextEqual("Link saving mode"))
	BotHandler.Handle(handler.MainKeyboardHandler(bot), th.TextEqual("Timer mode using messages"))

	BotHandler.Handle(handler.DifferentCommandsForDifferentModesHandler(bot), th.AnyMessage())

	//Let's launch the bot
	BotHandler.Start()

}
