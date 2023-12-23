package constans

import tba "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var (
	MyCommands = []tba.BotCommand{
		{
			Command:     "/newservice",
			Description: "Создать новую запись с сервисом",
		},
		{
			Command:     "/myservices",
			Description: "Создать новую запись с сервисом",
		},
		{
			Command:     "/editservices",
			Description: "Создать новую запись с сервисом",
		},
		{
			Command:     "/generatepassword",
			Description: "Генерация пароля",
		},
		{
			Command:     "/myservices",
			Description: "Вывод сервисов",
		},
	}
)
