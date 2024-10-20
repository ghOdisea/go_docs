package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("name is required for greeting")
	}
    message := fmt.Sprintf(randomFormat(), name)
//     message := fmt.Sprint(randomFormat()) // For testing purposes
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
		"Hello, %v!",
		"What's up, %v?",
		"Hello again, %v.",
		"You again, %v?",
		"Hi there, %v.",
		"Good day, %v.",
		"Hello, %v. How are you?",
		"Hi, %v. What's up?",
		"Good to see you, %v.",
	}

	return formats[rand.Intn(len(formats))]
}
