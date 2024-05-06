package main

type LoopCommand struct {
	block CommandBlock
}

func (lc *LoopCommand) execute() {
	for GetApplication().GetCurrentValue() != 0 {
		lc.block.execute()
	}
}
