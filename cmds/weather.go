package cmds

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

type WeatherResponse struct {
	CurrentCondition []struct {
		FeelsLikeC  string `json:"FeelsLikeC"`
		FeelsLikeF  string `json:"FeelsLikeF"`
		TempC       string `json:"temp_C"`
		TempF       string `json:"temp_F"`
		WeatherDesc []struct {
			Value string `json:"value"`
		} `json:"weatherDesc"`
	} `json:"current_condition"`
}

var weatherCmd = &cobra.Command{
	Use:   "weather [location]",
	Short: "Get the current weather for a specific location",
	Long:  `Fetches and displays the current weather conditions, including temperature, "feels like" temperature, and a description from the wttr.in API. Defaults to Celsius.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		location := args[0]
		url := fmt.Sprintf("https://wttr.in/%s?format=j1", location)

		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error fetching weather data:", err)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return
		}

		var weather WeatherResponse
		if err := json.Unmarshal(body, &weather); err != nil {
			fmt.Println("Error parsing JSON:", err)
			return
		}

		if len(weather.CurrentCondition) == 0 {
			fmt.Printf("Could not get weather for '%s'. Is it a valid location?\n", location)
			return
		}

		current := weather.CurrentCondition[0]
		description := ""
		if len(current.WeatherDesc) > 0 {
			description = current.WeatherDesc[0].Value
		}

		useFahrenheit, _ := cmd.Flags().GetBool("fahrenheit")

		if useFahrenheit {
			fmt.Printf("Weather in %s. Currently %s at %s°F but feels like %s°F.\n",
				strings.Title(location),
				description,
				current.TempF,
				current.FeelsLikeF)
		} else {
			// Default to Celsius
			fmt.Printf("Weather in %s. Currently %s at %s°C but feels like %s°C.\n",
				strings.Title(location),
				description,
				current.TempC,
				current.FeelsLikeC)
		}
	},
}

func init() {
	weatherCmd.Flags().BoolP("fahrenheit", "f", false, "Display temperature in Fahrenheit")

	rootCmd.AddCommand(weatherCmd)
}
