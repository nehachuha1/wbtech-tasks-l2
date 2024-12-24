package main

import "fmt"

type Builder interface {
	setWheelCount()
	setColor()
	setMaxSpeed()
	GetNewCar() Car
}

type Car interface {
	Ride()
	Honk()
}

type FirstTypeOfCar struct {
	WheelCount int
	Color      string
	MaxSpeed   int
}

func (f *FirstTypeOfCar) Ride() {
	fmt.Println("First type of car: I'm riding")
}

func (f *FirstTypeOfCar) Honk() {
	fmt.Println("Honk by first type of car")
}

func (f *FirstTypeOfCar) setWheelCount() {
	f.WheelCount = 4
}

func (f *FirstTypeOfCar) setColor() {
	f.Color = "blue"
}

func (f *FirstTypeOfCar) setMaxSpeed() {
	f.MaxSpeed = 96
}

func (f *FirstTypeOfCar) GetNewCar() Car {
	return &FirstTypeOfCar{
		WheelCount: f.WheelCount,
		Color:      f.Color,
		MaxSpeed:   f.MaxSpeed,
	}
}

type SecondTypeOfCar struct {
	WheelCount int
	Color      string
	MaxSpeed   int
}

func (s *SecondTypeOfCar) setWheelCount() {
	s.WheelCount = 8
}

func (s *SecondTypeOfCar) setColor() {
	s.Color = "Dark green"
}

func (s *SecondTypeOfCar) setMaxSpeed() {
	s.MaxSpeed = 40
}

func (s *SecondTypeOfCar) GetNewCar() Car {
	return &SecondTypeOfCar{
		WheelCount: s.WheelCount,
		Color:      s.Color,
		MaxSpeed:   s.MaxSpeed,
	}
}

func (s *SecondTypeOfCar) Ride() {
	fmt.Println("Second type of car: I'm riding")
}

func (s *SecondTypeOfCar) Honk() {
	fmt.Println("Honk by second type of car")
}

type Factory struct {
	CarBuilder Builder
}

func NewFactory(b Builder) *Factory {
	return &Factory{CarBuilder: b}
}

func (f *Factory) ChangeBuilder(b Builder) {
	f.CarBuilder = b
}

func (f *Factory) MakeNewCar() Car {
	f.CarBuilder.setColor()
	f.CarBuilder.setMaxSpeed()
	f.CarBuilder.setWheelCount()

	return f.CarBuilder.GetNewCar()
}

func main() {
	car1 := FirstTypeOfCar{}
	car2 := SecondTypeOfCar{}

	factory := NewFactory(&car1)
	factory.MakeNewCar()
	fmt.Printf("Car 1: %#v\n", factory.CarBuilder)

	factory.ChangeBuilder(&car2)
	factory.MakeNewCar()
	fmt.Printf("Car 2: %#v\n", factory.CarBuilder)
}
