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
# Suma: x1 + x2
ADD x3, x1, x2
# Literal entero: 1
MOV x4, #1
# Suma: x3 + x4
ADD x1, x3, x4
# Print statement
MOV X0, x1
BL print_integer
# Literal entero: 10
MOV x2, #10
# Literal entero: 5
MOV x3, #5
# Resta: x2 - x3
SUB x4, x2, x3
# Print statement
MOV X0, x4
BL print_integer
# Literal entero: 10
MOV x1, #10
# Literal entero: 2
MOV x2, #2
# Multiplicación: x1 * x2
MUL x3, x1, x2
# Print statement
MOV X0, x3
BL print_integer
# Literal entero: 20
MOV x4, #20
# Literal entero: 4
MOV x1, #4
# División: x4 / x1
UDIV x2, x4, x1
# Print statement
MOV X0, x2
BL print_integer
# Literal entero: 13
MOV x3, #13
# Literal entero: 5
MOV x4, #5
# Módulo: x3 % x4
UDIV x25, x3, x4
MSUB x1, x25, x4, x3
# Print statement
MOV X0, x1
BL print_integer
RET

//--------------------------------------------------------------
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
    .ascii "-"