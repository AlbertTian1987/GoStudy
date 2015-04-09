package maps

import (
    "reflect"
    "testing"
)

type Item struct {
    ID   int
    Name string
}

func GenetateMap() OrderedMap {
    return NewOrderedMap(
    func(i1 interface{}, i2 interface{}) int8 {
        item1 := i1.(Item)
        item2 := i2.(Item)

        if item1.ID < item2.ID {
            return -1
        } else if item1.ID > item2.ID {
            return 1
        } else {
            return 0
        }
    }, reflect.TypeOf(string("")), reflect.TypeOf(Item{}))
}

func TestGetAndPut(t *testing.T) {
    omap := GenetateMap()

    omap.Put("a",Item{5,"hellp"})
    omap.Put("b",Item{2,"kjhd"})
    t.Log("omap.Len()=",omap.Len())

    oldElem,c :=omap.Put("a",Item{1,"hello world"})
    if c {
        t.Log("oldElem = ",oldElem)
    }

}

func TestElems(t *testing.T){
    omap:=GenetateMap()
    omap.Put("1",Item{})
}
