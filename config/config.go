package config

import (
	"errors"
	"flag"
	"slices"
	"sync"
)

type Config struct {
	city   string
	format int
	Error  error
}

var availableFormats = []int{1, 2, 3, 4}

func NewConfig() Config {
	city := flag.String("city", "", "Город для прогноза погоды")
	format := flag.Int("format", 1, "Формат вывода погоды")

	if !slices.Contains(availableFormats, *format) {
		return Config{Error: errors.New("Формат должен быть в диапазоне от 1 до 4 включительно")}
	}

	return Config{city: *city, format: *format, Error: nil}
}

var AppConfig = sync.OnceValue(NewConfig)
