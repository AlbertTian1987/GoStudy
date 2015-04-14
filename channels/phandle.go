package main
import "fmt"

type Person struct {
    Name string
    Age int
    Addr Address
}

type Address struct {
    City string
    Distinct string
}

type PersonHandler interface {
    Batch(origs <-chan Person) <-chan Person
    Handle(orig *Person)
}

type PersonHandlerImpl struct {
}

func (this PersonHandlerImpl) Batch(origs <-chan Person) <-chan Person{
    dest:=make(chan Person,cap(origs))
    go func(){
        for p:= range origs{
            this.Handle(&p)
            dest<-p
        }
        close(dest)

    }()
    return dest
}


func (this PersonHandlerImpl) Handle(orig *Person){
    orig.Addr.Distinct = "朝阳"
}

func main() {
    persionHandler:= getPersonHandler()
    origs := make(chan Person,5)
    dest:=persionHandler.Batch(origs)
    fetchPerson(origs)
    sign := savePerson(dest)
    <-sign
}

func getPersonHandler() PersonHandler{
    var handler PersonHandlerImpl
    return handler
}

func fetchPerson(origs chan<- Person){
    go func(){
        for i:=0;i<5;i++{
            p:= Person{fmt.Sprintf("Allen%d",i),20,Address{"Beijing","海淀"}}
            origs<-p
        }
        close(origs)
    }()
}

func savePerson(dest <-chan Person) <-chan int{
    sign:=make(chan int,1)
    go func() {
        for {
            p,ok:= <-dest
            if !ok {
                break
            }
            fmt.Println(p)
        }

        sign<-1
    }()
    return sign;
}