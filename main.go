package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

const (
	webHook = "https://git.heroku.com/salty-fjord-82314.git"
)

func main() {
	botToken := "1378294896:AAECzw5BejQPLOy6mNRfhW45zYDjPmKXm6g"
	botAPI := "https://api.telegram.org/bot"
	botURL := botAPI + botToken
	offset := 0

	groupID := "-52457356"
	items, err := getPosts(groupID)
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range items {
		log.Print(item)
	}
	return

	for {
		updates, err := getUpdates(botURL, offset)
		if err != nil {
			log.Println("Somethingwent wrong: ", err.Error())
		}
		for _, update := range updates {
			err = respond(botURL, update)
			offset = update.UpdateId + 1
		}
		fmt.Println(updates)
	}
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal("creation bot: ", err)
	}
	log.Println("bot created")

	if _, err := bot.SetWebhook(tgbotapi.NewWebhook(webHook)); err != nil {
		log.Fatal("setting webHook %v; error: %v", webHook, err)
	}
	log.Println("vebHook set")

}

func getUpdates(botUrl string, offset int) ([]Update, error) {
	resp, err := http.Get(botUrl + "/getUpdates" + "?offset=" + strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var restResponse RestResponse
	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		return nil, err
	}
	return restResponse.Result, nil
}

func respond(botUrl string, update Update) error {
	var botMessage BotMessage
	botMessage.ChatId = update.Message.Chat.ChatId
	botMessage.Text = update.Message.Text
	buf, err := json.Marshal(botMessage)
	if err != nil {
		return err
	}
	_, err = http.Post(botUrl+"/sendMessage", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	return nil
}
