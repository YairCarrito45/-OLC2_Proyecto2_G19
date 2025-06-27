.global _start
.section .text
_start:
    BL main
    mov x8, #93
    svc #0

main:
# Literal entero: 2
MOV x1, #2
# Literal entero: 2
MOV x2, #2
# Suma de enteros
ADD x3, x1, x2
# Print entero
MOV X0, x3
BL print_integer
# Salto de línea después de println
ADR x1, .str_33
MOV X0, x1
BL print_string
# Literal entero: 2
MOV x4, #2
# Carga decimal 2.500000 desde memoria
ADR x1, .str_34
LDR d0, [x1]
# Convierte int a float64
SCVTF d0, x4
# Suma/Resta int y float
FADD d1, d0, d0
# Print float64
FMOV D0, d1
BL print_float
# Salto de línea después de println
ADR x1, .str_35
MOV X0, x1
BL print_string
# Carga decimal 2.500000 desde memoria
ADR x1, .str_36
LDR d0, [x1]
# Literal entero: 2
MOV x1, #2
# Convierte int a float64
SCVTF d2, x1
# Suma/Resta float y int
FADD d3, d0, d2
# Print float64
FMOV D0, d3
BL print_float
# Salto de línea después de println
ADR x1, .str_37
MOV X0, x1
BL print_string
# Carga decimal 2.500000 desde memoria
ADR x1, .str_38
LDR d0, [x1]
# Carga decimal 2.500000 desde memoria
ADR x1, .str_39
LDR d0, [x1]
# Suma/Resta de float64
FADD d4, d0, d0
# Print float64
FMOV D0, d4
BL print_float
# Salto de línea después de println
ADR x1, .str_40
MOV X0, x1
BL print_string
# Literal entero: 3
MOV x2, #3
# Literal entero: 1
MOV x3, #1
# Resta de enteros
SUB x4, x2, x3
# Print entero
MOV X0, x4
BL print_integer
# Salto de línea después de println
ADR x1, .str_41
MOV X0, x1
BL print_string
# Carga decimal 3.500000 desde memoria
ADR x1, .str_42
LDR d0, [x1]
# Literal entero: 1
MOV x1, #1
# Convierte int a float64
SCVTF d5, x1
# Suma/Resta float y int
FSUB d6, d0, d5
# Print float64
FMOV D0, d6
BL print_float
# Salto de línea después de println
ADR x1, .str_43
MOV X0, x1
BL print_string
# Literal entero: 3
MOV x2, #3
# Carga decimal 1.500000 desde memoria
ADR x1, .str_44
LDR d0, [x1]
# Convierte int a float64
SCVTF d7, x2
# Suma/Resta int y float
FSUB d0, d7, d0
# Print float64
FMOV D0, d0
BL print_float
# Salto de línea después de println
ADR x1, .str_45
MOV X0, x1
BL print_string
# Carga decimal 3.500000 desde memoria
ADR x1, .str_46
LDR d0, [x1]
# Carga decimal 1.500000 desde memoria
ADR x1, .str_47
LDR d0, [x1]
# Suma/Resta de float64
FSUB d1, d0, d0
# Print float64
FMOV D0, d1
BL print_float
# Salto de línea después de println
ADR x1, .str_48
MOV X0, x1
BL print_string
# Literal entero: 2
MOV x3, #2
# Literal entero: 2
MOV x4, #2
# Multiplicación de enteros
MUL x1, x3, x4
# Print entero
MOV X0, x1
BL print_integer
# Salto de línea después de println
ADR x1, .str_49
MOV X0, x1
BL print_string
# Literal entero: 2
MOV x2, #2
# Carga decimal 2.000000 desde memoria
ADR x1, .str_50
LDR d0, [x1]
# Convierte int a float64
SCVTF d2, x2
# Multiplicación int * float
FMUL d3, d2, d0
# Print float64
FMOV D0, d3
BL print_float
# Salto de línea después de println
ADR x1, .str_51
MOV X0, x1
BL print_string
# Carga decimal 2.000000 desde memoria
ADR x1, .str_52
LDR d0, [x1]
# Literal entero: 2
MOV x3, #2
# Convierte int a float64
SCVTF d4, x3
# Multiplicación float * int
FMUL d5, d0, d4
# Print float64
FMOV D0, d5
BL print_float
# Salto de línea después de println
ADR x1, .str_53
MOV X0, x1
BL print_string
# Carga decimal 2.000000 desde memoria
ADR x1, .str_54
LDR d0, [x1]
# Carga decimal 2.000000 desde memoria
ADR x1, .str_55
LDR d0, [x1]
# Multiplicación de flotantes
FMUL d6, d0, d0
# Print float64
FMOV D0, d6
BL print_float
# Salto de línea después de println
ADR x1, .str_56
MOV X0, x1
BL print_string
# Literal entero: 6
MOV x4, #6
# Literal entero: 2
MOV x1, #2
# División de enteros
UDIV x2, x4, x1
# Print entero
MOV X0, x2
BL print_integer
# Salto de línea después de println
ADR x1, .str_57
MOV X0, x1
BL print_string
# Carga decimal 6.000000 desde memoria
ADR x1, .str_58
LDR d0, [x1]
# Literal entero: 2
MOV x3, #2
# Convierte int a float64
SCVTF d7, x3
# División float / int
FDIV d0, d0, d7
# Print float64
FMOV D0, d0
BL print_float
# Salto de línea después de println
ADR x1, .str_59
MOV X0, x1
BL print_string
# Literal entero: 6
MOV x4, #6
# Carga decimal 2.000000 desde memoria
ADR x1, .str_60
LDR d0, [x1]
# Convierte int a float64
SCVTF d1, x4
# División int / float
FDIV d2, d1, d0
# Print float64
FMOV D0, d2
BL print_float
# Salto de línea después de println
ADR x1, .str_61
MOV X0, x1
BL print_string
# Carga decimal 6.000000 desde memoria
ADR x1, .str_62
LDR d0, [x1]
# Carga decimal 2.000000 desde memoria
ADR x1, .str_63
LDR d0, [x1]
# División de flotantes
FDIV d3, d0, d0
# Print float64
FMOV D0, d3
BL print_float
# Salto de línea después de println
ADR x1, .str_64
MOV X0, x1
BL print_string
# Literal entero: 10
MOV x1, #10
# Literal entero: 3
MOV x2, #3
# Módulo de enteros
UDIV x4, x1, x2
MSUB x3, x4, x2, x1
# Print entero
MOV X0, x3
BL print_integer
# Salto de línea después de println
ADR x1, .str_65
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

    // Copia original
    fmov d1, d0

    // Parte entera (redondeada hacia cero)
    frintz d2, d0
    fcvtzs x0, d2  // Convierte double a entero con redondeo hacia cero
    bl print_integer

    // Imprime punto decimal
    mov x0, #1
    ldr x1, =dot_char
    mov x2, #1
    mov w8, #64
    svc #0

    // Parte fraccionaria = d1 - d2
    fsub d3, d1, d2

    // Multiplica por 1000.0 para obtener 3 decimales
    ldr x19, =float1000
    ldr d4, [x19]
    fmul d5, d3, d4

    // Convierte a entero (sin signo)
    fcvtzu x20, d5
    bl print_3digit_integer

    ldp x21, x22, [sp], #16
    ldp x19, x20, [sp], #16
    ldp x29, x30, [sp], #16
    ret

