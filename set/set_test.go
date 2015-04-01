package set
import (
    "testing"
)


func TestAdd(t *testing.T){

    set1:=NewHashSet()
    set1.Add("a")
    set1.Add("a")
    set1.Add("a")
    set1.Add("a")
    set1.Add("a")
    t.Log("set1.len() = ",set1.Len())

}

func TestRemove(t *testing.T){
    set1:= NewHashSet()
    set1.Remove(1)
    set1.Add(1)
    set1.Remove(1)
    t.Log("set1 contains 1 = ",set1.Contains(1))
    set1.Add("1")
    t.Log(`set1 contains "1" = `,set1.Contains("1"))
}

func TestClear(t *testing.T){
    set1:=NewHashSet()

    set1.Add(1)
    set1.Add(2)
    set1.Add(3)
    set1.Add(4)

    set1.Clear()
    t.Log("set1.len() = ",set1.Len())
}

func TestEqual(t *testing.T){
    set1:=NewHashSet()
    set1.Add(1)
    set1.Add(2)
    set1.Add(3)
    set1.Add(4)

    set2:=NewHashSet()
    set2.Add(1)
    set2.Add(2)
    set2.Add(3)
    set2.Add(4)

    t.Log("set1 equals set2 ",set1.Equals(set2))

    set2.Add("1")
    t.Log("set1 equals set2 ",set1.Equals(set2))
}

func TestElements(t *testing.T){
    set1:=NewHashSet()
    set1.Add(1)
    set1.Add(2)
    set1.Add(3)
    set1.Add(4)
    set1.Add("adad")
    t.Log("set1 Elements = ",set1.Elements())
    t.Log("set1 String = ",set1)

}

func TestIsSuperSet(t *testing.T){
    set1:=NewHashSet()
    set1.Add(1)
    set1.Add(2)
    set1.Add(3)
    set1.Add(4)
    set1.Add("adad")

    set2:=NewHashSet()
    set2.Add(1)
    set2.Add(2)
    set2.Add(3)
    set2.Add(4)

    t.Log("set1 is set2 SuperSet = ",IsSuperSet(set1,set2))

    set2.Add("da")
    t.Log("set1 is set2 SuperSet = ",IsSuperSet(set1,set2))

}

func TestAddAll(t *testing.T){
    set1:=NewHashSet()
    set1.Add(1)
    set1.Add(2)
    set1.Add(3)
    set1.Add(4)
    set1.Add("adad")


    set2 := NewHashSet()
    set2.AddAll(set1)

    t.Log("set1 = ",set1)
    t.Log("set2 = ",set2)

    set1.Remove(2)
    set1.Remove("adad")
    t.Log("set1 = ",set1)
    t.Log("set2 = ",set2)
}

func TestUnion(t *testing.T){
    set1:=NewHashSet()
    set1.Add(1)
    set1.Add(2)
    set1.Add(3)
    set1.Add(4)
    set1.Add("adad")

    set2:=NewHashSet()
    set2.Add(1)
    set2.Add(2)
    set2.Add(3)
    set2.Add(4)
    set2.Add("djxchy")

    result:=NewHashSet()
    Union(set1,set2,result)
    t.Log("set1 and set2 union = ",result)
}

func TestIntersect(t *testing.T){
    set1:=NewHashSet()
    set1.Add(1)
    set1.Add(2)
    set1.Add(3)
    set1.Add(4)
    set1.Add("adad")

    set2:=NewHashSet()
    set2.Add(1)
    set2.Add(2)
    set2.Add(3)
    set2.Add(4)
    set2.Add("djxchy")
    set2.Add("1")

    result:=NewHashSet()
    Intersect(set1,set2,result)
    t.Log("set1 and set2 Intersect = ",result)
    result.Clear()
    Intersect(set1,nil,result)
    t.Log("set1 and nil Intersect = ",result)
}


func TestDifference(t *testing.T){
    set1:=NewHashSet()
    set1.Add(1)
    set1.Add(2)
    set1.Add(3)
    set1.Add(4)
    set1.Add("adad")

    set2:=NewHashSet()
    set2.Add(1)
    set2.Add(2)
    set2.Add(3)
    set2.Add(4)
    set2.Add("djxchy")
    set2.Add("1")

    result:=NewHashSet()
    Difference(set1,set2,result)
    t.Log("set1 and set2 Difference = ",result)

    result.Clear()
    Difference(set2,set1,result)

    t.Log("set2 and set1  Difference = ",result)

    result.Clear()
    Difference(set1,nil,result)
    t.Log("set1 and nil Difference = ",result)
}

func TestSymmetricDifference(t *testing.T){
    set1:=NewHashSet()
    set1.Add(1)
    set1.Add(2)
    set1.Add(3)
    set1.Add(4)
    set1.Add("adad")

    set2:=NewHashSet()
    set2.Add(1)
    set2.Add(2)
    set2.Add(3)
    set2.Add(4)
    set2.Add("djxchy")
    set2.Add("1")

    result:=NewHashSet()
    SymmetricDifference(set1,set2,result)
    t.Log("set1 and set2 SymmetricDifference = ",result)

    result.Clear()
    SymmetricDifference(set2,set1,result)
    t.Log("set2 and set1 SymmetricDifference = ",result)

    result.Clear()
    SymmetricDifference(set1,nil,result)
    t.Log("set1 and nil SymmetricDifference = ",result)
}

func TestSet(t *testing.T){
    set1:= NewHashSet()
    set2:= NewHashSet()
    s := Set(set1)
    s2 := Set(set2)


    s.Add(1)

    t.Log(s)
    t.Log(s.Equals(s2))
}