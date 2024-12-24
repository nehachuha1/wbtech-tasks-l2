package main

import "fmt"

type Command interface {
	Execute()
}

type MainController struct {
	command Command
}

func (c *MainController) setCommand(command Command) {
	c.command = command
}

func (c *MainController) executeCommand() {
	c.command.Execute()
}

type ProcessorOne struct {
}

func (p ProcessorOne) Execute() {
	fmt.Println("Doing some logic by first executer")
}

type ProcessorTwo struct {
}

func (p ProcessorTwo) Execute() {
	fmt.Println("Doing some logic by second executer")
}

func main() {
	mainController := MainController{}
	executerOne := ProcessorOne{}
	executerTwo := ProcessorTwo{}
	mainController.setCommand(executerOne)
	mainController.executeCommand()

	mainController.setCommand(executerTwo)
	mainController.executeCommand()
}
