package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"../scheduler"
)

const botAPI = "https://api.telegram.org/bot"

func StartBot(botToken string, errChan chan<- error) {
	var schedulerMessageChan = make(chan scheduler.Message)
	var updatesMessageChan = make(chan []Update)
	var offset = make(chan int)
	// offset <- 0
	botURL := botAPI + botToken
	go scheduler.StartScheduler(schedulerMessageChan)
	go getUpdates(botURL, updatesMessageChan, offset, errChan)
	for {
		select {
		case msg := <-schedulerMessageChan:
			go processScheduler(botURL, msg, errChan)
		case msg := <-updatesMessageChan:
			go processUpdates(botURL, msg, offset, errChan)
		}
	}
}

func getUpdates(botURL string, updagesChan chan<- []Update, offset <-chan int, errChan chan<- error) {
	off := 1 << 62
	for {
		resp, err := http.Get(botURL + "/getUpdates" + "?offset=" + strconv.Itoa(off))
		if err != nil {
			errChan <- err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			errChan <- err
		}
		var restResponse RestResponse
		err = json.Unmarshal(body, &restResponse)
		if err != nil {
			errChan <- err
		}
		if restResponse.Result != nil {
			updagesChan <- restResponse.Result
		}
		off = <-offset
	}
}

func processUpdates(botURL string, updates []Update, offset chan<- int, errChan chan<- error) {
	maxOffset := 0
	for _, update := range updates {
		var respondMessage string
		if update.Message.Text[0] == '/' {
			var err error
			respondMessage, err = commands(update.Message.Text)
			if err != nil {
				respondMessage = "Command not found ☹️"
			}
		} else {
			respondMessage = "Hello!"
		}

		if update.UpdateID > maxOffset {
			maxOffset = update.UpdateID
		}

		fmt.Println(respondMessage)
		err := sendMessage(botURL, respondMessage, update.Message.Chat.ChatID)
		if err != nil {
			errChan <- err
		}
	}
	offset <- maxOffset + 1
}

func processScheduler(botURL string, message scheduler.Message, errChan chan<- error) {
	sendMessage(botURL, message.Text, 265202007)
}

func sendMessage(botURL string, message string, chatID int) error {
	var botMessage BotMessage
	botMessage.ChatID = chatID
	botMessage.Text = message
	messageBytes, err := json.Marshal(botMessage)
	if err != nil {
		return err
	}
	_, err = http.Post(botURL+"/sendMessage", "application/json", bytes.NewBuffer(messageBytes))
	if err != nil {
		return err
	}
	return nil
}
