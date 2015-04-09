package maps

import (
	"reflect"
)

type OrderedMap interface {
	Get(key interface{}) interface{}
	Put(key interface{}, elem interface{}) (interface{}, bool)
	Remove(key interface{}) interface{}
	Clear()
	Len() int
	Contains(key interface{}) bool
	FirstKey() interface{}
	LastKey() interface{}
	HeadMap(toKey interface{}) OrderedMap
	SubMap(fromKey interface{}, toKey interface{}) OrderedMap
	TailMap(fromKey interface{}) OrderedMap
	Keys() []interface{}
	Elems() []interface{}
	ToMap() map[interface{}]interface{}
	KeyType() reflect.Type
	ElemType() reflect.Type
}

type myOrderedMap struct {
	keys     Keys
	keyType  reflect.Type
	elemType reflect.Type
	m        map[interface{}]interface{}
}

func NewOrderedMap(compareFunc func(interface{}, interface{}) int8, keyType, elemType reflect.Type) OrderedMap {
	return &myOrderedMap{
		keys:     NewKeys(compareFunc, elemType),
		keyType:  keyType,
		elemType: elemType,
		m:        make(map[interface{}]interface{}),
	}
}

func (orderedMap *myOrderedMap) Get(key interface{}) interface{} {
	if !orderedMap.isAcceptableKey(key) {
		return nil
	}
	return orderedMap.m[key]
}

func (orderedMap *myOrderedMap) Put(key interface{}, elem interface{}) (interface{}, bool) {

	if !orderedMap.isAcceptableKey(key) || !orderedMap.isAcceptableElem(elem) {
		return nil, false
	}

    oldElem,ok:=orderedMap.m[key]
    orderedMap.m[key] = elem
    if !ok {
        orderedMap.keys.Add(key)
    }
    return oldElem,true
}

func (orderedMap *myOrderedMap) Remove(key interface{}) interface{} {
	if !orderedMap.isAcceptableKey(key) {
		return nil
	}
	if !orderedMap.Contains(key) {
		return nil
	}
	orderedMap.keys.Remove(key)
	removedElem := orderedMap.Get(key)
	delete(orderedMap.m, key)
	return removedElem
}

func (orderedMap *myOrderedMap) Clear() {
	orderedMap.keys.Clear()
	orderedMap.m = make(map[interface{}]interface{})
}

func (orderedMap *myOrderedMap) Len() int {
	return len(orderedMap.m)
}

func (orderedMap *myOrderedMap) Contains(key interface{}) bool {
    _,ok:=orderedMap.m[key]
	return ok
}

func (orderedMap *myOrderedMap) FirstKey() interface{} {
	if orderedMap.Len() == 0 {
		return nil
	}
	return orderedMap.keys.Get(0)
}

func (orderedMap *myOrderedMap) LastKey() interface{} {
	if orderedMap.Len() == 0 {
		return nil
	}
	return orderedMap.keys.Get(orderedMap.Len() - 1)
}

func (orderedMap *myOrderedMap) HeadMap(toKey interface{}) OrderedMap {
	return orderedMap.SubMap(nil, toKey)
}

func (orderedMap *myOrderedMap) SubMap(fromKey interface{}, toKey interface{}) OrderedMap {

	newOMap := &myOrderedMap{
		keys:     NewKeys(orderedMap.keys.CompareFunc(), orderedMap.keyType),
		keyType:  orderedMap.keyType,
		elemType: orderedMap.elemType,
		m:        make(map[interface{}]interface{}),
	}

	omapLen := orderedMap.Len()

	if omapLen == 0 {
		return newOMap
	}

	startIndex, c := orderedMap.keys.Search(fromKey)
	if !c {
		startIndex = 0
	}

	endIndex, c := orderedMap.keys.Search(toKey)
	if !c {
		endIndex = omapLen
	}

	var key, elem interface{}
	for i := startIndex; i < endIndex; i++ {
		key = orderedMap.keys.Get(i)
		elem = orderedMap.m[key]
		newOMap.Put(key, elem)
	}

	return newOMap

}

func (orderedMap *myOrderedMap) TailMap(fromKey interface{}) OrderedMap {
	return orderedMap.SubMap(fromKey, nil)
}

func (orderedMap *myOrderedMap) Keys() []interface{} {
	return orderedMap.keys.GetAll()
}

func (orderedMap *myOrderedMap) Elems() []interface{} {
	sortedKeys := orderedMap.Keys()
	elems := make([]interface{}, orderedMap.Len())
	for _, key := range sortedKeys {
		elems = append(elems, orderedMap.m[key])
	}
	return elems
}

func (orderedMap *myOrderedMap) ToMap() map[interface{}]interface{} {
	mm := make(map[interface{}]interface{})
	sortedKeys := orderedMap.Keys()
	for _, key := range sortedKeys {
		mm[key] = orderedMap.m[key]
	}
	return mm
}

func (orderedMap *myOrderedMap) KeyType() reflect.Type {
	return orderedMap.keys.ElemType()
}

func (orderedMap *myOrderedMap) ElemType() reflect.Type {
	return orderedMap.elemType
}

func (orderedMap *myOrderedMap) isAcceptableElem(elem interface{}) bool {
	if elem == nil {
		return false
	}

	return reflect.TypeOf(elem) == orderedMap.elemType
}

func (orderedMap *myOrderedMap) isAcceptableKey(key interface{}) bool {
	if key == nil {
		return false
	}

	return reflect.TypeOf(key) == orderedMap.keyType
}
