package main
import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	fmt.Println(now)
	var year int = now.Year()
	fmt.Println(year)
}