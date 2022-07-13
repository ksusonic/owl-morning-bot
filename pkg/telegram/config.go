package telegram

import (
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
	Scheduler Scheduler `yaml:"scheduler"`
	Nlg       Nlg       `yaml:"nlg"`
}

type Scheduler struct {
	ChatId int64  `yaml:"chat_id"`
	Time   string `yaml:"time"`
}

type Nlg struct {
	GoodMorning []string `yaml:"good_morning"`
	Images      []string `yaml:"images"`
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
