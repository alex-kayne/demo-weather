package main

import (
	cfg "demo-weather/config"
	"fmt"
)

func main() {
	fmt.Println("Weather application started")

	config := cfg.AppConfig()

	if config.Error != nil {
		fmt.Println(config.Error)
		return
	}

}
