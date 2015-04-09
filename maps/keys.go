package maps

import (
	"reflect"
	"sort"
)

type Keys interface {
	sort.Interface
	Add(k interface{}) bool
	Remove(k interface{}) bool
	Clear()
	Get(index int) interface{}
	GetAll() []interface{}
	Search(k interface{}) (index int, contains bool)
	ElemType() reflect.Type
	CompareFunc() func(interface{}, interface{}) int8
}

func NewKeys(compareFunc func(interface{}, interface{}) int8, elemType reflect.Type) Keys {
	return &myKeys{
		container:   make([]interface{}, 0),
		compareFunc: compareFunc,
        elemType:    elemType,
    }
}

type myKeys struct {
	container   []interface{}
	compareFunc func(interface{}, interface{}) int8
	elemType    reflect.Type
}

func (keys *myKeys) Len() int {
	return len(keys.container)
}

func (keys *myKeys) Less(i, j int) bool {
	return keys.compareFunc(keys.container[i], keys.container[j]) == -1
}

func (keys *myKeys) Swap(i, j int) {
	keys.container[i], keys.container[j] = keys.container[j], keys.container[i]
}

func (keys *myKeys) isAcceptableElem(k interface{}) bool {
	if k == nil {
		return false
	}

	return reflect.TypeOf(k) == keys.elemType
}

func (keys *myKeys) Add(k interface{}) bool {
	ok := keys.isAcceptableElem(k)
	if !ok {
		return false
	}

	keys.container = append(keys.container, k)
	sort.Sort(keys)
	return true
}

func (keys *myKeys) Remove(k interface{}) bool {
	index, contains := keys.Search(k)

	if !contains {
		return false
	}

	keys.container = append(keys.container[0:index], keys.container[index+1:]...)
	return true
}

func (keys *myKeys) Search(k interface{}) (index int, contains bool) {
	if !keys.isAcceptableElem(k) {
		contains = false
		return
	}

	index = sort.Search(keys.Len(), func(i int) bool {
		return keys.compareFunc(keys.container[i], k) >= 0
	})

	if index < keys.Len() && keys.container[index] == k {
		contains = true
	}
	return
}

func (keys *myKeys) Clear() {
	keys.container = make([]interface{}, 0)
}

func (keys *myKeys) Get(index int) interface{} {
	if index >= keys.Len() {
		return nil
	}
	return keys.container[index]
}

func (keys *myKeys) GetAll() []interface{} {
	snapshot := make([]interface{}, keys.Len())
	for _, v := range keys.container {
		snapshot = append(snapshot, v)
	}
	return snapshot
}

func (keys *myKeys) ElemType() reflect.Type {
	return keys.elemType
}

func (keys *myKeys) CompareFunc() func(interface{}, interface{}) int8{
	return keys.compareFunc
}
