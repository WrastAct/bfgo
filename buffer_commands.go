package main

type GetValueCommand struct {
}

type IncrementValueCommand struct {
}

type DecrementValueCommand struct {
}

func (gp *GetValueCommand) execute() {
	GetApplication().GetCurrentValue()
}

func (ip *IncrementValueCommand) execute() {
	GetApplication().IncrementCurrentValue()
}

func (dp *DecrementValueCommand) execute() {
	GetApplication().DecrementCurrentValue()
}
