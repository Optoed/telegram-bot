# Telegram Bot

Этот проект представляет собой Telegram-бот, написанный на Go. Бот отвечает на сообщения "Привет" и "Пока".

## Запуск

1. Установите зависимости:
    ```sh
    go mod tidy
    ```

2. Создайте файл `.env` и добавьте ваш Telegram Bot Token:
    ```
    TELEGRAM_TOKEN=YOUR_TELEGRAM_BOT_TOKEN
    ```

3. Запустите бота:
    ```sh
    go run cmd/bot/main.go
    ```
