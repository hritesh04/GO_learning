package main

import(
	"fmt"
	"log"
	"example.com/greetings"
)

// Import function from diference package

// func main(){
// 	message := greetings.Hello("Hritesh")
// 	fmt.Println(message)
// }

// Error handling

func main(){
	log.SetPrefix("Greetings : ")
	log.SetFlags(0)
	message,err := greetings.Hello("Hritesh")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

}

