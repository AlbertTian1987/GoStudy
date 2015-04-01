package set
import (
    "bytes"
    "fmt"
)

type Set interface {
    Add(e interface{}) bool
    AddAll(other Set)
    Remove(e interface{})
    Clear()
    Contains(e interface{})bool
    Len() int
    Equals(other Set) bool
    Elements() []interface{}
    String() string
}

type HashSet struct {
    m map[interface{}]bool
}



func NewHashSet() *HashSet{
    return &HashSet{m:make(map[interface{}]bool)}
}

func (this *HashSet) Add(e interface{}) bool{
    if !this.m[e] {
        this.m[e] = true
        return true
    }
    return false
}

func (this *HashSet) Remove(e interface{}){
    delete(this.m,e)
}

func (this *HashSet) Clear(){
    this.m = make(map[interface{}]bool)
}

func (this *HashSet) Contains(e interface{}) bool{
    return this.m[e]
}

func (this *HashSet) Len() int{
    return len(this.m)
}

func (this *HashSet) Equals(other Set) bool{
    if other == nil {
        return false
    }

    if this.Len()!=other.Len() {
        return false
    }

    for key := range this.m{
        if !other.Contains(key) {
            return false
        }
    }

    return true
}

func (this *HashSet) Elements() []interface{}{
    initialLen := this.Len()
    snapShot := make([]interface{},initialLen)
    actualLen := 0
    for key:= range this.m {
        if actualLen < initialLen {
            snapShot[actualLen] = key
        }else {
            snapShot = append(snapShot,key)
        }
        actualLen++
    }

    if actualLen < initialLen {
        snapShot = snapShot[:actualLen]
    }
    return snapShot
}

func (this *HashSet) String() string{
    var buf bytes.Buffer
    buf.WriteString("Set{")
    first := true
    for key := range this.m{
        if first {
            first = false
        }else {
            buf.WriteString(" ")
        }
        buf.WriteString(fmt.Sprintf("%v",key))
    }
    buf.WriteString("}")
    return buf.String()
}


func (this *HashSet) AddAll(other Set){
    if other == nil {
        return
    }

    for _,e := range other.Elements() {
        this.Add(e);
    }
}

func IsSuperSet(this,other Set) bool{
    if other == nil {
        return false
    }

    thisLen := this.Len()
    otherLen := other.Len()

    if thisLen == 0 || thisLen == otherLen {
        return false
    }

    if thisLen > 0 && otherLen == 0 {
        return true
    }

    for _,v := range other.Elements(){
        if !this.Contains(v) {
            return false
        }
    }

    return true
}


func Union(this,other,result Set){
    result.AddAll(this)
    if other == nil {
        return
    }
    result.AddAll(other)
}

func Intersect(this,other,result Set){
    if other == nil {
        result = nil
        return
    }
    for _,v := range other.Elements(){
        if this.Contains(v) {
            result.Add(v)
        }
    }
}

func Difference(this,other,result Set){
    if other == nil {
        result.AddAll(this)
        return
    }

    for _,v := range this.Elements(){
        if !other.Contains(v) {
            result.Add(v)
        }
    }
}

func SymmetricDifference(this,other,result Set){
    if other == nil {
        result.AddAll(this)
        return
    }
    Difference(this,other,result)
    Difference(other,this,result)
}