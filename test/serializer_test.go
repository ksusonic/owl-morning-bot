package test

import (
	"encoding/json"
	"github.com/ksusonic/owl-morning-bot/pkg/ya_weather"
	"io/ioutil"
	"os"
	"testing"
)

func TestSerializeResponse(t *testing.T) {
	file, err := os.Open("static/test_response.json")
	if err != nil {
		return
	}
	defer func() {
		_ = file.Close()
	}()
	byteValue, _ := ioutil.ReadAll(file)

	response := ya_weather.YaWeatherResponse{}
	err = json.Unmarshal(byteValue, &response)
	if err != nil {
		t.Fatal("Error unmarshalling file!")
	}
}
