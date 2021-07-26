package main

import (
	"fmt"
)

//make建立slice
func main() {
	var notes []string
	notes = make([]string, 7)

	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"
	fmt.Println(notes[0])
	fmt.Println(notes[1])
	fmt.Println(len(notes))

	letters := []string{"a", "b", "c"} //多行的切片字面量
	for i := 0; i < len(letters); i++ {
		fmt.Println(letters[i])
	}

	for _, letter := range letters {
		fmt.Println(letter)
	}

	//切片運算子
	underlyArray := []string{"a", "b", "c", "d", "e", "f"}
	i, j := 1, 4
	slice2 := underlyArray[i:j]
	fmt.Println(slice2)
	slice3 := underlyArray[i:]
	fmt.Println(slice3)
	slice4 := underlyArray[:j]
	fmt.Println(slice4)

	//append
	appendTest := []int{1, 2, 3, 4, 5}
	fmt.Println(appendTest, len(appendTest))

	appendTest = append(appendTest, 10)
	fmt.Println(appendTest, len(appendTest))

	//底層陣列
}
