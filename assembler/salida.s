.global _start
.section .text
_start:
    BL main
    mov x8, #93
    svc #0

main:
# Carga decimal 1.300000 desde memoria
ADR x1, .str_12
LDR d0, [x1]
# Print float64
FMOV D0, d0
BL print_float
# Salto de línea después de println
ADR x1, .str_13
MOV X0, x1
BL print_string
# Cadena literal: "hola"
ADR x1, .str_14
# Cadena literal: "andrea"
ADR x2, .str_15
ADR x3, .str_16
# Print cadena
MOV X0, x3
BL print_string
# Salto de línea después de println
ADR x1, .str_17
MOV X0, x1
BL print_string
# Carga decimal 1.300000 desde memoria
ADR x1, .str_18
LDR d0, [x1]
# Literal entero: 1
MOV x4, #1
# Convierte int a float64
SCVTF d0, x4
# Suma float + int
FADD d1, d0, d0
# Print float64
FMOV D0, d1
BL print_float
# Salto de línea después de println
ADR x1, .str_19
MOV X0, x1
BL print_string
RET

.p2align 2
//--------------------------------------------------------------
// print_float - Prints a float64 with 3 decimal places
//
// Input:
//   d0 - The float64 value to print
//--------------------------------------------------------------
print_float:
    stp x29, x30, [sp, #-16]!
    stp x19, x20, [sp, #-16]!
    stp x21, x22, [sp, #-16]!

    fmov d1, d0
    fcvtzs x0, d0
    bl print_integer

    mov x0, #1
    ldr x1, =dot_char
    mov x2, #1
    mov w8, #64
    svc #0

    scvtf d2, x0
    fsub d3, d1, d2

    ldr x19, =float1000
    ldr d4, [x19]
    fmul d5, d3, d4

    fcvtzu x0, d5
    bl print_integer

    ldp x21, x22, [sp], #16
    ldp x19, x20, [sp], #16
    ldp x29, x30, [sp], #16
    ret

.p2align 2
//--------------------------------------------------------------
// print_string - Prints a null-terminated string to stdout
//
// Input:
//   x0 - Address of the null-terminated string
//--------------------------------------------------------------
print_string:
    stp x29, x30, [sp, #-16]!

    mov x1, x0
    mov x2, x0

find_length:
    ldrb w3, [x2]
    cmp w3, #0
    beq got_length
    add x2, x2, #1
    b find_length

got_length:
    sub x2, x2, x1
    mov x8, #64
    mov x0, #1
    mov x1, x1
    svc #0

    ldp x29, x30, [sp], #16
    ret

.p2align 2
//--------------------------------------------------------------
// print_integer - Prints a signed integer to stdout
//
// Input:
//   x0 - The integer value to print
//--------------------------------------------------------------
print_integer:
    stp x29, x30, [sp, #-16]!
    stp x19, x20, [sp, #-16]!
    stp x21, x22, [sp, #-16]!
    stp x23, x24, [sp, #-16]!
    stp x25, x26, [sp, #-16]!
    stp x27, x28, [sp, #-16]!

    mov x19, x0
    cmp x19, #0
    bge positive_number

    mov x20, x0
    mov x0, #1
    ldr x1, =minus_sign
    mov x2, #1
    mov w8, #64
    svc #0
    mov x0, x20

    neg x19, x19

positive_number:
    sub sp, sp, #32
    mov x22, sp
    mov x23, #0

    cmp x19, #0
    bne convert_loop

    mov w24, #48
    strb w24, [x22, x23]
    add x23, x23, #1
    b print_result

convert_loop:
    mov x24, #10
    udiv x25, x19, x24
    msub x26, x25, x24, x19
    add x26, x26, #48
    strb w26, [x22, x23]
    add x23, x23, #1
    mov x19, x25
    cbnz x19, convert_loop

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
.str_13:
	.asciz "\n"
.str_14:
	.asciz "hola"
.str_15:
	.asciz "andrea"
.str_16:
	.asciz "holaandrea"
.str_17:
	.asciz "\n"
.str_19:
	.asciz "\n"
.p2align 3
.p2align 3
.str_12:
    .double 1.300000
.p2align 3
.str_18:
    .double 1.300000
.p2align 2
dot_char:
	.asciz "."
.p2align 3
float1000:
	.double 1000.0
