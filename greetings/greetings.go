package greetings

import (
	"fmt"
	"errors"
	"math/rand"
)

// Error Handling

// func Hello(name string) (string,error) {
// 	if name == "" {
// 		return "", errors.New("No Name")
// 	}
// 	message := fmt.Sprintf("Hi %v, Hope you are doing well !",name)
// 	return message,nil
// }

// Returning Random Msg

func Hello(name string) (string,error) {
	if name == "" {
		return name,errors.New("No Name")
	}
	message := fmt.Sprintf(randomMessage(),name)
	return message,nil
}

func randomMessage() string {
	formats := []string{
		"Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}

// taking a list of arguments

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _,name := range names {
		message,err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages,nil
}