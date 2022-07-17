package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"time"
)

type Config struct {
	Bot       BotConfig       `yaml:"bot"`
	Scheduler SchedulerConfig `yaml:"scheduler"`
	YaWeather YaWeatherConfig `yaml:"ya_weather"`
}

type BotConfig struct {
	Debug      bool   `yaml:"debug"`
	UseWebhook bool   `yaml:"use_webhook"`
	WebhookUrl string `yaml:"webhook_url,omitempty"`
}

type SchedulerConfig struct {
	ChatId   int64  `yaml:"chat_id"`
	Time     string `yaml:"time"`
	Location string `yaml:"location"`
}

type YaWeatherConfig struct {
	Url  string `yaml:"url"`
	Lang string `yaml:"lang"`
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
	if err := Validate(config); err != nil {
		log.Fatal(err)
	}
	return config
}

func Validate(c *Config) error {
	// Bot config check
	if c.Bot.UseWebhook == true && c.Bot.WebhookUrl == "" {
		return fmt.Errorf("use webhook is true and url is empty")
	}
	// Scheduler config check
	if _, err := time.LoadLocation(c.Scheduler.Location); err != nil {
		return err
	}
	if incorrectTime := c.Scheduler.Time == ""; incorrectTime || c.Scheduler.ChatId == 0 {
		if incorrectTime {
			return fmt.Errorf("incorrect time format in scheduler config")
		} else {
			return fmt.Errorf("incorrect chatId in scheduler config")
		}
	}
	return nil
}
