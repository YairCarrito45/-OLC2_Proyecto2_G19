.section .text
.global _start

_start:
    mov x0, #1           // fd = 1 (stdout)
    adr x1, msg          // x1 = dirección del mensaje
    mov x2, len          // x2 = longitud del mensaje
    mov w8, #64          // syscall write
    svc #0

    mov x0, #0           // código de salida
    mov w8, #93          // syscall exit
    svc #0

.section .data
msg:
    .ascii "Hello, World!\n"
len = . - msg
