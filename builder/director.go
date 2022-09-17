package main

type Director struct {
	builder IBuilder
}

func newDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) buildHouse() House {
	d.builder.setWindowType()
	d.builder.setDoorType()

	return d.builder.getHouse()
}
