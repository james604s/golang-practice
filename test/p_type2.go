package main

import (
	"fmt"
)

func main() {
	var myInt int
	var myIntPointer *int //宣告一個只持有指向int指標的變數
	myIntPointer = &myInt //指派一個變數到這個變數
	fmt.Println(myIntPointer)
	var myFloat float64
	var myFloatPointer *float64
	myFloatPointer = &myFloat
	fmt.Println(myFloatPointer)
	var myBool bool
	var myBoolPointer *bool
	myBoolPointer = &myBool
	fmt.Println(myBoolPointer)
}
