package greetings

import (
	"fmt"
	"errors"
)

func Hello(name string) (string,error) {
	if name == "" {
		return "", errors.New("No Name")
	}
	message := fmt.Sprintf("Hi %v, Hope you are doing well !",name)
	return message,nil
}