package main



//the stack data structure
type Stack struct {
	items []interface{}
}

//returns a new, empty stack
func newStack() Stack {
	return Stack{make([]interface{}, 0)}
}

//pushs a new item onto the stack
func (self *Stack) push(value interface{}) {
	self.items = append(self.items, value)
}

//returns the top item on the stack
func (self *Stack) peek() interface{} {
	return self.items[len(self.items)-1]
}
