package value

import "fmt"

type SliceValue struct {
	ElementType string
	Elements    []IVOR
}

func (s *SliceValue) Value() interface{} {
	return s.Elements
}

func (s *SliceValue) Type() string {
	return "slice_" + s.ElementType
}

func (s *SliceValue) Copy() IVOR {
	copied := make([]IVOR, len(s.Elements))
	for i, elem := range s.Elements {
		copied[i] = elem.Copy()
	}
	return &SliceValue{
		ElementType: s.ElementType,
		Elements:    copied,
	}
}

// Para agregar elementos (simulando append)
func (s *SliceValue) Append(val IVOR) error {
	if val.Type() != s.ElementType {
		return fmt.Errorf("tipo incompatible: se esperaba %s pero se recibió %s", s.ElementType, val.Type())
	}
	s.Elements = append(s.Elements, val)
	return nil
}
