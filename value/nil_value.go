package value

type NilValue struct {
}


func IsValidType(t string) bool {
	return t == IVOR_INT || t == IVOR_FLOAT || t == IVOR_STRING ||
		t == IVOR_BOOL || t == IVOR_CHARACTER || isSliceType(t)
}

func isSliceType(t string) bool {
	return len(t) > 6 && t[:6] == "slice_"
}


func (s NilValue) Value() interface{} {
	return nil
}

func (s NilValue) Type() string {
	return IVOR_NIL
}

func (s NilValue) Copy() IVOR {
	return DefaultNilValue
}

var DefaultNilValue = &NilValue{}

type UnInitializedValue struct {
}

func (s UnInitializedValue) Value() interface{} {
	return nil
}

func (s UnInitializedValue) Type() string {
	return IVOR_UNINITIALIZED
}

func (s UnInitializedValue) Copy() IVOR {
	return DefaultUnInitializedValue
}

var DefaultUnInitializedValue = &UnInitializedValue{}
