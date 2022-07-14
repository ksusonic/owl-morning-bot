package telegram

import (
	"github.com/ksusonic/owl-morning-bot/pkg/scheduler"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type Config struct {
	Bot struct {
		Debug      bool   `yaml:"debug"`
		UseWebhook bool   `yaml:"use_webhook"`
		WebhookUrl string `yaml:"webhook_url,omitempty"`
	}
	Scheduler scheduler.Scheduler `yaml:"scheduler"`
}

func Load(path string) *Config {
	var config = new(Config)
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(file, config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}
