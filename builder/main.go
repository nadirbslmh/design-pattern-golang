package main

import "fmt"

func main() {
	normalBuilder := getBuilder(TYPE_NORMAL)

	director := newDirector(normalBuilder)

	house := director.buildHouse()

	fmt.Println(house.doorType, house.windowType)

}
