package main

type compiler struct {
	Code  string
	index int
}

func NewCompiler(code string) compiler {
	return compiler{Code: code, index: 0}
}

func (c *compiler) Compile() CommandBlock {
	var commands []Command

	for c.index < len(c.Code) {
		command := c.toCommand(c.Code[c.index])

		switch command.(type) {
		case *EndBlockCommand:
			return CommandBlock{commands}
		}

		commands = append(commands, command)

		c.index++
	}

	return CommandBlock{commands}
}

func (c *compiler) toCommand(r byte) Command {
	switch r {
	case '>':
		return &IncrementPointerCommand{}
	case '<':
		return &DecrementPointerCommand{}
	case '+':
		return &IncrementValueCommand{}
	case '-':
		return &DecrementValueCommand{}
	case '.':
		return &PrintCommand{}
	case '[':
		c.index++
		return &LoopCommand{block: c.Compile()}
	case ']':
		return &EndBlockCommand{}

	default:
		return &NilCommand{}
	}
}
