package handlers

import (
	"github.com/ent1k1377/my-password-bot/internal/constans"
	tba "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func UnknownMessageHandler(update tba.Update, bot *tba.BotAPI) {
	msg := tba.NewMessage(update.Message.From.ID, constans.UnknownMessage)
	bot.Send(msg)
}

func SendBotInfoMessageHandler(update tba.Update, bot *tba.BotAPI) {
	msg := tba.NewMessage(update.Message.From.ID, constans.StartCommand)
	bot.Send(msg)
}

func NewServiceHandler(update tba.Update, bot *tba.BotAPI) {
	msg := tba.NewMessage(update.Message.From.ID, constans.NameService)
	bot.Send(msg)
}

func WaitingForServiceNameHandler(update tba.Update, bot *tba.BotAPI) {
	msg := tba.NewMessage(update.Message.From.ID, constans.NamePassword)
	msg.ReplyMarkup = constans.NumericKeyboard20
	bot.Send(msg)
}

func PasswordHandler(update tba.Update, bot *tba.BotAPI) {
	msg := tba.NewMessage(update.Message.From.ID, constans.ConfigurePassword)
	msg.ReplyMarkup = constans.NumericKeyboard21
	bot.Send(msg)
}
