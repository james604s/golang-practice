package main

import {
	"fmt"
	"reflect"
}

func main() {
	var myInt int
	fmt.Println(reflect.TypeOf(&myInt)) //取得myInt的指標並且印出此指標型別。
	var myFloat float64
	fmt.Println(reflect.TypeOf(&myFloat)) ////取得myFloat的指標並且印出此指標型別。
	var myBool bool
	fmt.Println(reflect.TypeOf(&myBool)) ////取得myBool的指標並且印出此指標型別。
}

