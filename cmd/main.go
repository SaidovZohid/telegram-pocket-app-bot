package main

import (
	"github.com/SaidovZohid/telegram-pocket-app-bot/pkg/telegram"
	"github.com/zhashkevych/go-pocket-sdk"

	"github.com/SaidovZohid/telegram-pocket-app-bot/pkg/logger"

	"github.com/SaidovZohid/telegram-pocket-app-bot/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	cfg := config.ParseConfig()
	logger.Init()
	log := logger.GetLogger()
	bot, err := tgbotapi.NewBotAPI(cfg.TelegramApiToken)
	if err != nil {
		log.Fatalf("Failed to make new bot api: %v", err)
	}

	pocketClient, err := pocket.NewClient(cfg.PocketDesktopOther)
	if err != nil {
		log.Fatalf("Failed to get client pocket: %v", err)
	}

	telegramBot := telegram.NewBot(bot, &log, pocketClient)

	if err := telegramBot.Start(); err != nil {
		log.Errorf("Error while starting bot: %v", err)
	}
}
