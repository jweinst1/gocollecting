package main

//data structure very similar to a javascript object
type Object struct {
	items map[string]interface{}
}

//returns a new instance of an object
func newObject() Object {
	return Object{make(map[string]interface{})}
}