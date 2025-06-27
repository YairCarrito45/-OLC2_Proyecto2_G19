.global _start
.section .text
_start:
    BL main
    mov x8, #93
    svc #0

main:
# Cadena literal: "Nombre: "
ADR x1, .str_0
# Cadena literal: "jaunito"
ADR x2, .str_1
ADR x3, .str_2
# Print cadena
MOV X0, x3
BL print_string
# Salto de línea después de println
ADR x1, .str_3
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

.section .data
.str_0:
	.asciz "Nombre: "
.str_1:
	.asciz "jaunito"
.str_2:
	.asciz "Nombre: jaunito"
.str_3:
	.asciz "\n"
