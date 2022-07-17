package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/ksusonic/owl-morning-bot/config"
	"github.com/ksusonic/owl-morning-bot/pkg/weather/ya_weather"
	"log"
	"os"
)

type Bot struct {
	Bot       *tgbotapi.BotAPI
	Cfg       *config.Config
	YaWeather *ya_weather.YaWeatherClient
}

func NewBot(cfg *config.Config) Bot {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELETOKEN"))
	if err != nil {
		log.Panic("Check telegram token!\n", err)
	}
	yaWeather := ya_weather.NewYaWeatherClient(&cfg.YaWeather)

	bot.Debug = cfg.Bot.Debug
	return Bot{Bot: bot, Cfg: cfg, YaWeather: yaWeather}
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
