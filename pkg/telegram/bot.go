package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"math/rand"
	"os"
)

type Bot struct {
	Bot *tgbotapi.BotAPI
	Cfg *Config
}

func NewBot(cfg *Config) Bot {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELETOKEN"))
	if err != nil {
		log.Panic("Check telegram token!\n", err)
	}
	bot.Debug = cfg.Bot.Debug
	return Bot{Bot: bot, Cfg: cfg}
}

func (b *Bot) Start() {
	log.Printf("Authorized on account %s", b.Bot.Self.UserName)

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

func (b *Bot) SendScheduledMessage() error {
	text := b.Cfg.Nlg.GoodMorning[rand.Int()%len(b.Cfg.Nlg.GoodMorning)]
	_, err := b.Bot.Send(tgbotapi.NewMessage(b.Cfg.Scheduler.ChatId, text))
	if err != nil {
		return err
	}
	selectedImage := b.Cfg.Nlg.Images[rand.Int()%len(b.Cfg.Nlg.Images)]
	_, err = b.Bot.Send(tgbotapi.NewPhoto(b.Cfg.Scheduler.ChatId, tgbotapi.FileURL(selectedImage)))
	if err != nil {
		log.Printf("Probably, bad image url: ")
	}
	log.Println("Sent:", text, "with image:", selectedImage)
	return nil
}

func (b *Bot) handleMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	_, err := b.Bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}

func (b *Bot) initUpdateChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.Bot.GetUpdatesChan(u)
}
