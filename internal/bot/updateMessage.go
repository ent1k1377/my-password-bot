package telegrambot

import (
	sm "github.com/ent1k1377/my-password-bot/internal/statemachine"
	tba "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

func (tb *TelegramBot) updateMessage(update tba.Update) {
	msg := update.Message.Text

	currentEvent := tb.currentEvent(msg)
	currentState := tb.statemachine.CurrentState
	nextState := tb.statemachine.Transaction[currentState][currentEvent]

	tb.handlers[nextState](update, tb.Bot)
	if tb.statemachine.Transaction[nextState] == nil {
		tb.statemachine.CurrentState = sm.None
	} else {
		tb.statemachine.CurrentState = nextState
	}
}

func (tb *TelegramBot) currentEvent(msg string) sm.Event {
	var event sm.Event
	switch msg {
	case "/start":
		event = sm.Start
	case "/newservice":
		event = sm.NewService
	case "/generatepassword":
		event = sm.GeneratePassword
	default:
		isCommand := strings.HasPrefix(msg, "/")
		if isCommand {
			event = sm.UnknownCommand
		} else {
			event = sm.CustomMessage
		}
	}

	return event
}
