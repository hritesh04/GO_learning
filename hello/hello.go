package main

import(
	"fmt"
	"example.com/greetings"
)

func main(){
	message := greetings.Hello("Hritesh")
	fmt.Println(message)
}

