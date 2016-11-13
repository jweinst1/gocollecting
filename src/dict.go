package gocollecting



//struct that implements a dictionary.
type Dictionary struct {
	items map[string]interface{}
}

//function that returns a new dictionary instance.
func newDictionary() Dictionary {
	return Dictionary{make(map[string]interface{})}
}

//sets any item under an any type.
func (self *Dictionary) set(key string, value interface{}) {
	self.items[key] = value
}

//gets the value associated with some key.
func (self *Dictionary) get(key string) interface{} {
	val, has := self.items[key]
	if has {
		return val
	} else {
		return nil
	}
}

//checks if a dictionary contains a key.
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

//updates one dictionary with the keys and values of another
func (self *Dictionary) update(other Dictionary) {
	for key, val := range other.items {
		self.items[key] = val
	}
}

//Only updates a key and value if it doesn't already exist in the dictionary
func (self *Dictionary) fuse(other Dictionary) {
	for key, val := range other.items {
		if !self.contains(key) {
			self.items[key] = val
		}
	}
}

