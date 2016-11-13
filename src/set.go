package main

import "fmt"

//structure for sets

type Set struct {
	items map[string]bool
}

//makes and returns a new Set
func newSet() Set {
	return Set{make(map[string]bool)}
}

//adds an item to the set
func (self *Set) add(item string) {
	self.items[item] = true
}

//checks if an item is in the set
func (self *Set) contains(item string) bool {
	_, has := self.items[item]
	return has
}

//removes an item from the set
func (self *Set) remove(item string) {
	delete(self.items, item)
}

//makes a copy of the current set
func (self *Set) copy() Set {
	nwset := newSet()
	for key, _ := range self.items {
		nwset.add(key)
	}
	return nwset
}

//returns a new set that represents the intersection of both sets
func (self *Set) intersection(other Set) Set {
	nwset := newSet()
	for key, _ := range other.items {
		if self.contains(key) {
			nwset.add(key)
		}
	}
	return nwset
}

func main() {
	a := newSet()
	a.add("ee")
	a.add("rr")
	b := newSet()
	b.add("ee")
	b.add("g")
	fmt.Println(a.intersection(b))
}