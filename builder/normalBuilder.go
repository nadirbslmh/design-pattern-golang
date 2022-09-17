package main

type NormalBuilder struct {
	windowType string
	doorType   string
}

func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (b *NormalBuilder) setWindowType() {
	b.windowType = "Wooden Window"
}
func (b *NormalBuilder) setDoorType() {
	b.doorType = "Normal Door"
}

func (b *NormalBuilder) getHouse() House {
	return House{
		windowType: b.windowType,
		doorType:   b.doorType,
	}
}
