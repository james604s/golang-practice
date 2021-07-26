package main

import "fmt"

func paintNeeded(width float64, height float64) (float64, error){
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	}
	area := width * height
	return area / 10.0, nil
}

//建立值以代表錯誤的能力
//從paintNeeded回傳額外值的能力
func main {
	var amount, total float64
	amount = paintNeeded(4.2, 3.0)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%0.2f liters needed\n", amount)
	}
	// total += amount
	// amount = paintNeeded(5.2, 3.5)
	// fmt.Printf("%0.2f liters needed\n", amount)
	// total += amount
	// fmt.Printf("Total: %0.2f liters\n", total)
}