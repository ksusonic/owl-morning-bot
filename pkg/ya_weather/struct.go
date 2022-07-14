package ya_weather

import "time"

type YaWeatherResponse struct {
	Now   int64     `json:"now"`
	NowDt time.Time `json:"now_dt"`
	Info  struct {
		Lat    float64 `json:"lat"`
		Lon    float64 `json:"lon"`
		Tzinfo struct {
			Offset int    `json:"offset"`
			Name   string `json:"name"`
			Abbr   string `json:"abbr"`
			Dst    bool   `json:"dst"`
		} `json:"tzinfo"`
		DefPressureMm int    `json:"def_pressure_mm"`
		DefPressurePa int    `json:"def_pressure_pa"`
		Url           string `json:"url"`
	} `json:"info"`
	Fact struct {
		Temp            int     `json:"temp"`
		FeelsLike       int     `json:"feels_like"`
		Icon            string  `json:"icon"`
		Condition       string  `json:"condition"`
		WindSpeed       int     `json:"wind_speed"`
		WindGust        float64 `json:"wind_gust"`
		WindDir         string  `json:"wind_dir"`
		PressureMm      int     `json:"pressure_mm"`
		PressurePa      int     `json:"pressure_pa"`
		Humidity        int     `json:"humidity"`
		Daytime         string  `json:"daytime"`
		Polar           bool    `json:"polar"`
		Season          string  `json:"season"`
		PrecType        int     `json:"prec_type"`
		PrecStrength    float64 `json:"prec_strength"`
		IsThunder       bool    `json:"is_thunder"`
		Cloudness       int     `json:"cloudness"`
		ObsTime         int     `json:"obs_time"`
		PhenomIcon      string  `json:"phenom_icon"`
		PhenomCondition string  `json:"phenom-condition"`
	} `json:"fact"`
	Forecasts []struct {
		Date     string `json:"date"`
		DateTs   int    `json:"date_ts"`
		Week     int    `json:"week"`
		Sunrise  string `json:"sunrise"`
		Sunset   string `json:"sunset"`
		MoonCode int    `json:"moon_code"`
		MoonText string `json:"moon_text"`
		Parts    struct {
			Night struct {
				TempMin      int     `json:"temp_min"`
				TempMax      int     `json:"temp_max"`
				TempAvg      int     `json:"temp_avg"`
				FeelsLike    int     `json:"feels_like"`
				Icon         string  `json:"icon"`
				Condition    string  `json:"condition"`
				Daytime      string  `json:"daytime"`
				Polar        bool    `json:"polar"`
				WindSpeed    float64 `json:"wind_speed"`
				WindGust     int     `json:"wind_gust"`
				WindDir      string  `json:"wind_dir"`
				PressureMm   int     `json:"pressure_mm"`
				PressurePa   int     `json:"pressure_pa"`
				Humidity     int     `json:"humidity"`
				PrecMm       int     `json:"prec_mm"`
				PrecPeriod   int     `json:"prec_period"`
				PrecType     int     `json:"prec_type"`
				PrecStrength int     `json:"prec_strength"`
				Cloudness    float64 `json:"cloudness"`
			} `json:"night"`
			Evening struct {
				TempMin      int     `json:"temp_min"`
				TempMax      int     `json:"temp_max"`
				TempAvg      int     `json:"temp_avg"`
				FeelsLike    int     `json:"feels_like"`
				Icon         string  `json:"icon"`
				Condition    string  `json:"condition"`
				Daytime      string  `json:"daytime"`
				Polar        bool    `json:"polar"`
				WindSpeed    float64 `json:"wind_speed"`
				WindDir      string  `json:"wind_dir"`
				PressureMm   int     `json:"pressure_mm"`
				PressurePa   int     `json:"pressure_pa"`
				Humidity     int     `json:"humidity"`
				PrecMm       float64 `json:"prec_mm"`
				PrecPeriod   int     `json:"prec_period"`
				PrecType     int     `json:"prec_type"`
				PrecStrength float64 `json:"prec_strength"`
				Cloudness    float64 `json:"cloudness"`
			} `json:"evening"`
			DayShort struct {
				Temp         int     `json:"temp"`
				TempMin      int     `json:"temp_min"`
				FeelsLike    int     `json:"feels_like"`
				Icon         string  `json:"icon"`
				Condition    string  `json:"condition"`
				WindSpeed    float64 `json:"wind_speed"`
				WindGust     float64 `json:"wind_gust"`
				WindDir      string  `json:"wind_dir"`
				PressureMm   int     `json:"pressure_mm"`
				PressurePa   int     `json:"pressure_pa"`
				Humidity     int     `json:"humidity"`
				PrecType     int     `json:"prec_type"`
				PrecStrength float64 `json:"prec_strength"`
				Cloudness    int     `json:"cloudness"`
			} `json:"day_short"`
			NightShort struct {
				Temp         int     `json:"temp"`
				FeelsLike    int     `json:"feels_like"`
				Icon         string  `json:"icon"`
				Condition    string  `json:"condition"`
				WindSpeed    float64 `json:"wind_speed"`
				WindGust     float64 `json:"wind_gust"`
				WindDir      string  `json:"wind_dir"`
				PressureMm   int     `json:"pressure_mm"`
				PressurePa   int     `json:"pressure_pa"`
				Humidity     int     `json:"humidity"`
				PrecType     int     `json:"prec_type"`
				PrecStrength int     `json:"prec_strength"`
				Cloudness    float64 `json:"cloudness"`
			} `json:"night_short"`
		} `json:"parts"`
		Hours []struct {
			Hour         string  `json:"hour"`
			HourTs       int     `json:"hour_ts"`
			Temp         int     `json:"temp"`
			FeelsLike    int     `json:"feels_like"`
			Icon         string  `json:"icon"`
			Condition    string  `json:"condition"`
			WindSpeed    float64 `json:"wind_speed"`
			WindGust     float64 `json:"wind_gust"`
			WindDir      string  `json:"wind_dir"`
			PressureMm   int     `json:"pressure_mm"`
			PressurePa   int     `json:"pressure_pa"`
			Humidity     int     `json:"humidity"`
			PrecMm       float64 `json:"prec_mm"`
			PrecPeriod   int     `json:"prec_period"`
			PrecType     int     `json:"prec_type"`
			PrecStrength float64 `json:"prec_strength"`
			IsThunder    bool    `json:"is_thunder"`
			Cloudness    float64 `json:"cloudness"`
		} `json:"hours"`
	} `json:"forecasts"`
}
