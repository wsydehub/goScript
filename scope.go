package goScript

type ScopeType int8

const (
	CommonScopeType ScopeType = 0
	IfScopeType     ScopeType = 1
	ForScopeType    ScopeType = 2
	FuncScopeType   ScopeType = 3
)

type Scope struct {
	Type        ScopeType
	LocalVarMap map[string]*Variable
}

func NewScope(t ScopeType) *Scope {
	return &Scope{
		Type:        t,
		LocalVarMap: map[string]*Variable{},
	}
}

func (t *Scope) AddVar(v *Variable) {
	t.LocalVarMap[v.Name] = v
}

func (t *Scope) GetVar(name string) *Variable {
	return t.LocalVarMap[name]
}
