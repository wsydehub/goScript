package goScript

// ScopeStack is a LIFO stack for scopes.
type ScopeStack struct {
	stack []*Scope
}

// NewScopeStack creates an empty scope stack.
func NewScopeStack() *ScopeStack {
	return &ScopeStack{stack: []*Scope{}}
}

// Push adds a scope to the top.
func (t *ScopeStack) Push(ele *Scope) {
	t.stack = append(t.stack, ele)
}

// Pop removes and returns the top scope.
func (t *ScopeStack) Pop() *Scope {
	if t.IsEmpty() {
		panic("ScopeStack is empty, can't pop")
	}
	top := t.stack[len(t.stack)-1]
	t.stack = t.stack[:len(t.stack)-1]
	return top
}

// Top returns the top scope without removing it.
func (t *ScopeStack) Top() *Scope {
	if t.IsEmpty() {
		panic("ScopeStack is empty, can't top")
	}
	return t.stack[len(t.stack)-1]
}

// IsEmpty checks whether the stack has elements.
func (t *ScopeStack) IsEmpty() bool {
	return len(t.stack) == 0
}

// FuncStack is a LIFO stack for active functions.
type FuncStack struct {
	stack []*Function
}

// NewFuncStack creates an empty function stack.
func NewFuncStack() *FuncStack {
	return &FuncStack{stack: []*Function{}}
}

// Push adds a function to the top.
func (t *FuncStack) Push(ele *Function) {
	t.stack = append(t.stack, ele)
}

// Pop removes and returns the top function.
func (t *FuncStack) Pop() *Function {
	if t.IsEmpty() {
		panic("FuncStack is empty, can't pop")
	}
	top := t.stack[len(t.stack)-1]
	t.stack = t.stack[:len(t.stack)-1]
	return top
}

// Top returns the top function without removing it.
func (t *FuncStack) Top() *Function {
	if t.IsEmpty() {
		panic("FuncStack is empty, can't top")
	}
	return t.stack[len(t.stack)-1]
}

// IsEmpty checks whether the stack has elements.
func (t *FuncStack) IsEmpty() bool {
	return len(t.stack) == 0
}
