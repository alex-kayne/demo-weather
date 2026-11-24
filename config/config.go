package config

import (
	"errors"
	"flag"
	"slices"
)

type Config struct {
	City   string
	Format int
}

var availableFormats = []int{1, 2, 3, 4}

func NewConfig() (Config, error) {
	city := flag.String("city", "", "Город для прогноза погоды")
	format := flag.Int("format", 1, "Формат вывода погоды")

	flag.Parse()

	if !slices.Contains(availableFormats, *format) {
		return Config{}, errors.New("формат должен быть в диапазоне от 1 до 4 включительно")
	}

	return Config{City: *city, Format: *format}, nil
}
