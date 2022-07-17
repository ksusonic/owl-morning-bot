package test

import (
	"github.com/ksusonic/owl-morning-bot/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig(t *testing.T) {
	cfg := config.Load("../config/dev.yaml")
	assert.Equal(t, cfg.Bot, config.BotConfig{Version: "dev", Debug: true, UseWebhook: false})
	assert.Equal(t, cfg.Scheduler, config.SchedulerConfig{ChatId: 34912353, Time: cfg.Scheduler.Time, Location: "Europe/Moscow"})
	assert.Equal(t, cfg.YaWeather, config.YaWeatherConfig{Url: "https://api.weather.yandex.ru/v2/forecast", Lang: "ru_RU"})
}
