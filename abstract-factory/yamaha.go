package main

type Yamaha struct{}

func (y *Yamaha) makeMotorcycle() IMotorcycle {
	return &YamahaMotorcycle{Motorcycle{
		make: "Yamaha",
		name: "YZF-R1",
	}}
}
