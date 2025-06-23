package value

type CharValue struct {
	InternalValue rune
}

func (c CharValue) Value() interface{}   { return c.InternalValue }
func (c CharValue) Type() string         { return IVOR_CHARACTER }
func (c CharValue) Copy() IVOR           { return &CharValue{c.InternalValue} }

func NewChar(val rune) IVOR {
	return &CharValue{val}
}
