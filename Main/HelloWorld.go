package main

import (
	"fmt" //dealing with formatting
	"net/http"
	"log"
	"os"
	"strconv"
	"encoding/json"
)

var Myobj = make(map[string]string)

func main() {
	//Myobj["a thing"] = "3"


	for indx,elem := range os.Args[1:]{
		keyname := "Var_" + strconv.Itoa(indx)
		Myobj[keyname]=elem
	}

	/*
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

	fmt.Println(s2)
*/

	http.HandleFunc("/", handler)
	fmt.Println("Up and running")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	obj, err := json.Marshal(Myobj)
	if err != nil {
		fmt.Print(err)
		return
	}else {
		fmt.Fprintln(w, string(obj))	}

}


