package main

import "fmt"

//dictionary data structure file

type Dictionary struct {
	items map[string]interface{}
}

func newDictionary() Dictionary {
	return Dictionary{make(map[string]interface{})}
}

//sets any item under an any type
func (self *Dictionary) set(key string, value interface{}) {
	self.items[key] = value
}

func (self *Dictionary) get(key string) interface{} {
	val, has := self.items[key]
	if has {
		return val
	} else {
		return nil
	}
}

func (self *Dictionary) contains(key string) bool {
	_ , has := self.items[key]
	return has
}

//deletes a key in the dictionary
func (self *Dictionary) del(key string) {
	delete(self.items, key)
}

//returns amount of pairs in the Dictionary
func (self *Dictionary) length() int {
	return len(self.items)
}

//returns a slice of strings for the keys in the dictionary
func (self *Dictionary) keys() []string {
	keylst := make([]string, len(self.items))
	i := 0
	for key, _ := range self.items {
		keylst[i] = key
		i++
	}
	return keylst
}

//returns a slice of any typed values from the dictionary
func (self *Dictionary) values() []interface{} {
	vallst := make([]interface{}, len(self.items))
	i := 0
	for _, val := range self.items {
		vallst[i] = val
		i++
	}
	return vallst
}

func main() {
	f := newDictionary()
	f.set("foo", 6)
	f.set("doom", 3)
	fmt.Println(f.values())
}