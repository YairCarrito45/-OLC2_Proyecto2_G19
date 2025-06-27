.global _start
.section .text
_start:
    BL main
    mov x8, #93
    svc #0

main:
# Literal entero: 0
MOV x1, #0
STR x1, [SP, #0]
# Variable puntos inicializada con valor en x1
# Cadena literal: "=== Archivo de prueba básico ==="
ADR x2, .str_43
# Print cadena
MOV X0, x2
BL print_string
# Salto de línea después de println
ADR x1, .str_44
MOV X0, x1
BL print_string
# Cadena literal: "Validaciones manuales esperadas: 3"
ADR x3, .str_45
# Print cadena
MOV X0, x3
BL print_string
# Salto de línea después de println
ADR x1, .str_46
MOV X0, x1
BL print_string
# Cadena literal: "==== Declaración de variables ===="
ADR x4, .str_47
# Print cadena
MOV X0, x4
BL print_string
# Salto de línea después de println
ADR x1, .str_48
MOV X0, x1
BL print_string
# Literal entero: 0
MOV x1, #0
STR x1, [SP, #8]
# Variable puntosDeclaracion inicializada con valor en x1
# Cadena literal: "Declaración explícita con tipo y valor"
ADR x2, .str_49
# Print cadena
MOV X0, x2
BL print_string
# Salto de línea después de println
ADR x1, .str_50
MOV X0, x1
BL print_string
# Literal entero: 42
MOV x3, #42
STR x3, [SP, #16]
# Variable entero inicializada con valor en x3
# Decimal como cadena literal: 3.14
ADR x4, .str_51
STR x4, [SP, #24]
# Variable decimal inicializada con valor en x4
# Cadena literal: "Hola!"
ADR x1, .str_52
STR x1, [SP, #32]
# Variable texto inicializada con valor en x1
# Booleano literal: true
ADR x2, .str_53
STR x2, [SP, #40]
# Variable booleano inicializada con valor en x2
# Cadena literal: "entero:"
ADR x3, .str_54
# Print cadena
MOV X0, x3
BL print_string
# Accediendo variable entero
LDR x4, [SP, #16]
# Print entero
MOV X0, x4
BL print_integer
# Salto de línea después de println
ADR x1, .str_55
MOV X0, x1
BL print_string
# Cadena literal: "decimal:"
ADR x1, .str_56
# Print cadena
MOV X0, x1
BL print_string
# Accediendo variable decimal
LDR x2, [SP, #24]
# Print cadena
MOV X0, x2
BL print_string
# Salto de línea después de println
ADR x1, .str_57
MOV X0, x1
BL print_string
# Cadena literal: "texto:"
ADR x3, .str_58
# Print cadena
MOV X0, x3
BL print_string
# Accediendo variable texto
LDR x4, [SP, #32]
# Print cadena
MOV X0, x4
BL print_string
# Salto de línea después de println
ADR x1, .str_59
MOV X0, x1
BL print_string
# Cadena literal: "booleano:"
ADR x1, .str_60
# Print cadena
MOV X0, x1
BL print_string
# Accediendo variable booleano
LDR x2, [SP, #40]
# Print cadena
MOV X0, x2
BL print_string
# Salto de línea después de println
ADR x1, .str_61
MOV X0, x1
BL print_string
# Cadena literal: ""
ADR x3, .str_62
# Print cadena
MOV X0, x3
BL print_string
# Salto de línea después de println
ADR x1, .str_63
MOV X0, x1
BL print_string
# Accediendo variable puntos
LDR x4, [SP, #0]
# Literal entero: 1
MOV x1, #1
# Suma: x4 + x1
ADD x2, x4, x1
# Print entero
MOV X0, x2
BL print_integer
# Salto de línea después de println
ADR x1, .str_64
MOV X0, x1
BL print_string
RET

//--------------------------------------------------------------
// print_string - Prints a null-terminated string to stdout
//
// Input:
//   x0 - Address of the null-terminated string
//--------------------------------------------------------------
print_string:
    stp x29, x30, [sp, #-16]!  // Save frame pointer and link register

    mov x1, x0        // Start of string
    mov x2, x0        // Pointer to scan for null terminator

find_length:
    ldrb w3, [x2]
    cmp w3, #0
    beq got_length
    add x2, x2, #1
    b find_length

got_length:
    sub x2, x2, x1    // x2 = length
    mov x8, #64       // syscall write
    mov x0, #1        // stdout
    mov x1, x1        // buffer
    svc #0

    ldp x29, x30, [sp], #16
    ret

//--------------------------------------------------------------
// print_integer - Prints a signed integer to stdout
//
// Input:
//   x0 - The integer value to print
//--------------------------------------------------------------
print_integer:
    stp x29, x30, [sp, #-16]!  // Save frame pointer and link register
    stp x19, x20, [sp, #-16]!
    stp x21, x22, [sp, #-16]!
    stp x23, x24, [sp, #-16]!
    stp x25, x26, [sp, #-16]!
    stp x27, x28, [sp, #-16]!

    mov x19, x0                // Save original number
    cmp x19, #0
    bge positive_number

    // Handle negative number
    mov x20, x0
    mov x0, #1                 // stdout
    ldr x1, =minus_sign

    mov x2, #1                 // length = 1
    mov w8, #64                // syscall: write
    svc #0
    mov x0, x20                // Restore x0

    neg x19, x19               // Make number positive

positive_number:
    sub sp, sp, #32            // Reserve space for buffer
    mov x22, sp                // x22 -> buffer
    mov x23, #0                // digit counter

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

    add x26, x26, #48
    strb w26, [x22, x23]
    add x23, x23, #1

    mov x19, x25
    cbnz x19, convert_loop

    // Reverse digits
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
    mov w24, #10
    strb w24, [x22, x23]
    add x23, x23, #1

    mov x0, #1
    mov x1, x22
    mov x2, x23
    mov w8, #64
    svc #0

    add sp, sp, #32
    ldp x27, x28, [sp], #16
    ldp x25, x26, [sp], #16
    ldp x23, x24, [sp], #16
    ldp x21, x22, [sp], #16
    ldp x19, x20, [sp], #16
    ldp x29, x30, [sp], #16
    ret

minus_sign:
    .ascii "-"

.section .data
.str_47:
	.asciz "==== Declaración de variables ===="
.str_53:
	.asciz "true"
.str_57:
	.asciz "\n"
.str_59:
	.asciz "\n"
.str_60:
	.asciz "booleano:"
.str_62:
	.asciz ""
.str_56:
	.asciz "decimal:"
.str_43:
	.asciz "=== Archivo de prueba básico ==="
.str_44:
	.asciz "\n"
.str_45:
	.asciz "Validaciones manuales esperadas: 3"
.str_48:
	.asciz "\n"
.str_50:
	.asciz "\n"
.str_51:
	.asciz "3.14"
.str_52:
	.asciz "Hola!"
.str_49:
	.asciz "Declaración explícita con tipo y valor"
.str_54:
	.asciz "entero:"
.str_55:
	.asciz "\n"
.str_58:
	.asciz "texto:"
.str_61:
	.asciz "\n"
.str_63:
	.asciz "\n"
.str_64:
	.asciz "\n"
.str_46:
	.asciz "\n"
