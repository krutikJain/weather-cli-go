/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// currentWeatherCmd represents the currentWeather command
var currentWeatherCmd = &cobra.Command{
	Use:   "currentWeather",
	Short: "CurrentWeather Of Location",
	Long:  `Gives You Information About Current Weather Of Location You Provided`,
	Run: func(cmd *cobra.Command, args []string) {
		getCurrentWeather()
	},
}

func init() {
	rootCmd.AddCommand(currentWeatherCmd)
}

type Weather struct {
	Location struct {
		Name string `json:"name"`
	} `json:"location"`
	CurrentData struct {
		Temperature float64 `json:"temp_c"`
	} `json:"current"`
}

func getCurrentWeather() {

	color.New(color.FgGreen).Println("Enter Location For Which You Want To Get Weather Info")

	var loc string

	fmt.Scan(&loc)

	resByte := getData("http://api.weatherapi.com/v1/current.json?key=YOUR_KEY&q=" + loc)

	weather := Weather{}

	err := json.Unmarshal(resByte, &weather)

	if err != nil {
		log.Println("Error Unmarshaling", err)
	}

	color.New(color.FgBlue).Printf("Current Weather For %s is %0.1fC", weather.Location.Name, weather.CurrentData.Temperature)
}

func getData(api string) []byte {

	req, err := http.NewRequest(
		http.MethodGet,
		api,
		nil,
	)

	if err != nil {
		log.Println("Request Error = ", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "Learning Purpose")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Println("Response Error = ", err)
	}

	resByte, err := io.ReadAll(res.Body)

	if err != nil {
		log.Println("Error in Reading Data = ", err)
	}

	return resByte
}
