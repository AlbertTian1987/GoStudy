package reflect
import (
    "testing"
    "reflect"
    "fmt"
)
type DD struct {
    Key string
    Value int
}

func Test_Reflect(t *testing.T) {

    D := DD{Key:"hello",Value:90}

    s := reflect.ValueOf(&D).Elem()

    typeOfT := s.Type()
    for i := 0; i < s.NumField(); i++ {
        f := s.Field(i)
        t.Log(fmt.Sprintf("%d: %s %s = %v\n", i,
        typeOfT.Field(i).Name, f.Type(), f.Interface()))
    }


    i := 27
    ir := reflect.ValueOf(&i)

    p := ir.Elem()
    t.Log("ir is ",ir.Kind()," and CanSet ",ir.CanSet())
    t.Log("p is ",p.Kind()," and CanSet ",p.CanSet())

    kk := reflect.Indirect(reflect.ValueOf(&i))
    t.Log("kk is ",kk.Kind()," and CanSet ",kk.CanSet(),"kk value",kk.Interface())
    kk.SetInt(876)
    t.Log("and i change to ",i)

}

func indirectType(v reflect.Type) reflect.Type {
    switch v.Kind() {
        case reflect.Ptr:
        return indirectType(v.Elem())
        default:
        return v
    }
}