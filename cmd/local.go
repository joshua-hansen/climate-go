/*
Copyright © 2022 Joshua Hansen joshuahansen@outlook.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/joshua-hansen/climate-go/util"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// localCmd represents the local command
var localCmd = &cobra.Command{
	Use:   "local",
	Short: "Gets your local weather data",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Fetching Local Weather Data...")
		if viper.GetBool("debugapp") {
			fmt.Fprintln(os.Stdout, "weather-api-key:", viper.Get("climateapikey"))
		}
		body := fetchweather()

		var unit = '?'
		if viper.GetString("units") == "imperial" {
			unit = 'F'
		} else if viper.GetString("units") == "metric" {
			unit = 'C'
		} else {
			unit = 'K'
		}
		fmt.Fprintf(os.Stdout, "Temp: %v\u00B0%c Feel: %v\u00B0%c Max: %v\u00B0%c Min: %v\u00B0%c\n",
			body.Main.Temp, unit, body.Main.Feel, unit, body.Main.TMax, unit, body.Main.TMin, unit)
	},
}

func fetchweather() (body util.WeatherAPI) {
	wurl := viper.GetString("climateapiurl")
	i, _ := url.Parse(wurl)
	qry := i.Query()
	qry.Add("lat", fmt.Sprintf("%f", viper.GetFloat64("latitude")))
	qry.Add("lon", fmt.Sprintf("%f", viper.GetFloat64("longitude")))
	qry.Add("units", viper.GetString("units"))
	qry.Add("appid", viper.GetString("climateapikey"))
	i.RawQuery = qry.Encode()
	wurl = i.String()

	wc := http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequest(http.MethodGet, wurl, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to create request:", err.Error())
	}

	res, err := wc.Do(req)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to get request:", err.Error())
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	decoder := json.NewDecoder(res.Body)
	var data = util.WeatherAPI{}
	decoder.Decode(&data)

	return data
}

func init() {
	rootCmd.AddCommand(localCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// localCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// localCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
