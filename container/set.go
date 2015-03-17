package container
import (
    "bytes"
    "fmt"
)

type HashSet struct {
    m map[interface{}]bool
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

func (this *HashSet) Equals(other *HashSet) bool{
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

func (this *HashSet) IsSuperSet(other *HashSet) bool{
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