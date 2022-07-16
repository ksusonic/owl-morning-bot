package scheduler

import (
	"github.com/go-co-op/gocron"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"time"
)

type Scheduler struct {
	ChatId   int64  `yaml:"chat_id"`
	Time     string `yaml:"time"`
	Location string `yaml:"location"`

	Cron *gocron.Scheduler `yaml:"-"`
}

func NewScheduler(shedConf *Scheduler) *Scheduler {
	location, err := time.LoadLocation(shedConf.Location)
	if err != nil {
		log.Fatal("Incorrect location in config!")
	}
	return &Scheduler{
		ChatId:   shedConf.ChatId,
		Time:     shedConf.Time,
		Location: shedConf.Location,
		Cron:     gocron.NewScheduler(location),
	}
}

func (s *Scheduler) MakeCronTasks(b *tgbotapi.BotAPI) {
	_, err := s.Cron.Every(1).Day().At(s.Time).Do(func() {
		if err := s.SendGoodMorning(b); err != nil {
			log.Println("Error sending scheduled message: ", err)
		}
	})
	if err != nil {
		log.Println("Scheduler could not Do: ", err)
	}
	s.Cron.StartBlocking()
}

func (s *Scheduler) SendGoodMorning(b *tgbotapi.BotAPI) error {
	text := RandomGoodMorningText()
	image := RandomGoodMorningImage()

	_, err := b.Send(
		tgbotapi.NewMessage(s.ChatId, text))
	if err != nil {
		return err
	}
	_, err = b.Send(tgbotapi.NewPhoto(s.ChatId, tgbotapi.FileURL(image)))
	if err != nil {
		log.Printf("Probably, bad image url: ")
	}
	log.Println("Sent:", text, "with image:", image)
	return nil
}
