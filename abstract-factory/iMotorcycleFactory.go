package main

const (
	BRAND_YAMAHA = "Yamaha"
	BRAND_DUCATI = "Ducati"
)

type IMotorcycleFactory interface {
	makeMotorcycle() IMotorcycle
}

func GetMotorcycleFactory(brand string) IMotorcycleFactory {
	switch brand {
	case BRAND_YAMAHA:
		return &Yamaha{}
	case BRAND_DUCATI:
		return &Ducati{}
	default:
		return nil
	}
}
