package main
import (
	"fmt"
	"strings"
)
//Replacer搜尋字串並替代
//值 now.Year() 方法
func main() {
	broken := "G# r#cks!"
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)
}