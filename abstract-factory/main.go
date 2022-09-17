package main

import "fmt"

func main() {
	// motorcycle ducati
	ducatiFactory := GetMotorcycleFactory(BRAND_DUCATI)

	ducatibike := ducatiFactory.makeMotorcycle()

	fmt.Println(ducatibike.getMake(), ducatibike.getName())
}
