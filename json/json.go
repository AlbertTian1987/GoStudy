package json
import (
    "encoding/json"
    . "github.com/bitly/go-simplejson"
    "fmt"
)

type User struct {
    Id int
    Name string
    Age int
}

type Class struct {
    Name string
}

type ComplexUser struct {
    UserInfo User
    OwnClass ClassSlice
    OwnClass2 []Class
}

type ClassSlice []Class

var USER_JSON = `{"id":123,"name":"Albert","age":24,"OwnClasses":[{"name":"class1"},{"name":"class2"}]} `
var Class_JSON = `[{"name":"class1"},{"name":"class2"}]`
var Complex_Json = `{"UserInfo":{"id":1,"name":"emellend","age":2},"Ownclass":[{"name":"hah1"},{"name":"aaa2"}],"Ownclass2":[{"name":"hah2"},{"name":"aaa3"}]}`

func (user *User) Decode_from_json(jsonString string)error{
    return json.Unmarshal([]byte(jsonString),user)
}

func (classes *ClassSlice) Decode_from_json(jsonstring string) error{
    return json.Unmarshal([]byte(jsonstring),classes)
}

func (cpx *ComplexUser) Decode_from_json(jsonString string) error{
    return json.Unmarshal([]byte(jsonString),cpx)
}

var UnKnowJson = `
{
    "test":{
        "array":[1,"2",3],
        "int":10,
        "float":5.150,
        "bignum":9223372036864775807,
        "string":"simplejson",
        "bool":true
    }
}
`

func UnkonwJson(){
    js,_ := NewJson([]byte(UnKnowJson))

    arr,_:=js.Get("test").Get("array").Array()
    i:=js.Get("test").Get("int").MustInt()
    str := js.Get("test").Get("string").MustString()
    fmt.Println("arr=",arr)
    fmt.Println("i=",i)
    fmt.Println("str=",str)
}

type Server struct{
    Int int `json:"-"`
    ServerName string `json:"server_name"`
    Score float64 `json:"score,string"`
    ServerIp string `json:"server_ip,omitempty"`
    ServerId int `json:"id"`
    ConvertId int `json:"conver_id,string"`
    Bool bool `json:"valid"`
}


func EncodeToJson(){
    slice := make([]Server,2)
    slice[0] = Server{0,`s1 "1.0"`,7.65,"",45,123,true}
    slice[1] = Server{1,`s2 "1.0"`,99.8,"127.0.0.1",7,1123,false}

    data,err:=json.MarshalIndent(&slice,"","    ")
    if err!=nil {
        fmt.Println(err)
    }else {
        fmt.Println(string(data))
    }
}
