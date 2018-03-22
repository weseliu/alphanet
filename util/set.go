package util

type Set struct {
	m map[interface{}]bool
}

func NewSet() *Set {
	return &Set{
		m: map[interface{}]bool{},
	}
}

func (Self *Set) Add(item interface{}) {
	Self.m[item] = true
}

func (Self *Set) Remove(item interface{}) {
	delete(Self.m, item)
}

func (Self *Set) Has(item interface{}) bool {
	_, ok := Self.m[item]
	return ok
}

func (Self *Set) Len() int {
	return len(Self.List())
}

func (Self *Set) Clear() {
	Self.m = map[interface{}]bool{}
}

func (Self *Set) IsEmpty() bool {
	return Self.Len() == 0
}

func (Self *Set) List() []interface{} {
	list := make([]interface{}, 0)
	for item := range Self.m {
		list = append(list, item)
	}
	return list
}

func (Self *Set) Traverse(callback func(item interface{}) bool) {
	for item := range Self.m {
		if callback != nil {
			if callback(item) == false {
				break
			}
		}
	}
}
