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
	"fmt"
	"os"

	"github.com/joshua-hansen/climate-go/util"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// localCmd represents the local command
var localCmd = &cobra.Command{
	Use:   "local",
	Short: "Gets your local weather data",
	Long: `local gets weather data from a predefined location
	in the configuration file. Make the edit to latitude and longitude
	of your location.`,
	Run: func(cmd *cobra.Command, args []string) {
		if viper.GetBool("debugapp") {
			fmt.Fprintln(os.Stdout, "weather-api-key:", viper.Get("climateapikey"))
		}
		res := fetchweather()

		var unit = '?'
		if viper.GetString("units") == "imperial" {
			unit = 'F'
		} else if viper.GetString("units") == "metric" {
			unit = 'C'
		} else {
			unit = 'K'
		}
		fmt.Fprintf(os.Stdout, "Temp: %v\u00B0%c Feel: %v\u00B0%c Max: %v\u00B0%c Min: %v\u00B0%c\n",
			res.Main.Temp, unit, res.Main.Feel, unit, res.Main.TMax, unit, res.Main.TMin, unit)
	},
}

func fetchweather() (body util.WeatherAPI) {
	lat := viper.GetFloat64("latitude")
	lon := viper.GetFloat64("longitude")
	util.Log(fmt.Sprintf("Calling to fetch weather at lat: %f lon: %f ", lat, lon))
	return util.FetchWeather(lat, lon)
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
