package bot

type Question struct {
	Text    string
	Answer1 string
	Answer2 string
}

var answers = make(map[int64]map[int]string) // мапа для ответов, содержит сам ответ и индекс вопроса к нему
var currentQuestion = make(map[int64]int)    // мапа содержит номер юзера и мапу ответа

func NewQuestions() []Question {
	var questions = []Question{
		// сюда добавим вопросы
	}
	questions = append(questions, Question{
		Text:    "где ты сейчас находишься?",
		Answer1: "я сейчас в центре",
		Answer2: "я сейчас не в центре",
	})

	questions = append(questions, Question{
		Text:    "готов(а) немного проехать ради хорошей кофейни?",
		Answer1: "да, готов(а) — расстояние не проблема",
		Answer2: "нет, хочу что-нибудь поближе",
	})

	questions = append(questions, Question{
		Text:    "что ты хочешь делать в кофейне?",
		Answer1: "спокойно отдохнуть и пообщаться",
		Answer2: "поработать / поучиться",
	})

	questions = append(questions, Question{
		Text:    "что ты скорее хочешь выпить?",
		Answer1: "кофе",
		Answer2: "не кофе",
	})

	questions = append(questions, Question{
		Text:    "что для тебя важнее в кофейне?",
		Answer1: "главное — хороший кофе и напитки",
		Answer2: "хочу ещё вкусно поесть",
	})

	questions = append(questions, Question{
		Text:    "какая атмосфера тебе сегодня ближе?",
		Answer1: "спокойная и сдержанная",
		Answer2: "яркая и атмосферная, чтобы было красиво и можно было сделать фото",
	})

	questions = append(questions, Question{
		Text:    "ты идёшь в кофейню...",
		Answer1: "один / одна",
		Answer2: "с кем-то",
	})

	questions = append(questions, Question{
		Text:    "с тобой будут дети?",
		Answer1: "да",
		Answer2: "нет",
	})

	questions = append(questions, Question{
		Text:    "хочешь ли ты после кофе погулять?",
		Answer1: "да, хочу кофейню в красивом месте",
		Answer2: "нет, хочу просто посидеть в кофейне",
	})
	return questions
}
