package mode_logic_handlers

import (
	"echo-tgbot/internal/repository"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"log"
	"net/url"
	"strconv"
)

// This is the logic for sending links to the database:

func LinkModeLogicHandler(update telego.Update, bot *telego.Bot, chatID telego.ChatID, lastMessage *string,
	penultimateMessage *string) {

	//in our case we need to first establish a connection to the database:
	repository.Repository()

	if (*penultimateMessage == "Link saving mode" || *penultimateMessage == "⤴️ Come back #2") &&
		(*lastMessage != "") {

		//in our case, we will use chat ID as user ID, to simplify communication,
		//store data by which users can be identified

		//we check whether a real link was sent (we create our own function for this):
		if isValidURL(*lastMessage) == true {
			replyToMessage := *lastMessage + " - correct link."
			message := tu.Message(chatID, replyToMessage)
			_, _ = bot.SendMessage(message)
		} else {
			replyToMessage := *lastMessage + " - incorrect link"
			message := tu.Message(chatID, replyToMessage)
			_, _ = bot.SendMessage(message)
			return //if there is an error and the link is incorrect,
			// then we do not allow it to survive in its database
		}

		name_link := *lastMessage
		chat_id := strconv.Itoa(int(update.Message.Chat.ID))

		_, err := repository.Database.Exec("INSERT INTO user_links (user_id, name_link) VALUES ($1, $2)", chat_id, name_link)
		if err != nil {
			log.Println(err)
		}

	} else if (*penultimateMessage == "List of saved links" || *penultimateMessage == "⤴️ Come back #3") &&
		(*lastMessage == "Show list of my links") {

		//string array
		links, err := gettingLinks(update, bot, chatID)
		if err != nil { //We check for errors while the program is running.
			log.Println(err)
			return
		}

		// Send a list of saved links
		if len(links) == 0 {
			message := tu.Message(chatID, "There are no saved links.")
			_, _ = bot.SendMessage(message)
		} else {
			var replyToMessage string
			for iter, link := range links {
				replyToMessage += strconv.Itoa(iter+1) + " - " + link + "\n" //We create a message with links
			}
			message := tu.Message(chatID, replyToMessage)
			_, _ = bot.SendMessage(message)
		}
		//log.Println("Last Message: %s", *penultimateMessage)
	} else if *penultimateMessage == "link removal mode" {
		if NumberOfUserLinks, err := strconv.Atoi(*lastMessage); err == nil {

			// *lastMessage is an integer
			//you also need to verify the correct number entered by the user
			//does the user have such a reference number at all?
			if NumberOfUserLinks <= 0 {
				var replyToMessage string = "You entered an incorrect number that is not in the list of saved links."
				message := tu.Message(chatID, replyToMessage)
				_, _ = bot.SendMessage(message)
				return //finishing the function
			}

			//Here we will have to partially use the code that was already used before,
			//so we can put it in a separate block in the function
			links, err := gettingLinks(update, bot, chatID)
			if err != nil { //We check for errors while the program is running.
				log.Println(err)
				return //finishing the function
			}

			//checking if the number entered by the user is correct
			idLink, _ := strconv.Atoi(*lastMessage)
			var flag bool = false
			linkToDelete := ""
			for iter, link := range links {
				if iter == idLink {
					flag = true
					//Accordingly, we need to remember the link, which we will then delete
					linkToDelete = link
				}
			}

			//if the flag has not changed, it means the user entered the incorrect number
			if flag == false {
				var replyToMessage string = "You entered an incorrect number that is not in the list of saved links."
				message := tu.Message(chatID, replyToMessage)
				_, _ = bot.SendMessage(message)
				return //finishing the function
			}

			//Now we need to write a request here. which will delete the desired link in the database table, after
			//which we will send a message that the send link was successfully deleted, and we will send again a
			//list of links, this is already described below.
			chat_id := strconv.Itoa(int(update.Message.Chat.ID))
			if _, err = repository.Database.Exec("DELETE FROM user_links WHERE user_id = $1 AND name_link = $2",
				chat_id, linkToDelete); err != nil {
				// Handling an SQL query execution error
				replyToMessage := "An error occurred while trying to delete the link. Please try again later."
				message := tu.Message(chatID, replyToMessage)
				_, _ = bot.SendMessage(message)
				return //finishing the function
			}

			var replyToMessage string = "The link number you sent has been successfully deleted."
			for iter, link := range links {
				replyToMessage += strconv.Itoa(iter+1) + " - " + link + "\n" //We create a message with links
			}

		} else {
			// *lastMessage is not an integer
			var replyToMessage string = "You entered a non-integer number! Try again!"
			message := tu.Message(chatID, replyToMessage)
			_, _ = bot.SendMessage(message)
		}
	}
}

// We describe our function to check the validity of the send link:
func isValidURL(link string) bool {
	parsedURL, err := url.Parse(link)
	if err != nil {
		return false // A parse error means this is an invalid URL.
	}

	// Check that the scheme (for example, http, https) and host are not empty:
	return (parsedURL.Scheme != "") && (parsedURL.Host != "")
}

// Access to a database of links saved by each user
func gettingLinks(update telego.Update, bot *telego.Bot, chatID telego.ChatID) ([]string, error) {

	//in our case, we will use chat ID as user ID, to simplify communication,
	//store data by which users can be identified

	chat_id := strconv.Itoa(int(update.Message.Chat.ID))

	// write a request
	rows, err := repository.Database.Query("SELECT name_link FROM user_links WHERE user_id = $1", chat_id)

	// error check:
	if err != nil {
		log.Println(err)
		message := tu.Message(chatID, "Error while retrieving saved links.")
		_, _ = bot.SendMessage(message)
		return nil, err
	}

	defer rows.Close() // Close lines after use (when all the rest of the
	// code is executed, they will close themselves)

	var links []string //create a slice of rows where our links from the database will be stored to be sent to the user
	for rows.Next() == true {
		var name_link string

		//error check:
		if err := rows.Scan(&name_link); err != nil {
			log.Println(err)
			message := tu.Message(chatID, "Error reading link.")
			_, _ = bot.SendMessage(message)
			return nil, err
		}

		links = append(links, name_link) // Adding a link to the slice
	}

	//After completing the iteration, it is recommended to check for errors
	if err := rows.Err(); err != nil {
		log.Println(err)
		message := tu.Message(chatID, "Error while retrieving saved links.")
		_, _ = bot.SendMessage(message)
		return nil, err
	}

	return links, nil
}
