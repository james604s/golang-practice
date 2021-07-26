package main
import "fmt"
func paintNeeded(width float64, height float64){
	area := width * height
	fmt.Printf("%.2f liters needed\n", area/10.0)
}

//第一個float64是傳入的變數值型別, 第二個是return值型別
func double(number float64) float64{
	return number * 2
}

func main(){
	paintNeeded(4.2, 3.0)
	paintNeeded(5.2, 3.5)
	paintNeeded(5.0, 3.3)
	dozen := double(6.0)
	fmt.Println(dozen)
	fmt.Println(double(4.2))
}
