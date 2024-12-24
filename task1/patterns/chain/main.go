package main

import "fmt"

type RequestHandler interface {
	setNext(handler RequestHandler)
	doSomeLogic()
}

type FirstHandler struct {
	next RequestHandler
}

func (h *FirstHandler) setNext(handler RequestHandler) {
	h.next = handler
}

func (h *FirstHandler) doSomeLogic() {
	fmt.Println("Do some logic in first handler")
	h.next.doSomeLogic()
}

type SecondHandler struct {
	next RequestHandler
}

func (h *SecondHandler) setNext(handler RequestHandler) {
	h.next = handler
}

func (h *SecondHandler) doSomeLogic() {
	fmt.Println("Do some logic in second handler")
	h.next.doSomeLogic()
}

func (h *ThirdHandler) setNext(handler RequestHandler) {
	h.next = handler
}

func (h *ThirdHandler) doSomeLogic() {
	fmt.Println("Do some logic in third handler")
}

type ThirdHandler struct {
	next RequestHandler
}

func main() {
	firstHandler := &FirstHandler{}
	secondHandler := &SecondHandler{}
	thirdHandler := &ThirdHandler{}

	firstHandler.setNext(secondHandler)
	secondHandler.setNext(thirdHandler)

	firstHandler.doSomeLogic()
}
