package bot

import (
	"log"
	"telegram-bot/internal/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func Start(cfg *config.Config) error {
	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		return err
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		response := "Привет! Я ваш Telegram-бот. Как я могу помочь?"
		if update.Message.Text == "Привет" {
			response = "Привет!"
		} else if update.Message.Text == "Пока" {
			response = "До свидания!"
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
		bot.Send(msg)
	}

	return nil
}
