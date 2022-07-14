package main

import (
	"flag"
	"github.com/ksusonic/owl-morning-bot/pkg/scheduler"
	"github.com/ksusonic/owl-morning-bot/pkg/telegram"
)

func main() {
	configPathPtr := flag.String("c", "config/dev.yaml", "Path to config")
	flag.Parse()

	cfg := telegram.Load(*configPathPtr)

	teleBot := telegram.NewBot(cfg)

	cronScheduler := scheduler.NewScheduler(&cfg.Scheduler)
	go cronScheduler.MakeCronTasks(teleBot.Bot)

	teleBot.Start()
}
