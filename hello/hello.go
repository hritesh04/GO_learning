package main

import(
	"fmt"
	"log"
	"example.com/greetings"
)

// func main(){
// 	message := greetings.Hello("Hritesh")
// 	fmt.Println(message)
// }

func main(){
	log.SetPrefix("Greetings : ")
	log.SetFlags(0)
	message,err := greetings.Hello("Hritesh")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

}

