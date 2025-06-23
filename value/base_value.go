package value


import "strconv"

// IVOR stands for Internal Value Object Representation
type IVOR interface {
	Value() interface{}
	Type() string
	Copy() IVOR
	
}


const (
	IVOR_INT              = "Int"
	IVOR_FLOAT            = "Float"
	IVOR_STRING           = "String"
	IVOR_BOOL             = "Bool"
	IVOR_CHARACTER        = "Character"
	IVOR_NIL              = "nil"
	IVOR_BUILTIN_FUNCTION = "builtin_function"
	IVOR_FUNCTION         = "function"
	IVOR_VECTOR           = "vector"
	IVOR_OBJECT           = "object"
	IVOR_ANY              = "any"
	IVOR_POINTER          = "pointer"
	IVOR_MATRIX           = "matrix"
	IVOR_SELF             = "self"
	IVOR_UNINITIALIZED    = "uninitialized"
)



func NewInt(val int) IVOR {
	return &IntValue{InternalValue: val}
}

func NewFloat(val float64) IVOR {
	return &FloatValue{InternalValue: val}
}

func NewString(val string) IVOR {
	unquoted, err := strconv.Unquote(val)
	if err != nil {
		unquoted = val
	}
	return &StringValue{InternalValue: unquoted}
}


func NewBool(val bool) IVOR {
	return &BoolValue{InternalValue: val}
}

func NewNil() IVOR {
	return DefaultNilValue
}
