package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

const (
	start = "start"
	ping  = "ping"
	pwd   = "pwd"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) (tgbotapi.Message, error) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Я не знаю этого, шайтан")
	switch message.Command() {
	case start:
		msg.Text = "Здарова! Если есть вопросы, пиши @ksusonic"
	case ping:
		msg.Text = fmt.Sprintf("pong v%s 🏓", b.Version)
	case pwd:
		msg.Text = strconv.FormatInt(message.Chat.ID, 10)
	}

	return b.Bot.Send(msg)
}
