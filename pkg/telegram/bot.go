package telegram

import (
	"github.com/SaidovZohid/telegram-pocket-app-bot/pkg/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/zhashkevych/go-pocket-sdk"
)

type Bot struct {
	bot          *tgbotapi.BotAPI
	logger       *logger.Logger
	pocketClient *pocket.Client
}

func NewBot(bot *tgbotapi.BotAPI, logger *logger.Logger, pocketClient *pocket.Client) *Bot {
	return &Bot{
		bot:          bot,
		logger:       logger,
		pocketClient: pocketClient,
	}
}

func (b *Bot) Start() error {
	b.logger.Infof("Authired on Account: %v", b.bot.Self.UserName)

	updates := b.initUpdatesChannel()

	b.handleUpdates(updates)

	return nil
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			b.handleCommand(update.Message)
			continue
		}

		b.handleMessage(update.Message)
	}
}

func (b *Bot) initUpdatesChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u)
}
