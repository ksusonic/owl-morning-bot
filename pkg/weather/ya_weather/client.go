package ya_weather

import (
	"encoding/json"
	"fmt"
	"github.com/ksusonic/owl-morning-bot/config"
	"github.com/ksusonic/owl-morning-bot/pkg/weather"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type YaWeatherClient struct {
	ApiToken string
	URL      string
	Lang     string
}

func NewYaWeatherClient(c *config.YaWeatherConfig) *YaWeatherClient {
	token := os.Getenv("YA_WEATHER_TOKEN")
	if token == "" {
		log.Fatal("empty YA_WEATHER_TOKEN")
	}
	return &YaWeatherClient{ApiToken: token, URL: c.Url, Lang: c.Lang}
}

func (c *YaWeatherClient) Request(city weather.City) string {
	request := c.makeRequest(city)
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("ya_weather_client: could not request weather: %s\n", err)
		return ""
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("ya_weather_client: could not read response body: %s\n", err)
		return ""
	}

	if res.StatusCode != 200 {
		log.Printf("ya_weather_client: status is %d: %s\n", res.StatusCode, resBody)
		return ""
	}

	log.Printf("successful response: %s\n", resBody)

	parsed, err := c.parseResponse(res)
	if err != nil {
		log.Printf("ya_weather_client: could not parse weather: %s\n", err)
		return ""
	}

	return c.readableResponse(parsed, city)
}

func (c *YaWeatherClient) makeRequest(city weather.City) *http.Request {
	req, err := http.NewRequest("GET", c.URL, nil)
	if err != nil {
		log.Println(err)
	}

	req.Header.Add("X-Yandex-API-Key", c.ApiToken)

	q := req.URL.Query()
	q.Add("lat", city.Lat())
	q.Add("lon", city.Lon())
	q.Add("lang", c.Lang)
	q.Add("limit", "1")

	req.URL.RawQuery = q.Encode()
	log.Printf("request ready: %s\n", req.URL)
	return req
}

func (c *YaWeatherClient) parseResponse(res *http.Response) (*YaWeatherResponse, error) {
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := YaWeatherResponse{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *YaWeatherClient) readableResponse(res *YaWeatherResponse, city weather.City) string {
	var builder strings.Builder
	builder.WriteString(
		fmt.Sprintf(
			"Сейчас в %s %d градусов, %s. Ветер %dм/с\n",
			city.NameIn(),
			res.Fact.Temp,
			conditionParser(res.Fact.Condition),
			res.Fact.WindSpeed,
		),
	)
	builder.WriteString(
		fmt.Sprintf(
			"Днем %d, %s\n",
			res.Forecasts[0].Parts.DayShort.TempAvg,
			conditionParser(res.Forecasts[0].Parts.DayShort.Condition),
		),
	)
	builder.WriteString(
		fmt.Sprintf(
			"Вечером %d, %s\n",
			res.Forecasts[0].Parts.Evening.TempAvg,
			conditionParser(res.Forecasts[0].Parts.Evening.Condition),
		),
	)
	builder.WriteString(
		fmt.Sprintf(
			"Ночью %d, %s\n",
			res.Forecasts[0].Parts.Night.TempAvg,
			conditionParser(res.Forecasts[0].Parts.Night.Condition),
		),
	)
	return builder.String()
}

func conditionParser(condition string) string {
	switch condition {
	case "clear":
		return "ясно ☀️"
	case "partly-cloudy":
		return "малооблачно 🌤️"
	case "cloudy":
		return "облачно с прояснениями ⛅️️"
	case "overcast":
		return "пасмурно ☁️"
	case "drizzle":
		return "морось 🌨"
	case "lightrain":
		return "небольшой дождь 🌧"
	case "rain":
		return "дождь 🌧☔️"
	case "moderate-rain":
		return "умеренно сильный дождь 🌧☔️"
	case "heavy-rain":
		return "сильный дождь 🌧☔️"
	case "continuous-heavy-rain":
		return "длительный сильный дождь 🌧☔️"
	case "showers":
		return "ливень 🌧☔️"
	case "wet-snow":
		return "дождь со снегом 🌨"
	case "light-snow":
		return "небольшой снег ❄️"
	case "snow":
		return "снег ❄️"
	case "snow-showers":
		return "снегопад ❄️"
	case "hail":
		return "град 🌧"
	case "thunderstorm":
		return "гроза 🌩"
	case "thunderstorm-with-rain":
		return "дождь с грозой ⛈"
	case "thunderstorm-with-hail":
		return "гроза с градом ⛈"
	default:
		return ""
	}
}