.p2align 2
//--------------------------------------------------------------
// print_3digit_integer - Always prints 3 digits with leading zeros
// Input: x20 = integer [0-999]
print_3digit_integer:
    sub sp, sp, #16
    mov x1, sp

    // Hundreds
    mov x2, #100
    udiv x3, x20, x2
    msub x20, x3, x2, x20
    add x3, x3, #48
    strb w3, [x1]

    // Tens
    mov x2, #10
    udiv x3, x20, x2
    msub x20, x3, x2, x20
    add x3, x3, #48
    strb w3, [x1, #1]

    // Units
    add x3, x20, #48
    strb w3, [x1, #2]

    // Null-terminator
    mov w3, #0
    strb w3, [x1, #3]      // <- Terminación del string

    // Syscall
    mov x0, #1
    mov x2, #3             // aún puedes usar 3 bytes si lo prefieres
    mov w8, #64
    svc #0

    add sp, sp, #16
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
.str_40:
	.asciz "\n"
.str_48:
	.asciz "\n"
.str_56:
	.asciz "\n"
.str_59:
	.asciz "\n"
.str_65:
	.asciz "\n"
.str_35:
	.asciz "\n"
.str_51:
	.asciz "\n"
.str_61:
	.asciz "\n"
.str_64:
	.asciz "\n"
.str_33:
	.asciz "\n"
.str_37:
	.asciz "\n"
.str_41:
	.asciz "\n"
.str_43:
	.asciz "\n"
.str_45:
	.asciz "\n"
.str_53:
	.asciz "\n"
.str_49:
	.asciz "\n"
.str_57:
	.asciz "\n"
.p2align 3
.p2align 3
.str_34:
    .double 2.500000
.p2align 3
.str_36:
    .double 2.500000
.p2align 3
.str_38:
    .double 2.500000
.p2align 3
.str_39:
    .double 2.500000
.p2align 3
.str_42:
    .double 3.500000
.p2align 3
.str_44:
    .double 1.500000
.p2align 3
.str_46:
    .double 3.500000
.p2align 3
.str_47:
    .double 1.500000
.p2align 3
.str_50:
    .double 2.000000
.p2align 3
.str_52:
    .double 2.000000
.p2align 3
.str_54:
    .double 2.000000
.p2align 3
.str_55:
    .double 2.000000
.p2align 3
.str_58:
    .double 6.000000
.p2align 3
.str_60:
    .double 2.000000
.p2align 3
.str_62:
    .double 6.000000
.p2align 3
.str_63:
    .double 2.000000
.p2align 2
dot_char:
	.asciz "."
.p2align 3
float1000:
	.double 1000.0
