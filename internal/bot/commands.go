package bot

func HandlerCommands(cmd string) string {
	switch {
	case cmd == "start":
		return "привет, давай выберем, где попить кофе сегодня!"
	case cmd == "help":
		return "я подберу кофейню под твоё настроение. " +
			"ответь на несколько вопросов — и узнаешь, где провести сегодняшний кофейный день."
	default:
		return "Неизвестная команда"
	}
}
