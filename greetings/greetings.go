package greetings

import (
	"errors"
	"fmt"
	"math/rand"

	"rsc.io/quote"
)

// return a "greeting" based on the named person
func Greeting(name string) (string, error) {
	if name == "" {
		return quote.Hello(), errors.New("empty name")
	}

	message := fmt.Sprintf(randomGreeting(), name)
	return message, nil
}

func Greetings(names []string) map[string]string {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Greeting(name)
		if err == nil {
			messages[name] = message
		}
	}

	return messages
}

// return a random format of "greeting"
func randomGreeting() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v",
		"Hello %v, Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
