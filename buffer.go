package main

const MemorySize = 1 << 16

type Buffer struct {
	Data    [MemorySize]uint8
	Pointer uint64
}
