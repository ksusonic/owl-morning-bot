package telegram

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	Bot struct {
		Debug      bool   `yaml:"debug"`
		UseWebhook bool   `yaml:"use_webhook"`
		WebhookUrl string `yaml:"webhook_url,omitempty"`
	}
}

func Load(path string) (*Config, error) {
	var config = new(Config)
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(file, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
