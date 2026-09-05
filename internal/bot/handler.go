package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handler struct {
	bot *tgbotapi.BotAPI
}

func NewHandler(bot *tgbotapi.BotAPI) *Handler {
	return &Handler{bot: bot}
}

func (h *Handler) HandlerUpdate(update *tgbotapi.Update) {
	msg := update.Message

	if msg != nil && msg.IsCommand() {
		cmd := msg.Command()
		userID := msg.From.ID
		chatID := msg.Chat.ID

		switch cmd {
		case "start":
			questions := NewQuestions()
			question := questions[0]

			currentQuestion[userID] = 0
			answers[userID] = make(map[int]string)
			delete(recommendedShops, userID)
			delete(currentShop, userID)

			h.sendText(chatID, HandlerCommands(cmd))

			keyboard := NewQuestionKeyboard(question)
			message := tgbotapi.NewMessage(chatID, question.Text)
			message.ReplyMarkup = keyboard

			_, err := h.bot.Send(message)
			if err != nil {
				log.Println(err)
			}
		default:
			h.sendText(chatID, HandlerCommands(cmd))
		}
		return
	}

	callback := update.CallbackQuery
	if callback == nil {
		return
	}

	userID := callback.From.ID
	if answers[userID] == nil {
		return
	}

	_, err := h.bot.Request(tgbotapi.NewCallback(callback.ID, ""))
	if err != nil {
		log.Println(err)
	}

	if callback.Data == "next" || callback.Data == "back" {
		h.handleShopCallback(callback)
		return
	}

	questions := NewQuestions()
	questionIndex := currentQuestion[userID]
	if questionIndex >= len(questions) {
		return
	}

	answers[userID][questionIndex] = callback.Data
	currentQuestion[userID]++

	nextIndex := currentQuestion[userID]
	chatID := callback.Message.Chat.ID

	if nextIndex >= len(questions) {
		h.sendRecommendations(chatID, userID)
		return
	}

	question := questions[nextIndex]
	keyboard := NewQuestionKeyboard(question)
	message := tgbotapi.NewMessage(chatID, question.Text)
	message.ReplyMarkup = keyboard

	_, err = h.bot.Send(message)
	if err != nil {
		log.Println(err)
	}
}

func (h *Handler) handleShopCallback(callback *tgbotapi.CallbackQuery) {
	userID := callback.From.ID
	shops := recommendedShops[userID]
	if len(shops) == 0 {
		return
	}

	index := currentShop[userID]
	if callback.Data == "next" && index < len(shops)-1 {
		currentShop[userID]++
	}
	if callback.Data == "back" && index > 0 {
		currentShop[userID]--
	}
	if currentShop[userID] == index {
		return
	}

	keyboard := NewCoffeeShopKeyboard()
	edit := tgbotapi.NewEditMessageText(
		callback.Message.Chat.ID,
		callback.Message.MessageID,
		shopCardText(userID),
	)
	edit.ReplyMarkup = &keyboard

	_, err := h.bot.Send(edit)
	if err != nil {
		log.Println(err)
	}
}

func (h *Handler) sendRecommendations(chatID int64, userID int64) {
	shops := loadRecommendations(userID)
	recommendedShops[userID] = shops
	currentShop[userID] = 0

	if len(shops) == 0 {
		h.sendText(chatID, "не получилось подобрать кофейню, попробуй /start ещё раз")
		return
	}

	h.sendText(chatID, "спасибо! вот кофейни, которые тебе подойдут:")

	keyboard := NewCoffeeShopKeyboard()
	message := tgbotapi.NewMessage(chatID, shopCardText(userID))
	message.ReplyMarkup = keyboard

	_, err := h.bot.Send(message)
	if err != nil {
		log.Println(err)
	}
}

func (h *Handler) sendText(chatID int64, text string) {
	_, err := h.bot.Send(tgbotapi.NewMessage(chatID, text))
	if err != nil {
		log.Println(err)
	}
}
