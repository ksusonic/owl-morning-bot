package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	commandStart = "start"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) (tgbotapi.Message, error) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Я не знаю этого, шайтан")
	switch message.Command() {
	case commandStart:
		msg.Text = "Приветик!"
	}

	return b.bot.Send(msg)
}
