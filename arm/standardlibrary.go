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
    cmp x19, #0                // Compare with zero
    bge positive_number        // Branch if greater or equal to zero
    
    // Handle negative number
    mov x0, #1                 // fd = 1 (stdout)
    adr x1, minus_sign         // Address of minus sign
    mov x2, #1                 // Length = 1
    mov w8, #64                // Syscall write
    svc #0
    
    neg x19, x19               // Make number positive
    
positive_number:
    // Prepare buffer for converting result to ASCII
    sub sp, sp, #32            // Reserve space on stack
    mov x22, sp                // x22 points to buffer
    
    // Initialize digit counter
    mov x23, #0                // Digit counter
    
    // Handle special case for zero
    cmp x19, #0
    bne convert_loop
    
    // If number is zero, just write '0'
    mov w24, #48               // ASCII '0'
    strb w24, [x22, x23]       // Store in buffer
    add x23, x23, #1           // Increment counter
    b print_result             // Skip conversion loop
    
convert_loop:
    // Divide the number by 10
    mov x24, #10
    udiv x25, x19, x24         // x25 = x19 / 10 (quotient)
    msub x26, x25, x24, x19    // x26 = x19 - (x25 * 10) (remainder)
    
    // Convert remainder to ASCII and store in buffer
    add x26, x26, #48          // Convert to ASCII ('0' = 48)
    strb w26, [x22, x23]       // Store digit in buffer
    add x23, x23, #1           // Increment digit counter
    
    // Prepare for next iteration
    mov x19, x25               // Quotient becomes the new number
    cbnz x19, convert_loop     // If number is not zero, continue
    
    // Reverse the buffer since digits are in reverse order
    mov x27, #0                // Start index
reverse_loop:
    sub x28, x23, x27          // x28 = length - current index
    sub x28, x28, #1           // x28 = length - current index - 1
    
    cmp x27, x28               // Compare indices
    bge print_result           // If crossed, finish reversing
    
    // Swap characters
    ldrb w24, [x22, x27]       // Load character from start
    ldrb w25, [x22, x28]       // Load character from end
    strb w25, [x22, x27]       // Store end character at start
    strb w24, [x22, x28]       // Store start character at end
    
    add x27, x27, #1           // Increment start index
    b reverse_loop             // Continue reversing
    
print_result:
    // Add newline
    mov w24, #10               // Newline character
    strb w24, [x22, x23]       // Add to end of buffer
    add x23, x23, #1           // Increment counter
    
    // Print the result
    mov x0, #1                 // fd = 1 (stdout)
    mov x1, x22                // Buffer address
    mov x2, x23                // Buffer length
    mov w8, #64                // Syscall write
    svc #0
    
    // Clean up and restore registers
    add sp, sp, #32            // Free buffer space
    ldp x27, x28, [sp], #16    // Restore callee-saved registers
    ldp x25, x26, [sp], #16
    ldp x23, x24, [sp], #16
    ldp x21, x22, [sp], #16
    ldp x19, x20, [sp], #16
    ldp x29, x30, [sp], #16    // Restore frame pointer and link register
    ret                        // Return to caller

minus_sign:
    .ascii "-"               // Minus sign`,
	"for_normal": `//--------------------------------------------------------------
// Macro: for_loop
// Realiza un bucle tipo for en ARM64
//
// Parámetros:
//   %1 -> Registro a usar como índice (ej. x0)
//   %2 -> Valor inicial
//   %3 -> Límite (no inclusivo)
//   %4 -> Paso (positivo o negativo)
//
// Uso típico:
//   for_loop x0, 0, 10, 1
//       bl print_integer
//   end_for_loop x0, 1
//--------------------------------------------------------------

    .macro for_loop reg:req, init:req, limit:req, step:req
    mov     \reg, #\init             // Inicializa el contador
    .if (\step) > 0
1:
    cmp     \reg, #\limit            // Comparar con el límite
    bge     2f                       // Si reg >= límite, salir
    .else
1:
    cmp     \reg, #\limit            // Comparar con el límite
    ble     2f                       // Si reg <= límite, salir
    .endif
    .endm

    .macro end_for_loop reg:req, step:req
    .if (\step) > 0
    add     \reg, \reg, #\step       // reg += step
    .else
    sub     \reg, \reg, #-\step      // reg -= abs(step)
    .endif
    b       1b                       // Volver al inicio del bucle
2:
    .endm`, "while_loop": `//--------------------------------------------------------------
//--------------------------------------------------------------
// Macro: while_loop
// Simula un bucle while(condición) en ARM64
//
// Parámetros:
//   %1 -> Registro a comparar (ej. x0)
//   %2 -> Límite o valor de comparación (ej. 10)
//   %3 -> Operador lógico inverso (ej. bge, beq, etc.)
//         Si la condición se cumple, el bucle TERMINA
//
// Ejemplo:
//   while_loop x0, 10, bge   ; while (x0 < 10)
//       bl print_integer
//       add x0, x0, #1
//   end_while_loop
//--------------------------------------------------------------

    .macro while_loop reg:req, limit:req, branch:req
1:
    cmp     \reg, #\limit
    \branch 2f                 // Si la condición se cumple, salta al final
    .endm

    .macro end_while_loop
    b       1b                // Salta al inicio del bucle
2:
    .endm`,
}