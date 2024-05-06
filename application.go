package main

import (
	"fmt"
	"sync"
)

var once sync.Once

type application struct {
	buffer        Buffer
	commandBuffer CommandBlock
}

var applicationInstance *application

func GetApplication() *application {
	if applicationInstance == nil {
		once.Do(
			func() {
				applicationInstance = &application{}
			})
	}

	return applicationInstance
}

func (a *application) SetCommandBuffer(cb CommandBlock) {
	a.commandBuffer = cb
}

func (a *application) Run() {
	a.commandBuffer.execute()
}

func (a *application) GetCurrentValue() uint8 {
	return a.buffer.Data[a.buffer.Pointer]
}

func (a *application) GetPointer() uint64 {
	return a.buffer.Pointer
}

func (a *application) IncrementCurrentValue() {
	a.buffer.Data[a.buffer.Pointer]++
}

func (a *application) DecrementCurrentValue() {
	a.buffer.Data[a.buffer.Pointer]--
}

func (a *application) IncrementPointer() {
	a.buffer.Pointer++
}

func (a *application) DecrementPointer() {
	a.buffer.Pointer--
}

func (a *application) PrintCurrentValue() {
	fmt.Printf("%c", a.GetCurrentValue())
}
