package greetings

import (
	"fmt"
)

func Hello(name string)string {
	message := fmt.Sprintf("Hi %v, Hope you are doing well !",name);
	return message
}