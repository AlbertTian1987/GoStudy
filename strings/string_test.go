package strings
import "testing"


func Test_StringFmt(t *testing.T){
    t.Log(StringsFmt1("a","b","c","d",1,false,5.754))
    t.Log(StringsFmt2("a","b","c","d",1,false,5.754))

    t.Log(StringsFmt1("aaaaaaa"))
    t.Log(StringsFmt2("aaaaaaa"))
}