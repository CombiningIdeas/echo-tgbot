package mode_logic_handlers

import (
	"fmt"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"strconv"
	"time"
)

//This is the logic for messages with timers
//We handle clearing all your SMS sent by time in timers  mode via message:

func TimerModeLogicHandler(bot *telego.Bot, chatID telego.ChatID, lastMessage *string, penultimateMessage *string, timers *[]*time.Timer,
	flagAll *bool, flagOne *bool, messageWithTimer *string) {
	if (*lastMessage == "clear all message creation with timer") && (*penultimateMessage == "Timer mode using messages") {
		//then we reset the final line that needs to be displayed and reset the final seconds (this is the time after which the message will arrive).
		//As a result, we will send an empty message and send it immediately, rather than wait and take up system resources.

		//THIS IS A VERY IMPORTANT PART OF THE CODE FOR CLEARING ALL TIMERS:
		// Stop all active timers
		for _, timer := range *timers {
			timer.Stop() // Stop timer
		}

		//reset the list of timers completely
		*timers = []*time.Timer{}

		*flagAll = false

		message := tu.Message(chatID, "All Messages that should have been sent at the right time were successfully deleted.")
		_, _ = bot.SendMessage(message)

		//We process the clearing of your last SMS sent by time in Alarm mode with the following message:
	} else if (*lastMessage == "clear last message creation with timer") && (*penultimateMessage == "Timer mode using messages") {
		//then we reset the final line that needs to be displayed and reset the final seconds (this is the time after which the message will arrive).
		//As a result, we will send an empty message and send it immediately, rather than wait and take up system resources.

		(*timers)[len(*timers)-1].Stop() // Stop the last timer

		//here you don’t have to write cleaning the timer, since it will still be automatically executed in the main code,
		//after clearing the timer, in the example above we wrote it because we knew that we needed to clear absolutely the
		//entire array, and here is a specific timer, which means it’s best do not duplicate the code and let the goroutine
		//below do the cleaning of the timer from the timer array

		*flagOne = false

		message := tu.Message(chatID, "A last message that should have been sent by the right time was successfully deleted.")
		_, _ = bot.SendMessage(message)
	} else if (*lastMessage != "") && (*penultimateMessage == "Timer mode using messages") {

		//we introduce restrictions on the timer, up to 10 timers, the user can create:
		if len(*timers) >= 10 {
			message := tu.Message(chatID, "Cannot create more than 10 timers.")
			_, _ = bot.SendMessage(message)
			return
		}

		//Only when a new SMS is sent, everything is updated and automatically becomes true:
		*flagAll = true
		*flagOne = true
		*messageWithTimer = *lastMessage

		var tmpMessage string
		var tmpTime string

		ii := 0
		for ; ii < len(*lastMessage); ii++ {
			if string((*messageWithTimer)[ii]) == ":" {
				break
			}
			tmpMessage = tmpMessage + string((*messageWithTimer)[ii])
		}
		ii++

		for ; ii < len(*lastMessage); ii++ {
			tmpTime = tmpTime + string((*messageWithTimer)[ii])
		}

		number, err := strconv.ParseFloat(tmpTime, 64) // Convert the string to float64
		if err != nil {
			fmt.Println("Ошибка:", err)
			return
		}

		seconds := int64(number * 60) //convert to seconds so that they are also supported

		// Creating a new timer:
		timer := time.NewTimer(time.Duration(seconds) * time.Second)
		*timers = append(*timers, timer)

		if *flagOne == false {
			timer.Stop()
			return
		}

		go func() {
			<-timer.C
			//we must write this way, the whole point is that we always first create a timer, and then we can only change our mind and cancel,
			//either all timers or just one, since this cancellation must be carried out precisely after blocking the channel, since the
			//initial the value is still true and the channel will be blocked in any case, and if the check is written before it, then when
			//the value changes to false, it will be too late, since initially it was true and the channel has already been blocked and the
			//message will still be send:

			if *flagAll == true {
				message := tu.Message(chatID, tmpMessage)
				_, _ = bot.SendMessage(message)
			}

			//free the array of timers:
			for i := 0; i < len(*timers); i++ {
				if (*timers)[i] == timer {
					*timers = append((*timers)[:i], (*timers)[i+1:]...)
					break
				}
			}
		}()

		//You can, of course, write without goroutines and anonymous functions, but then it will
		//be less efficient and more complicated:
		//time.AfterFunc(time.Duration(seconds)*time.Second, func() {
		//	if (h.flagAll == true) && (h.flagOne == true) {
		//		message := tu.Message(chatID, tmpMessage)
		//		_, _ = bot.SendMessage(message)
		//	}
		//	h.flagOne = true //since we are focusing on the last message, we return the flag to the desired position
		//})

	}
}
