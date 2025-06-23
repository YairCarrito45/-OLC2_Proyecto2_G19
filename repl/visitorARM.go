package repl

import (
	parser "compiler/parser"
)

// Visitor personalizado para recorrer el árbol de sintaxis

type ARMVisitor struct {
	parser.BaseVlangVisitor
	Code        []string
	Data        []string
	ScopeTrace  *ScopeTrace
	LabelCount  int
	Console     *Console
}

// Traza de scopes (útil para funciones, if, loops, etc.)
type ScopeTrace struct {
	Scopes []string
}

// Constructor del visitor ARM
func NewARMVisitor(console *Console) *ARMVisitor {
	return &ARMVisitor{
		Code:       []string{},
		Data:       []string{},
		ScopeTrace: &ScopeTrace{},
		LabelCount: 0,
		Console:    console,
	}
}
