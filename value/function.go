// function.go
package value

import p "compiler/parser"

type FunctionParam struct {
	Name            string
	Type            string
	PassByReference bool // ✅ NUEVO CAMPO
}

type StructFunction struct {
    ReceiverName string
    ReceiverType string
    IsMutable    bool // Nuevo: indica si el receptor es mutable
    Params       []FunctionParam
    ReturnType   string // Nuevo: tipo de retorno
    Body         p.IBloqueContext
}

type Function struct {
    Name       string
    ReturnType string
    Params     []FunctionParam
    Body       p.IBloqueContext
}

var GlobalStructFunctions = map[string]map[string]*StructFunction{}
var GlobalFunctions = map[string]*Function{}