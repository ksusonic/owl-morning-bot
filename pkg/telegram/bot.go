package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

type Bot struct {
	bot *tgbotapi.BotAPI
	cfg *Config
}

func NewBot(cfg *Config) Bot {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELETOKEN"))
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = cfg.Bot.Debug
	return Bot{bot: bot, cfg: cfg}
}

func (b *Bot) Start() {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	for update := range b.initUpdateChannel() {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			sent, err := b.handleCommand(update.Message)
			if err != nil {
				log.Printf("Sending error: %v\ntext: %s\nto: %s", err, sent.Text, sent.Chat.UserName)
			}
			continue
		}

		b.handleMessage(update.Message)
	}
}

func (b *Bot) SendMessage(chatId int64, text string) error {
	_, err := b.bot.Send(tgbotapi.NewMessage(chatId, text))
	if err != nil {
		return err
	}
	return nil
}

func (b *Bot) handleMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	_, err := b.bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}

func (b *Bot) initUpdateChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u)
}
