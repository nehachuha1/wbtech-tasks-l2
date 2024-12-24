package main

import "fmt"

type State interface {
	DoAction()
}

type Computer struct {
	powered      State
	ready        State
	currentState State
}

func (c *Computer) NewComputer() {
	c.powered = &PoweredComputer{computer: c}
	c.ready = &ReadyComputer{computer: c}
	c.currentState = c.powered
}

type PoweredComputer struct {
	computer *Computer
}

func (c *PoweredComputer) DoAction() {
	fmt.Println("from powered to ready")
	c.computer.currentState = c.computer.ready
}

type ReadyComputer struct {
	computer *Computer
}

func (c *ReadyComputer) DoAction() {
	fmt.Println("from ready to powered")
	c.computer.currentState = c.computer.powered
}

func main() {
	newComputer := &Computer{}
	newComputer.NewComputer()

	fmt.Printf("%#v\n", newComputer)
	newComputer.currentState.DoAction()
	fmt.Printf("%#v\n", newComputer)
	newComputer.currentState.DoAction()
}
