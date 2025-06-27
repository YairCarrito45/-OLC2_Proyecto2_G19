package arm

type SymbolTable struct {
	vars map[string]Symbol
}

func NewSymbolTable() *SymbolTable {
	return &SymbolTable{vars: make(map[string]Symbol)}
}

func (s *SymbolTable) Set(id string, sym Symbol) {
	s.vars[id] = sym
}

func (s *SymbolTable) Get(id string) (Symbol, bool) {
	val, ok := s.vars[id]
	return val, ok
}
