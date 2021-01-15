package main

import (
	"fmt"
	"math"
)


type Register uint
const (
	R_R0 Register = iota
	R_R1
	R_R2
	R_R3
	R_R4
	R_R5
	R_R6
	R_R7
	R_PC
	R_COND
	R_COUNT
)

const (
	OP_BR = iota
	OP_ADD
	OP_LD
	OP_ST
	OP_JSR
	OP_AND
	OP_LDR
	OP_STR
	OP_RTI
	OP_NOT
	OP_LDI
	OP_STI
	OP_JMP
	OP_RES
	OP_LEA
	OP_TRAP
)

const (
	FL_POS = 1 << iota
	FL_ZRO = 1 << iota
	FL_NEG = 1 << iota
)

var memory [math.MaxInt16]uint16
var registers [R_COUNT]uint16

func main() {
	var foo uint16 = 0xFFFF
	for foo > 0 {
		fmt.Printf("Foo: %d\n", foo)
		foo = foo >> 1
	}
	registers[9] = 12
}

