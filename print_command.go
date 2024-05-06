package main

type PrintCommand struct {
}

func (pc *PrintCommand) execute() {
	GetApplication().PrintCurrentValue()
}
