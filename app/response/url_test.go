package response

import (
	"fmt"
	"os"
	"testing"
)

func TestGetUrl(t *testing.T) {
	chatId := "123"
	message := "hello"
	botToken := "qwerty123"

	query := fmt.Sprintf("chat_id=%s&text=%s", chatId, message)
	expected := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s", botToken, chatId, message)

	os.Setenv("TELEGRAM_BOT_TOKEN", "qwerty123")
	defer os.Unsetenv("TELEGRAM_BOT_TOKEN")

	actual := GetUrl(query)

	if actual != expected {
		t.Errorf("GetUrl() = %s; want %s", actual, expected)
	}
}

func TestGetUrlQuery(t *testing.T) {
	chatId := "123"
	message := "hello"
	expected := "chat_id=123&text=hello"

	actual := GetUrlQuery(message, chatId)

	if actual != expected {
		t.Errorf("GetUrlQuery() = %s; want %s", actual, expected)
	}
}
