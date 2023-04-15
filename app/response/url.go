package response

import (
	"fmt"
	"net/url"
	"os"
)

func GetUrlQuery(message, chatId string) string {
	query := url.Values{}
	query.Add("chat_id", chatId)
	query.Add("text", message)

	return query.Encode()
}

func GetUrl(query string) string {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")

	return fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?%s", token, query)
}
