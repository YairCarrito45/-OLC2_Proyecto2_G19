.global _start
.section .text
_start:
    BL main
    mov x8, #93
    svc #0

main:
# Literal entero: 1
MOV x1, #1
# Literal entero: 1
MOV x2, #1
# Comparación ==
CMP x1, x2
CSET x3, EQ
# Print entero
MOV X0, x3
BL print_integer
# Salto de línea después de println
ADR x1, .str_0
MOV X0, x1
BL print_string
# Literal entero: 1
MOV x4, #1
# Literal entero: 2
MOV x1, #2
# Comparación !=
CMP x4, x1
CSET x2, NE
# Print entero
MOV X0, x2
BL print_integer
# Salto de línea después de println
ADR x1, .str_1
MOV X0, x1
BL print_string
# Literal entero: 1
MOV x3, #1
# Literal entero: 2
MOV x4, #2
# Comparación ==
CMP x3, x4
CSET x1, EQ
# Print entero
MOV X0, x1
BL print_integer
# Salto de línea después de println
ADR x1, .str_2
MOV X0, x1
BL print_string
# Carga decimal 10.000000 desde memoria
ADR x1, .str_3
LDR d0, [x1]
# Literal entero: 10
MOV x2, #10
# Comparación ==
# ❌ Tipos incompatibles en comparación
# ⚠️ Valor nulo, omitiendo impresión
# Salto de línea después de println
ADR x1, .str_4
MOV X0, x1
BL print_string
# Carga decimal 3.500000 desde memoria
ADR x1, .str_5
LDR d0, [x1]
# Carga decimal 3.500000 desde memoria
ADR x1, .str_6
LDR d0, [x1]
# Comparación !=
FCMP d0, d0
CSET x4, NE
# Print entero
MOV X0, x4
BL print_integer
# Salto de línea después de println
ADR x1, .str_7
MOV X0, x1
BL print_string
# Cadena literal: "hola"
ADR x1, .str_8
STR x1, [SP, #0]
# Variable a inicializada con valor en x1
# Cadena literal: "hola"
ADR x2, .str_9
STR x2, [SP, #8]
# Variable b inicializada con valor en x2
# Cadena literal: "adios"
ADR x3, .str_10
STR x3, [SP, #16]
# Variable c inicializada con valor en x3
# Accediendo variable a
LDR x4, [SP, #0]
# Accediendo variable b
LDR x1, [SP, #8]
# Comparación ==
CMP x4, x1
CSET x2, EQ
# Print entero
MOV X0, x2
BL print_integer
# Salto de línea después de println
ADR x1, .str_11
MOV X0, x1
BL print_string
# Accediendo variable a
LDR x3, [SP, #0]
# Accediendo variable c
LDR x4, [SP, #16]
# Comparación !=
CMP x3, x4
CSET x1, NE
# Print entero
MOV X0, x1
BL print_integer
# Salto de línea después de println
ADR x1, .str_12
MOV X0, x1
BL print_string
# Accediendo variable a
LDR x2, [SP, #0]
# Accediendo variable c
LDR x3, [SP, #16]
# Comparación ==
CMP x2, x3
CSET x4, EQ
# Print entero
MOV X0, x4
BL print_integer
# Salto de línea después de println
ADR x1, .str_13
MOV X0, x1
BL print_string
# Booleano literal: true
ADR x1, .str_14
# Booleano literal: false
ADR x2, .str_15
# Comparación ==
CMP x1, x2
CSET x3, EQ
# Print entero
MOV X0, x3
BL print_integer
# Salto de línea después de println
ADR x1, .str_16
MOV X0, x1
BL print_string
# Booleano literal: true
ADR x4, .str_17
# Booleano literal: false
ADR x1, .str_18
# Comparación !=
CMP x4, x1
CSET x2, NE
# Print entero
MOV X0, x2
BL print_integer
# Salto de línea después de println
ADR x1, .str_19
MOV X0, x1
BL print_string
RET

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

.section .data
.str_7:
	.asciz "\n"
.str_13:
	.asciz "\n"
.str_17:
	.asciz "true"
.str_9:
	.asciz "hola"
.str_11:
	.asciz "\n"
.str_14:
	.asciz "true"
.str_18:
	.asciz "false"
.str_0:
	.asciz "\n"
.str_1:
	.asciz "\n"
.str_8:
	.asciz "hola"
.str_12:
	.asciz "\n"
.str_15:
	.asciz "false"
.str_16:
	.asciz "\n"
.str_2:
	.asciz "\n"
.str_4:
	.asciz "\n"
.str_10:
	.asciz "adios"
.str_19:
	.asciz "\n"
.p2align 3
.p2align 3
.str_3:
    .double 10.000000
.p2align 3
.str_5:
    .double 3.500000
.p2align 3
.str_6:
    .double 3.500000
.p2align 2
dot_char:
	.asciz "."
.p2align 3
float1000:
	.double 1000.0
