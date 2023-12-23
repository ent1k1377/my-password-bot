package telegrambot

import tba "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (tb *TelegramBot) updateCallbackQuery(update tba.Update) {
	callback := tba.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
	if _, err := tb.Bot.Request(callback); err != nil {
		panic(err)
	}

	msg := tba.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
	if _, err := tb.Bot.Send(msg); err != nil {
		panic(err)
	}
}
