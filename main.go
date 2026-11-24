package main

import (
	cfg "demo-weather/config"
	"fmt"
)

func main() {
	fmt.Println("Weather application started")

	config, err := cfg.NewConfig()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", config)

}
