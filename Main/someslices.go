package main

import "fmt"

func main() {

//	var newarray []int
//	var newarray2 []int32

	intarray := []int{2, 3, 4, 5, 6, 7, 8, 9, 3}
	intslice := intarray[3:8] //56789
	for _, v := range intslice {
//		newarray2 = append(newarray, v) //joins new value to empty array and stores result. Which is [3] after whole loop finishes
		intarray = append(intarray, v)  //append the slice of itself to intarray. I'm satisfied with that.
	}

	fmt.Println(intarray)

}
