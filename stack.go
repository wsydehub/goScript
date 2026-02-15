package goScript

type ScopeStack struct {
	stack []*Scope
}

func NewScopeStack() *ScopeStack {
	return &ScopeStack{stack: []*Scope{}}
}

func (t *ScopeStack) Push(ele *Scope) {
	t.stack = append(t.stack, ele)
}

func (t *ScopeStack) Pop() *Scope {
	if t.IsEmpty() {
		panic("ScopeStack is empty, can't pop")
	}
	top := t.stack[len(t.stack)-1]
	t.stack = t.stack[:len(t.stack)-1]
	return top
}

func (t *ScopeStack) Top() *Scope {
	if t.IsEmpty() {
		panic("ScopeStack is empty, can't top")
	}
	return t.stack[len(t.stack)-1]
}

func (t *ScopeStack) IsEmpty() bool {
	return len(t.stack) == 0
}

type FuncStack struct {
	stack []*Function
}

func NewFuncStack() *FuncStack {
	return &FuncStack{stack: []*Function{}}
}

func (t *FuncStack) Push(ele *Function) {
	t.stack = append(t.stack, ele)
}

func (t *FuncStack) Pop() *Function {
	if t.IsEmpty() {
		panic("FuncStack is empty, can't pop")
	}
	top := t.stack[len(t.stack)-1]
	t.stack = t.stack[:len(t.stack)-1]
	return top
}

func (t *FuncStack) Top() *Function {
	if t.IsEmpty() {
		panic("FuncStack is empty, can't top")
	}
	return t.stack[len(t.stack)-1]
}

func (t *FuncStack) IsEmpty() bool {
	return len(t.stack) == 0
}
