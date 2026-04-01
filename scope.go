package goScript

// ScopeType marks the lifecycle stage of a scope.
type ScopeType int8

const (
	CommonScopeType ScopeType = 0
	IfScopeType     ScopeType = 1
	ForScopeType    ScopeType = 2
	FuncScopeType   ScopeType = 3
)

// Scope stores variables for a single lexical level.
type Scope struct {
	Type        ScopeType
	LocalVarMap map[string]*Variable
}

// NewScope creates an empty scope for a given type.
func NewScope(t ScopeType) *Scope {
	return &Scope{
		Type:        t,
		LocalVarMap: map[string]*Variable{},
	}
}

// AddVar binds a variable into the current scope.
func (t *Scope) AddVar(v *Variable) {
	t.LocalVarMap[v.Name] = v
}

// GetVar resolves a variable from the current scope only.
func (t *Scope) GetVar(name string) *Variable {
	return t.LocalVarMap[name]
}
