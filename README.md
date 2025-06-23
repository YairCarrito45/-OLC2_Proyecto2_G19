# OLC2_Proyecto2_202204578

## Compilador VLangCherry → ARM64

Este proyecto corresponde al desarrollo de un **compilador funcional** para el lenguaje de programación **VLangCherry**, como parte del curso **Organización de Lenguajes y Compiladores 2** de la **Facultad de Ingeniería** de la **Universidad de San Carlos de Guatemala**.

A diferencia del Proyecto 1 (intérprete), este compilador traduce código fuente escrito en VLangCherry a **código ensamblador ARM64**, el cual puede ejecutarse en entornos virtualizados como **QEMU** o físicos como **Raspberry Pi**.

---

## 🧩 Características principales

- ✅ **Analizador léxico y sintáctico** con ANTLR
- ✅ **Análisis semántico** completo con verificación de tipos y ámbitos
- ✅ Generación de **AST (Árbol de Sintaxis Abstracta)**
- ✅ **Generación de código ARM64** en archivos `.s`
- ✅ Reportes de **errores** y **tabla de símbolos**
- ✅ **Interfaz gráfica** para escribir, editar y compilar archivos `.vch`
- ✅ Consola integrada para mostrar la salida y mensajes del compilador
- ✅ Soporte para estructuras de control, funciones, slices y expresiones

---

## 📚 Tecnologías utilizadas

- [Go](https://golang.org/) – Backend del compilador
- [ANTLR](https://www.antlr.org/) – Generador léxico/sintáctico
- [Fyne](https://fyne.io/) – (o alternativa web) Interfaz gráfica
- [QEMU](https://www.qemu.org/) – Emulador ARM64
- [ARM64 Assembly](https://github.com/ARM-software/arm-architecture) – Instrucciones de bajo nivel

---

## 📦 Estructura del proyecto

```bash
├── antlr/                 # Gramática ANTLR (.g4)
├── ast/                   # Construcción del AST
├── assembler/             # Generación de código ARM64
├── gui/                   # Interfaz gráfica (Fyne o alternativa)
├── main.go                # Punto de entrada
├── parser/                # Código generado por ANTLR
├── reports/               # Reportes generados (errores, símbolos, AST)
└── tests/                 # Casos de prueba en .vch
