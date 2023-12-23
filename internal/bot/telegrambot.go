package telegrambot

import (
	"github.com/ent1k1377/my-password-bot/internal/constans"
	"github.com/ent1k1377/my-password-bot/internal/handlers"
	sm "github.com/ent1k1377/my-password-bot/internal/statemachine"
	tba "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type HandlerFunc func(update tba.Update, bot *tba.BotAPI)

type TelegramBot struct {
	Bot          *tba.BotAPI
	UpdateChan   tba.UpdatesChannel
	statemachine *sm.StateMachine
	handlers     map[sm.State]HandlerFunc
}

func New(token string) (*TelegramBot, error) {
	bot, err := tba.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	bot.Debug = true

	update := tba.NewUpdate(0)
	update.Timeout = 60

	updates := bot.GetUpdatesChan(update)
	statemachine := sm.New()
	statemachine.Config()

	return &TelegramBot{
		Bot:          bot,
		UpdateChan:   updates,
		statemachine: statemachine,
		handlers:     make(map[sm.State]HandlerFunc),
	}, nil
}

func (tb *TelegramBot) Start() {
	for update := range tb.UpdateChan {
		if update.Message != nil {
			tb.updateMessage(update)
		} else if update.CallbackQuery != nil {
			tb.updateCallbackQuery(update)
		} else {
			log.Fatal("update is missing", update)
		}
	}
}

func (tb *TelegramBot) SetConfig() error {
	_, err := tb.setMyCommands()
	if err != nil {
		return err
	}

	tb.setHandlers()

	return nil
}

func (tb *TelegramBot) setMyCommands() (*tba.APIResponse, error) {
	commands := tba.SetMyCommandsConfig{
		Commands:     constans.MyCommands,
		Scope:        nil,
		LanguageCode: "",
	}
	return tb.Bot.Request(commands)
}

func (tb *TelegramBot) setHandlers() {
	tb.addHandler(sm.SendBotInfoMessage, handlers.SendBotInfoMessageHandler)
	tb.addHandler(sm.WaitingForServiceName, handlers.WaitingForServiceNameHandler)
}

func (tb *TelegramBot) addHandler(state sm.State, handler HandlerFunc) {
	tb.handlers[state] = handler
}
