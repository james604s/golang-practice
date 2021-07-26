package main

import (
	"bufio"
	"fmt"
	"os"
)
func main() {
	fmt.Print("Enter a grade:") //使用者輸入分數
	reader := bufio.NewReader(os.Stdin)  //緩衝讀取器, 從鍵盤取字串
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
}

