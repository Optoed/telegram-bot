package main

import (
	"log"
	"telegram-bot/internal/bot"
	"telegram-bot/internal/config"
)

func main() {
	// Загрузка конфигурации
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	//Запуск бота
	if err := bot.Start(cfg); err != nil {
		log.Fatalf("Error starting bot: %v", err)
	}
}
