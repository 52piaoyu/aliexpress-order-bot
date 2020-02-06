package bot

import "fmt"

func commands(message string) (string, error) {
	switch message {
	case "/help":
		return commandHelp(), nil
	case "/start":
		return commandHelp(), nil
	default:
		return "", fmt.Errorf("Command " + message + " not found")
	}
}

func commandHelp() string {
	message := "Hello world!"
	return message
}
