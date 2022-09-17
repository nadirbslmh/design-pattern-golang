package main

type Ducati struct{}

func (d *Ducati) makeMotorcycle() IMotorcycle {
	return &DucatiMotorcycle{Motorcycle{
		make: "Ducati",
		name: "Paningale V4",
	}}
}
