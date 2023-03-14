package telegram

import (
	"github.com/SaidovZohid/telegram-pocket-app-bot/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) handleMessage(message *tgbotapi.Message) {
	b.logger.Infof("[%s] -> %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)

	b.bot.Send(msg)
}

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case config.CommandStart:
		return b.handleStartCommand(message)
	default:
		return b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, config.ReplyCommand)
	_, err := b.bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "wrong command :(")
	_, err := b.bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}
