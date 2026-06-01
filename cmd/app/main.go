package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/Elunded/lb1/internal/estimator" // ОБОВ'ЯЗКОВО ЗАМІНИ [твій-нік] НА ВЛАСНИЙ
)

func main() {
	// Налаштування Viper для читання config.yaml
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// Читаємо файл
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal().Err(err).Msg("Помилка читання конфіг-файлу")
	}

	// Налаштування Zerolog для красивого виводу в консоль
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Встановлюємо рівень логування з конфігу
	logLevelStr := viper.GetString("logger.level")
	logLevel, err := zerolog.ParseLevel(logLevelStr)
	if err != nil {
		logLevel = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(logLevel)

	log.Info().Msg("Запуск будівельного калькулятора")

	// Беремо площу цегли з файлу налаштувань
	brickArea := viper.GetFloat64("brick.area")
	wallArea := 50.0 // Припустимо, стіна 50 квадратів

	log.Debug().Float64("wall_area", wallArea).Float64("brick_area", brickArea).Msg("Вхідні дані для розрахунку")

	result, err := estimator.BricksNeeded(wallArea, brickArea)
	if err != nil {
		log.Error().Err(err).Msg("Помилка розрахунку")
		return
	}

	log.Info().Float64("bricks_needed", result).Msg("Успішний розрахунок матеріалів")
}
