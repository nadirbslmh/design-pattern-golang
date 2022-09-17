package main

type IglooBuilder struct {
	windowType string
	doorType   string
}

func newIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (b *IglooBuilder) setWindowType() {
	b.windowType = "Snow Window"
}
func (b *IglooBuilder) setDoorType() {
	b.doorType = "Snow Door"
}

func (b *IglooBuilder) getHouse() House {
	return House{
		windowType: b.windowType,
		doorType:   b.doorType,
	}
}
