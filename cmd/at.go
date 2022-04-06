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
	"strconv"

	"github.com/joshua-hansen/climate-go/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// atCmd represents the at command
var atCmd = &cobra.Command{
	Use:   "at",
	Short: "Gets weather location at the provided location",
	Long: `at gets weather location at the provided location.
	
	So far the only argument supported is zip code
	
	US is the default country code. You can add a country code
	after you provide the zip code.
	
	Usage:
	climate at [zip] [country-code]`,
	Run: func(cmd *cobra.Command, args []string) {
		if viper.GetBool("debugapp") {
			fmt.Fprintln(os.Stdout, "weather-api-key:", viper.Get("climateapikey"))
			fmt.Fprintln(os.Stdout, "arguments: ", args)
		}

		if !argIntegrityCheck(args) {
			fmt.Println("Terminating....")
			os.Exit(1)
		}
		var zip, _ = strconv.Atoi(args[0])
		var code string
		if len(args) == 2 {
			code = args[1]
		}
		res := fetchweatherbyzip(zip, code)

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

func fetchweatherbyzip(zip int, code string) (body util.WeatherAPI) {
	util.Log(fmt.Sprintf("Calling to fetch weather at zip: %v and country: %v", zip, code))
	return util.FetchWeatherByZip(zip)
}

func argIntegrityCheck(args []string) bool {
	argLen := len(args)
	var pass = true
	if argLen != 1 && argLen != 2 {
		fmt.Println("Please check inputs, refer to help section.")
		pass = false
	}
	if _, err := strconv.Atoi(args[0]); err != nil {
		fmt.Println("Please check zipcode is an integer")
		pass = false
	}
	var c string
	if argLen == 2 {
		c = args[1]
		if len(c) != 2 {
			fmt.Println("Country code should be two characters")
			pass = false
		}
	}
	return pass
}

func init() {
	rootCmd.AddCommand(atCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// atCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// atCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
