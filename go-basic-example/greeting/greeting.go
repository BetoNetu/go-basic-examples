package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	formats := []string{
		"Hi, %v",
		"Wassup, %v",
		"Sup, %v",
		"F* u, %v",
	}
	return formats[rand.Intn(len(formats))]
}

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)

	// Return one of the message formats selected at random.
	return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	// map[key-type]value-type. Maps is the dict of go

	messages := make(map[string]string)
	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved message with
		// the name.
		messages[name] = message
	}
	return messages, nil
}
