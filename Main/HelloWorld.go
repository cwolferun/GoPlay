package main

import (
	"fmt"		//dealing with formatting

	"os"
)

func main(){

	fmt.Print("Hello")

	//regular for loop where strings and sep are declared at once
	var s, sep string
	for i := 1; i < len(os.Args);i++{
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	//traditional while loop
	var condition bool		//how to declare variables in go
	condition2 := true		//dynamic declaration based on what its set to
	for !condition && condition2{
		condition = true
		condition2 = false
	}

	s2,sep2 := "",""
	for _, arg := range os.Args[1:]{			//returns two values, index and element value
		s2 += sep2 + arg
		sep2 = " "
	}
	fmt.Println(s2)



}