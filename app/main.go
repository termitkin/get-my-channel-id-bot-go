package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/termitkin/get-my-channel-id-bot-go/app/request"
	"github.com/termitkin/get-my-channel-id-bot-go/app/response"
	"github.com/termitkin/get-my-channel-id-bot-go/app/setups"
	"github.com/termitkin/get-my-channel-id-bot-go/app/utils"
)

func handleRequest(res http.ResponseWriter, req *http.Request) {
	contentTypeIsValid := request.CheckContentType(res, req)

	if !contentTypeIsValid {
		fmt.Println("Content type is not valid")

		return
	}

	_, err := res.Write([]byte("ok"))

	if err != nil {
		fmt.Println(err)
	}

	body := request.GetJsonBody(res, req)

	if body == nil {
		fmt.Println("Body is empty")

		return
	}

	fmt.Println(body)

	if body["channel_post"] == nil {
		fmt.Println("Message is not a channel post")

		return
	}

	channelPost := body["channel_post"].(map[string]any)

	if channelPost["text"] == nil {
		fmt.Println("Message is not a text")

		return
	}

	text := channelPost["text"].(string)

	fmt.Println(text)

	if text != "@get_my_channel_id_bot" {
		fmt.Println("Message is not a command")

		return
	}

	chatId := channelPost["chat"].(map[string]any)["id"].(float64)
	chatTitle := channelPost["chat"].(map[string]any)["title"].(string)
	chatUsername := channelPost["chat"].(map[string]any)["username"].(string)
	chatType := channelPost["chat"].(map[string]any)["type"].(string)

	message := ""

	if chatId != 0 {
		s := utils.FloatToStr(chatId)

		message += fmt.Sprintf("Chat ID: %s\n", s)
	}

	if chatTitle != "" {
		message += fmt.Sprintf("Chat title: %s\n", chatTitle)
	}

	if chatUsername != "" {
		message += fmt.Sprintf("Chat username: %s\n", chatUsername)
	}

	if chatType != "" {
		message += fmt.Sprintf("Chat type: %s\n", chatType)
	}

	fmt.Println(message)

	query := response.GetUrlQuery(message, utils.FloatToStr(chatId))
	url := response.GetUrl(query)

	fmt.Println(url)

	response.SendMessage(url)
}

func init() {
	varExists := setups.CheckEnvVariable("TELEGRAM_BOT_TOKEN")

	if !varExists {
		log.Fatalln("TELEGRAM_BOT_TOKEN variable is not set")
	}
}

func main() {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8080", nil)
}
