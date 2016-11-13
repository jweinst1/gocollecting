package main

import "fmt"



//the stack data structure
type Stack struct {
	items []interface{}
}

func newStack() Stack {
	return Stack{make([]interface{}, 0)}
}

func (self *Stack) push(value interface{}) {
	self.items = append(self.items, value)
}

func main() {
	f := newStack()
	f.push(7)
	fmt.Println(f)
}