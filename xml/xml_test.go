package xml
import "testing"


func Test_xml(t *testing.T){

    obj := Decode_Servers()
    t.Log(obj)

    t.Log("XMLName:",obj.XMLName)
    t.Log("Svs:",obj.Svs)
    t.Log("Version:",obj.Version)
    t.Log("description:",obj.Description)
    for _,v:=range obj.Svs{
        t.Log("Key=",v.ServerName.Key)
        t.Log("Value=",v.ServerName.Value)
    }
}

func Test_Encode(t *testing.T){

    Encode_Servers()

}