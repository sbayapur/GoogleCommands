package p

import (
	"fmt"
	"io/ioutil"
	"net/http"
    "time"
    //"bytes"
	"encoding/json"
    //"log"
)
type Weather struct {
	ConsolidatedWeather []struct {
		ID                   int64     `json:"id"`
		WeatherStateName     string    `json:"weather_state_name"`
		WeatherStateAbbr     string    `json:"weather_state_abbr"`
		WindDirectionCompass string    `json:"wind_direction_compass"`
		Created              time.Time `json:"created"`
		ApplicableDate       string    `json:"applicable_date"`
		MinTemp              float64   `json:"min_temp"`
		MaxTemp              float64   `json:"max_temp"`
		TheTemp              float64   `json:"the_temp"`
		WindSpeed            float64   `json:"wind_speed"`
		WindDirection        float64   `json:"wind_direction"`
		AirPressure          float64   `json:"air_pressure"`
		Humidity             int       `json:"humidity"`
		Visibility           float64   `json:"visibility"`
		Predictability       int       `json:"predictability"`
	} `json:"consolidated_weather"`
	Time         time.Time `json:"time"`
	SunRise      time.Time `json:"sun_rise"`
	SunSet       time.Time `json:"sun_set"`
	TimezoneName string    `json:"timezone_name"`
	Parent       struct {
		Title        string `json:"title"`
		LocationType string `json:"location_type"`
		Woeid        int    `json:"woeid"`
		LattLong     string `json:"latt_long"`
	} `json:"parent"`
	Sources []struct {
		Title     string `json:"title"`
		Slug      string `json:"slug"`
		URL       string `json:"url"`
		CrawlRate int    `json:"crawl_rate"`
	} `json:"sources"`
	Title        string `json:"title"`
	LocationType string `json:"location_type"`
	Woeid        int    `json:"woeid"`
	LattLong     string `json:"latt_long"`
	Timezone     string `json:"timezone"`
}

func GetWeather(w http.ResponseWriter, r *http.Request) {
	apiurl := "https://www.metaweather.com/api/location/44418"
    client := &http.Client{}
    req, err := http.NewRequest("GET", apiurl,nil)
    if err != nil {
        fmt.Fprintf(w, "Error!!")
    }
    res, err := client.Do(req)
    defer res.Body.Close()
    if err != nil {
        fmt.Fprintf(w, "Error!!!")
    }
    body, _ := ioutil.ReadAll(res.Body)
	weather := &Weather{}
	json.Unmarshal(body, weather)
    
    final_response := "The weather in London today is " + weather.ConsolidatedWeather[0].WeatherStateName

    fmt.Fprintf(w,final_response)
}
