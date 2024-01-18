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

// func main(){
// 	log.SetPrefix("Greetings : ")
// 	log.SetFlags(0)
// 	message,err := greetings.Hello("Hritesh")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(message)

// }

// Multiple Arguments

func main(){
	log.SetPrefix("Greetings : ")
	log.SetFlags(0)
	names := []string{"Hritesh","Jake","Alex"}
	message,err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	for name,msg := range message {
		fmt.Println(name,":",msg)
	}

}