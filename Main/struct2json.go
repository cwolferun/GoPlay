package main

import (
	"encoding/json"
	"fmt"
)

type twoD struct {
	X float32
	Y float32
}

type thirdD struct {
	twoD
	Z float32			`json:",omitempty"`		//dont show entry if the thing is empty
}

type fourthD struct {
	thirdD
	Time string
}

type Spacetime struct {

	Name string			`json:"ObjectName"`		//change the key name in the output
	fourthD

}

func main(){


var	mycord = Spacetime{
	fourthD:
		fourthD{
			Time:"23:00",
			thirdD:
				thirdD{Z:3.4,
				twoD:
					twoD{
						X:3.4,Y:4.2}}},

		Name:"victor",

	}
	out,_ := json.Marshal(mycord)
	fmt.Println(string(out))


}