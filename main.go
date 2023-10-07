package main

import "fmt"

type Car interface {
	Build()
}

type AudiCar struct{}

func (s *AudiCar) Build() {
	fmt.Println("Assembly of a audi car.")
}

type BmwCar struct{}

func (s *BmwCar) Build() {
	fmt.Println("Assembly of a BMW car.")
}

type MercedesCar struct{}

func (s *MercedesCar) Build() {
	fmt.Println("Assembly of a Mercedes car.")
}

type SuzukiCar struct{}

func (s *SuzukiCar) Build() {
	fmt.Println("Assembly of a Suzuki car.")
}

type ToyotaCar struct{}

func (s *ToyotaCar) Build() {
	fmt.Println("Assembly of a Toyota car.")
}

type AudiCarFactory struct{}

func (f *AudiCarFactory) ConstructorCar() Car {
	return &AudiCar{}
}

type BmwCarFactory struct{}

func (f *BmwCarFactory) ConstructorCar() Car {
	return &BmwCar{}
}

type MercedesCarFactory struct{}

func (f *MercedesCarFactory) ConstructorCar() Car {
	return &MercedesCar{}
}

type SuzukiCarFactory struct{}

func (f *SuzukiCarFactory) ConstructorCar() Car {
	return &SuzukiCar{}
}

type ToyotaCarFactory struct{}

func (f *ToyotaCarFactory) ConstructorCar() Car {
	return &ToyotaCar{}
}

func main() {
	AudiCarFactory := &AudiCarFactory{}
	AudiCar := AudiCarFactory.ConstructorCar()
	AudiCar.Build()

	BmwCarFactory := &BmwCarFactory{}
	BmwCar := BmwCarFactory.ConstructorCar()
	BmwCar.Build()

	MercedesCarFactory := &MercedesCarFactory{}
	MercedesCar := MercedesCarFactory.ConstructorCar()
	MercedesCar.Build()

	SuzukiCarFactory := &SuzukiCarFactory{}
	SuzukiCar := SuzukiCarFactory.ConstructorCar()
	SuzukiCar.Build()

	ToyotaCarFactory := &ToyotaCarFactory{}
	ToyotaCar := ToyotaCarFactory.ConstructorCar()
	ToyotaCar.Build()
}

//output
//Assembly of a audi car.
//Assembly of a BMW car.
//Assembly of a Mercedes car.
//Assembly of a Suzuki car.
//Assembly of a Toyota car.
