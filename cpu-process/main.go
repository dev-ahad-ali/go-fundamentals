package main

import (
	"errors"
	"fmt"
)

// OpCode represents operation types supported by the simulated ALU.
type OpCode int

const (
	OpADD OpCode = iota // Addition
	OpSUB               // Subtraction
	OpMUL               // Multiplication
	OpDIV               // Division
	OpAND               // Bitwise AND
	OpOR                // Bitwise OR
	OpNOT               // Bitwise NOT
)

// Instruction represents a binary instruction stored in RAM.
type Instruction struct {
	Op      OpCode // Operation code for ALU
	RegSrcA int    // Source Register A index
	RegSrcB int    // Source Register B index (unused for NOT)
	RegDest int    // Destination Register index
}

// CPU represents the core hardware components (Registers + PC).
type CPU struct {
	Registers [8]int64 // Simulated Register Set (R0 - R7)
	PC        int      // Program Counter (points to memory address in RAM)
}

// ExecuteALU simulates the 7 basic Arithmetic Logic Unit operations.
func (cpu *CPU) ExecuteALU(op OpCode, valA, valB int64) (int64, error) {
	switch op {
	case OpADD:
		return valA + valB, nil
	case OpSUB:
		return valA - valB, nil
	case OpMUL:
		return valA * valB, nil
	case OpDIV:
		if valB == 0 {
			return 0, errors.New("ALU runtime error: division by zero")
		}
		return valA / valB, nil
	case OpAND:
		return valA & valB, nil
	case OpOR:
		return valA | valB, nil
	case OpNOT:
		return ^valA, nil
	default:
		return 0, errors.New("ALU runtime error: unrecognized opcode")
	}
}

// RunControlUnit simulates the Control Unit (CU) Fetch-Decode-Execute cycle.
func (cpu *CPU) RunControlUnit(ram []Instruction) error {
	fmt.Println("--- Control Unit: Starting Fetch-Decode-Execute Cycle ---")

	for cpu.PC < len(ram) {
		// Step 1: Fetch instruction pointed to by Program Counter (PC)
		currentPC := cpu.PC
		instruction := ram[currentPC]

		// Advance Program Counter to point to the next instruction in memory
		cpu.PC++

		// Step 2 & 3: Decode instruction and read operands from registers
		valA := cpu.Registers[instruction.RegSrcA]
		valB := cpu.Registers[instruction.RegSrcB]

		// Step 4: Pass operands to ALU for execution
		result, err := cpu.ExecuteALU(instruction.Op, valA, valB)
		if err != nil {
			return fmt.Errorf("error at PC [%d]: %w", currentPC, err)
		}

		// Step 5: Write result back to destination register
		cpu.Registers[instruction.RegDest] = result

		fmt.Printf("PC [%d] Executed Op: %d | R%d (%d) & R%d (%d) -> Dest R%d = %d\n",
			currentPC, instruction.Op, instruction.RegSrcA, valA, instruction.RegSrcB, valB, instruction.RegDest, result)
	}

	fmt.Println("--- Execution Cycle Complete ---")
	return nil
}

func main() {
	// Initialize CPU registers with sample values (R0 = 12, R1 = 4, R2 = 2)
	simulatedCPU := CPU{
		Registers: [8]int64{12, 4, 2, 0, 0, 0, 0, 0},
		PC:        0, // PC starts at address 0
	}

	// Simulated RAM memory array containing binary instructions loaded from disk
	simulatedRAM := []Instruction{
		{Op: OpADD, RegSrcA: 0, RegSrcB: 1, RegDest: 3}, // R3 = R0 + R1  (12 + 4 = 16)
		{Op: OpMUL, RegSrcA: 3, RegSrcB: 2, RegDest: 4}, // R4 = R3 * R2  (16 * 2 = 32)
		{Op: OpSUB, RegSrcA: 4, RegSrcB: 1, RegDest: 5}, // R5 = R4 - R1  (32 - 4 = 28)
		{Op: OpAND, RegSrcA: 0, RegSrcB: 1, RegDest: 6}, // R6 = R0 & R1   (12 & 4 = 4)
	}

	// Control Unit executes program instructions stored in RAM
	err := simulatedCPU.RunControlUnit(simulatedRAM)
	if err != nil {
		fmt.Println("Execution failed:", err)
		return
	}

	// Output final state of CPU registers
	fmt.Println("\nFinal Register State:")
	for index, val := range simulatedCPU.Registers {
		fmt.Printf("Register R%d: %d\n", index, val)
	}
}
