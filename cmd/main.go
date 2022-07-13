package main

import (
	"flag"
	"github.com/ksusonic/owl-morning-bot/pkg/telegram"
	"log"
)

func main() {
	configPathPtr := flag.String("c", "config/dev.yaml", "Path to config")
	flag.Parse()

	cfg, err := telegram.Load(*configPathPtr)
	if err != nil {
		log.Fatal(err)
	}

	teleBot := telegram.NewBot(cfg)
	teleBot.Start()
}
