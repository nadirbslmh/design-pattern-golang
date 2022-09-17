package main

const (
	TYPE_NORMAL = "normal"
	TYPE_IGLOO  = "igloo"
)

type IBuilder interface {
	setWindowType()
	setDoorType()
	getHouse() House
}

func getBuilder(builderType string) IBuilder {
	switch builderType {
	case TYPE_NORMAL:
		return newNormalBuilder()
	case TYPE_IGLOO:
		return newIglooBuilder()
	default:
		return nil
	}
}
