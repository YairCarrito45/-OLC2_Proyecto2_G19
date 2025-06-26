package arm

import "strings"

// StandardLibrary contiene las funciones de la biblioteca estándar de ARM

type StandardLibrary struct {
	usedFunctions map[string]struct{}
}

func NewStandardLibrary() *StandardLibrary {
	return &StandardLibrary{
		usedFunctions: make(map[string]struct{}),
	}
}

// Use registra una función como utilizada
func (s *StandardLibrary) Use(function string) {
	s.usedFunctions[function] = struct{}{}
}

// GetFunctionDefinitions retorna el código ensamblador de las funciones utilizadas
func (s *StandardLibrary) GetFunctionDefinitions() string {
	var functions []string
	for fn := range s.usedFunctions {
		if def, ok := functionDefinitions[fn]; ok {
			functions = append(functions, def)
		}
	}
	return strings.Join(functions, "\n\n")
}

// AddFunctions -> Agregarian Macros
var functionDefinitions = map[string]string{
"print_integer": `//--------------------------------------------------------------
// print_integer - Prints a signed integer to stdout
//
// Input:
//   x0 - The integer value to print
//--------------------------------------------------------------
print_integer:
    // Save registers
    stp x29, x30, [sp, #-16]!  // Save frame pointer and link register
    stp x19, x20, [sp, #-16]!  // Save callee-saved registers
    stp x21, x22, [sp, #-16]!
    stp x23, x24, [sp, #-16]!
    stp x25, x26, [sp, #-16]!
    stp x27, x28, [sp, #-16]!

    // Check if number is negative
    mov x19, x0                // Save original number
    cmp x19, #0
    bge positive_number

    // Handle negative number
    mov x20, x0                // Backup x0
    mov x0, #1                 // fd = 1 (stdout)
    adr x1, minus_sign         // Address of minus sign
    mov x2, #1                 // Length = 1
    mov w8, #64                // Syscall write
    svc #0
    mov x0, x20                // Restore x0

    neg x19, x19               // Make number positive

positive_number:
    sub sp, sp, #32            // Reserve space for buffer
    mov x22, sp                // x22 -> buffer
    mov x23, #0                // Digit counter

    cmp x19, #0
    bne convert_loop

    mov w24, #48               // ASCII '0'
    strb w24, [x22, x23]
    add x23, x23, #1
    b print_result

convert_loop:
    mov x24, #10
    udiv x25, x19, x24         // x25 = x19 / 10
    msub x26, x25, x24, x19    // x26 = remainder

    add x26, x26, #48          // Convert to ASCII
    strb w26, [x22, x23]
    add x23, x23, #1

    mov x19, x25
    cbnz x19, convert_loop

    // Reverse the buffer
    mov x27, #0
reverse_loop:
    sub x28, x23, x27
    sub x28, x28, #1

    cmp x27, x28
    bge print_result

    ldrb w24, [x22, x27]
    ldrb w25, [x22, x28]
    strb w25, [x22, x27]
    strb w24, [x22, x28]

    add x27, x27, #1
    b reverse_loop

print_result:
    mov w24, #10               // Newline
    strb w24, [x22, x23]
    add x23, x23, #1

    mov x0, #1                 // fd = 1
    mov x1, x22                // Buffer address
    mov x2, x23                // Length
    mov w8, #64                // Syscall write
    svc #0

    // Clean up and return
    add sp, sp, #32
    ldp x27, x28, [sp], #16
    ldp x25, x26, [sp], #16
    ldp x23, x24, [sp], #16
    ldp x21, x22, [sp], #16
    ldp x19, x20, [sp], #16
    ldp x29, x30, [sp], #16
    ret

minus_sign:
    .ascii "-"`,

}