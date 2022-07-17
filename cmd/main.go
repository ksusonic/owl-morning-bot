package main

import (
	"flag"
	"github.com/ksusonic/owl-morning-bot/config"
	"github.com/ksusonic/owl-morning-bot/pkg/scheduler"
	"github.com/ksusonic/owl-morning-bot/pkg/telegram"
)

func main() {
	configPathPtr := flag.String("c", "config/dev.yaml", "Path to config")
	flag.Parse()

	cfg := config.Load(*configPathPtr)

	teleBot := telegram.NewBot(cfg)

	cronScheduler := scheduler.NewScheduler(cfg, teleBot.Bot)
	go cronScheduler.MakeCronTasks()

	teleBot.Start()
}
