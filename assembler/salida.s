.global _start
.section .text
_start:
    BL main
    mov x8, #93
    svc #0

main:
# Literal entero: 5
MOV x1, #5
# Literal entero: 3
MOV x2, #3
# Comparación relacional entre enteros
CMP x1, x2
BGT rel_true_label_0
MOV x3, #0
B rel_end_label_1
rel_true_label_0:
MOV x3, #1
rel_end_label_1:
# Print entero
MOV X0, x3
BL print_integer
# Salto de línea después de println
ADR x1, .str_0
MOV X0, x1
BL print_string
# Carga decimal 2.500000 desde memoria
ADR x1, .str_1
LDR d0, [x1]
# Carga decimal 1.000000 desde memoria
ADR x1, .str_2
LDR d1, [x1]
# Comparación relacional entre float64
FCMP d0, d1
BLT relf_true_label_2
MOV x4, #0
B relf_end_label_3
relf_true_label_2:
MOV x4, #1
relf_end_label_3:
# Print entero
MOV X0, x4
BL print_integer
# Salto de línea después de println
ADR x1, .str_3
MOV X0, x1
BL print_string
# Carga decimal 3.140000 desde memoria
ADR x1, .str_4
LDR d2, [x1]
# Carga decimal 3.140000 desde memoria
ADR x1, .str_5
LDR d3, [x1]
# Comparación relacional entre float64
FCMP d2, d3
BGE relf_true_label_4
MOV x1, #0
B relf_end_label_5
relf_true_label_4:
MOV x1, #1
relf_end_label_5:
# Print entero
MOV X0, x1
BL print_integer
# Salto de línea después de println
ADR x1, .str_6
MOV X0, x1
BL print_string
# Literal entero: 10
MOV x2, #10
# Literal entero: 20
MOV x3, #20
# Comparación relacional entre enteros
CMP x2, x3
BLE rel_true_label_6
MOV x4, #0
B rel_end_label_7
rel_true_label_6:
MOV x4, #1
rel_end_label_7:
# Print entero
MOV X0, x4
BL print_integer
# Salto de línea después de println
ADR x1, .str_7
MOV X0, x1
BL print_string
# Carga decimal 1.100000 desde memoria
ADR x1, .str_8
LDR d4, [x1]
# Literal entero: 2
MOV x1, #2
# Convirtiendo entero a float64 para comparar
SCVTF d5, x1
# Comparación relacional entre float64
FCMP d4, d5
BGT relf_true_label_8
MOV x2, #0
B relf_end_label_9
relf_true_label_8:
MOV x2, #1
relf_end_label_9:
# Print entero
MOV X0, x2
BL print_integer
# Salto de línea después de println
ADR x1, .str_9
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

.section .data
.str_3:
	.asciz "\n"
.str_6:
	.asciz "\n"
.str_7:
	.asciz "\n"
.str_9:
	.asciz "\n"
.str_0:
	.asciz "\n"
.p2align 3
.p2align 3
.str_1:
    .double 2.500000
.p2align 3
.str_2:
    .double 1.000000
.p2align 3
.str_4:
    .double 3.140000
.p2align 3
.str_5:
    .double 3.140000
.p2align 3
.str_8:
    .double 1.100000
.p2align 2
dot_char:
	.asciz "."
.p2align 3
float1000:
	.double 1000.0
