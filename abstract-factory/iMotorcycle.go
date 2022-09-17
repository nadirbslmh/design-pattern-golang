package main

type IMotorcycle interface {
	setMake(make string)
	setName(name string)
	getMake() string
	getName() string
}

type Motorcycle struct {
	make string
	name string
}

func (m *Motorcycle) setMake(make string) {
	m.make = make
}
func (m *Motorcycle) setName(name string) {
	m.name = name
}
func (m *Motorcycle) getMake() string {
	return m.make
}
func (m *Motorcycle) getName() string {
	return m.name
}
