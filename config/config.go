package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	TelegramApiToken   string
	PocketDesktopOther string
}

func ParseConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Print(err)
	}

	conf := viper.New()
	conf.AutomaticEnv()

	return Config{
		TelegramApiToken:   conf.GetString("TELEGRAM_APITOKEN"),
		PocketDesktopOther: conf.GetString("POCKET_DESKTOP_OTHER"),
	}
}
