package main

import (
	cfg "demo-weather/config"
	reader2 "demo-weather/reader"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Weather application started")

	config, err := cfg.NewConfig()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", config)

	reader := strings.NewReader("Hello World")
	res, _ := reader2.ReadChunks(reader, 4)
	fmt.Println(string(res))
}
