package constans

import tba "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var (
	NumericKeyboard1 = tba.NewInlineKeyboardMarkup(
		tba.NewInlineKeyboardRow(
			tba.NewInlineKeyboardButtonData("Свой пароль", "/yourpassword"),
			tba.NewInlineKeyboardButtonData("Сгенерировать пароль", "/generatepassword"),
		),
	)

	NumericKeyboard2 = tba.NewInlineKeyboardMarkup(
		tba.NewInlineKeyboardRow(
			tba.NewInlineKeyboardButtonData("8 символов", "/chars8"),
		),
		tba.NewInlineKeyboardRow(
			tba.NewInlineKeyboardButtonData("16 символов", "/chars16"),
		),
		tba.NewInlineKeyboardRow(
			tba.NewInlineKeyboardButtonData("32 символа", "/chars32"),
		),
	)

	NumericKeyboard3 = tba.NewInlineKeyboardMarkup(
		tba.NewInlineKeyboardRow(
			tba.NewInlineKeyboardButtonData("Английский алфавит", "/englishalphabet"),
		),
		tba.NewInlineKeyboardRow(
			tba.NewInlineKeyboardButtonData("Русский алфавит", "/russianalphabet"),
		),
	)

	NumericKeyboard4 = tba.NewInlineKeyboardMarkup(
		tba.NewInlineKeyboardRow(
			tba.NewInlineKeyboardButtonData("Строчные, пропичные символы и цифры", "/qwe"),
		),
		tba.NewInlineKeyboardRow(
			tba.NewInlineKeyboardButtonData("Строчные, пропичные, специальные символы и цифры", "/chars16"),
		),
	)
)
var (
	NumericKeyboard20 = tba.NewReplyKeyboard(
		tba.NewKeyboardButtonRow(
			tba.NewKeyboardButton("Да"),
			tba.NewKeyboardButton("Нет"),
		),
	)

	NumericKeyboard21 = tba.NewReplyKeyboard(
		tba.NewKeyboardButtonRow(
			tba.NewKeyboardButton("8 символов"),
		),
		tba.NewKeyboardButtonRow(
			tba.NewKeyboardButton("16 символов"),
		),
		tba.NewKeyboardButtonRow(
			tba.NewKeyboardButton("32 символов"),
		),
	)
	NumericKeyboard22 = tba.NewReplyKeyboard(
		tba.NewKeyboardButtonRow(
			tba.NewKeyboardButton("Английский алфавит"),
		),
		tba.NewKeyboardButtonRow(
			tba.NewKeyboardButton("Русский алфавит"),
		),
	)
)
