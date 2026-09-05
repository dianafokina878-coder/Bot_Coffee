package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func NewQuestionKeyboard(question Question) tgbotapi.InlineKeyboardMarkup {
	button1 := tgbotapi.NewInlineKeyboardButtonData(
		question.Answer1,
		"answer1",
	)
	button2 := tgbotapi.NewInlineKeyboardButtonData(
		question.Answer2,
		"answer2",
	)
	row1 := tgbotapi.NewInlineKeyboardRow(button1)
	row2 := tgbotapi.NewInlineKeyboardRow(button2)
	keyboard := tgbotapi.NewInlineKeyboardMarkup(row1, row2)
	return keyboard
}

func NewCoffeeShopKeyboard() tgbotapi.InlineKeyboardMarkup {
	button1 := tgbotapi.NewInlineKeyboardButtonData(
		"Вперёд",
		"next",
	)

	button2 := tgbotapi.NewInlineKeyboardButtonData(
		"Назад",
		"back",
	)

	row1 := tgbotapi.NewInlineKeyboardRow(button1)
	row2 := tgbotapi.NewInlineKeyboardRow(button2)

	newKeyboard := tgbotapi.NewInlineKeyboardMarkup(row1, row2)

	return newKeyboard
}
