package bot

import "fmt"

func commands(message string) (string, int, error) {
	var text string
	var stage int
	switch message {
	case "/help":
		text, stage = commandHelp()
		return text, stage, nil
	case "/start":
		text, stage = commandStart()
		return text, stage, nil
	default:
		return "", StageDefault, fmt.Errorf("Command " + message + " not found")
	}
}

func commandHelp() (string, int) {
	message := "Hello world!"
	return message, StageDefault
}

func commandStart() (string, int) {
	return "Введдите логин/пароль от аккаунта Aliexpress через пробел", StageLogin
}
