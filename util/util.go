package util

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
)

func FetchWeather(lat, lon float64) (body WeatherAPI) {
	Log("Building weather url query")
	wurl := viper.GetString("climateapiurl")
	i, _ := url.Parse(wurl)
	qry := i.Query()
	qry.Add("lat", fmt.Sprintf("%f", lat))
	qry.Add("lon", fmt.Sprintf("%f", lon))
	qry.Add("units", viper.GetString("units"))
	qry.Add("appid", viper.GetString("climateapikey"))
	i.RawQuery = qry.Encode()
	wurl = i.String()

	Log("Creating HTTP Client")
	wc := http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequest(http.MethodGet, wurl, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to create request:", err.Error())
	}

	Log("Sending Request")
	res, err := wc.Do(req)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to get request:", err.Error())
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	Log("Decoding JSON Response")
	decoder := json.NewDecoder(res.Body)
	data := WeatherAPI{}
	decoder.Decode(&data)

	Log("Returning data to caller")
	return data
}

func FetchWeatherByZip(zip, code string) (body WeatherAPI) {
	Log("Building weather url query")
	wurl := viper.GetString("climateapiurl")
	i, _ := url.Parse(wurl)
	qry := i.Query()
	qry.Add("zip", fmt.Sprintf("%s,%s", zip, code))
	qry.Add("units", viper.GetString("units"))
	qry.Add("appid", viper.GetString("climateapikey"))
	i.RawQuery = qry.Encode()
	wurl = i.String()
	wurl = strings.Replace(wurl, "%2C", ",", -1)
	Log(wurl)

	Log("Creating HTTP Client")
	wc := http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequest(http.MethodGet, wurl, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to create request:", err.Error())
	}

	Log("Sending Request")
	res, err := wc.Do(req)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to get request:", err.Error())
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	Log("Decoding JSON Response")
	decoder := json.NewDecoder(res.Body)
	data := WeatherAPI{}
	decoder.Decode(&data)

	Log("Returning data to caller")
	return data
}

func Log(str string) {
	if viper.GetBool("debugapp") {
		fmt.Println(str)
	}
}
