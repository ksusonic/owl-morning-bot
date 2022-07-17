package scheduler

import (
	"github.com/go-co-op/gocron"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/ksusonic/owl-morning-bot/config"
	"github.com/ksusonic/owl-morning-bot/pkg/weather"
	"github.com/ksusonic/owl-morning-bot/pkg/weather/ya_weather"
	"log"
	"time"
)

type Scheduler struct {
	BotApi    *tgbotapi.BotAPI
	YaWeather *ya_weather.YaWeatherClient
	ChatId    int64
	Time      string
	Location  string
	Cron      *gocron.Scheduler
}

func NewScheduler(conf *config.Config, botApi *tgbotapi.BotAPI) *Scheduler {
	location, err := time.LoadLocation(conf.Scheduler.Location)
	if err != nil {
		log.Fatal("Incorrect location in config!")
	}
	return &Scheduler{
		BotApi:    botApi,
		YaWeather: ya_weather.NewYaWeatherClient(&conf.YaWeather),
		ChatId:    conf.Scheduler.ChatId,
		Time:      conf.Scheduler.Time,
		Location:  conf.Scheduler.Location,
		Cron:      gocron.NewScheduler(location),
	}
}

func (s *Scheduler) MakeCronTasks() {
	_, err := s.Cron.Every(1).Day().At(s.Time).Do(func() {
		if err := s.SendGoodMorning(); err != nil {
			log.Println("Error sending scheduled message: ", err)
		}
	})
	if err != nil {
		log.Println("Scheduler could not Do: ", err)
	}
	s.Cron.StartBlocking()
}

func (s *Scheduler) SendGoodMorning() error {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("recovered: %v\n", r)
		}
	}()
	text := RandomGoodMorningText()
	image := RandomGoodMorningImage()

	_, err := s.BotApi.Send(
		tgbotapi.NewMessage(s.ChatId, text))
	if err != nil {
		return err
	}
	_, err = s.BotApi.Send(tgbotapi.NewPhoto(s.ChatId, tgbotapi.FileURL(image)))
	if err != nil {
		log.Printf("Probably, bad image url: ")
	}
	log.Println("Sent:", text, "with image:", image)

	log.Println("getting Kazan weather info...")
	weatherNow := s.YaWeather.Request(&weather.Kazan{})
	if weatherNow != "" {
		_, _ = s.BotApi.Send(tgbotapi.NewMessage(s.ChatId, weatherNow))
	}

	return nil
}
