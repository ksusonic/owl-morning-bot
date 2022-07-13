package scheduler

import (
	"github.com/go-co-op/gocron"
	"github.com/ksusonic/owl-morning-bot/pkg/telegram"
	"log"
	"time"
)

func MakeCronTasks(b *telegram.Bot) {
	location, _ := time.LoadLocation("Europe/Moscow")
	s := gocron.NewScheduler(location)
	_, err := s.Every(1).Day().At(b.ShConf().Time).Do(func() {
		if err := b.SendScheduledMessage(); err != nil {
			log.Println("Error sending scheduled message: ", err)
		}
	})
	if err != nil {
		log.Println("Scheduler could not Do: ", err)
	}
	s.StartBlocking()
}
