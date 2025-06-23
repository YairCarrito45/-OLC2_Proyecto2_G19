package value

import (
	"fmt"
	"strings"
)

type StructField struct {
	Name string
	Type string
}

type StructDefinition struct {
	Name   string
	Fields []StructField
}

type StructInstance struct {
	Def    *StructDefinition
	Fields map[string]IVOR
}

func NewStructInstance(def *StructDefinition, fields map[string]IVOR) *StructInstance {
	return &StructInstance{Def: def, Fields: fields}
}

func (s *StructInstance) Type() string  { return s.Def.Name }
func (s *StructInstance) Copy() IVOR    { return s } // ✅ Por referencia
func (s *StructInstance) Value() any    { return s }

// ✅ Esto habilita impresión legible: print(p)
func (s *StructInstance) String() string {
	var sb strings.Builder
	sb.WriteString(s.Def.Name + "{")
	first := true
	for _, field := range s.Def.Fields {
		if !first {
			sb.WriteString(", ")
		}
		first = false
		val := s.Fields[field.Name]
		sb.WriteString(fmt.Sprintf("%s: %v", field.Name, val.Value()))
	}
	sb.WriteString("}")
	return sb.String()
}


