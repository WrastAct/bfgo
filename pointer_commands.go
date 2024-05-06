package main

type GetPointerCommand struct {
}

type IncrementPointerCommand struct {
}

type DecrementPointerCommand struct {
}

func (gp *GetPointerCommand) execute() {
	GetApplication().GetPointer()
}

func (ip *IncrementPointerCommand) execute() {
	GetApplication().IncrementPointer()
}

func (dp *DecrementPointerCommand) execute() {
	GetApplication().DecrementPointer()
}
