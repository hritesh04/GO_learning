package main

import (
	"fmt"
	"unicode/utf8"
	"errors"
)

func main(){
	input := "The future is what we do for it today"
	reversed,reversedErr := Reverse(input)
	reRev,reRevErr := Reverse(reversed)
	fmt.Printf("The original msg : %q ",input)
	fmt.Printf("The reversed msg : %q, err: %v",reversed,reversedErr)
	fmt.Printf("The Re-reversed msg : %q, err: %v",reRev,reRevErr)
}

func Reverse(s string)(string,error) {
	fmt.Printf("Input : %q\n",s)
	//b := []byte(s)
	b:=[]rune(s)
	fmt.Printf("Rune : %q\n",b)
	if !utf8.ValidString(s) {
        return s, errors.New("input is not valid UTF-8")
    }
	for i,j := 0,len(b)-1;i<len(b)/2;i,j = i+1,j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b),nil
}