package main

type Command interface {
	execute()
}

type CommandBlock struct {
	commands []Command
}

type EndBlockCommand struct {
}

type NilCommand struct {
}

func (cb *CommandBlock) execute() {
	for _, val := range cb.commands {
		val.execute()
	}
}

func (eb *EndBlockCommand) execute() {
}

func (nc *NilCommand) execute() {
}
