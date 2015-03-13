package regexp
import (
    "net/http"
    "fmt"
    "io/ioutil"
    "regexp"
    "strings"
)

func ParseBaidu(){

    resp,err:=http.Get("http://www.baidu.com")

    checkError(err)

    defer resp.Body.Close()

    body,err := ioutil.ReadAll(resp.Body)
    checkError(err)

    src:=string(body)

    //将所有的HTML标签全部转化为小写
    re,_:=regexp.Compile(`\<[\S\s]+?\>`)
    src = re.ReplaceAllStringFunc(src,strings.ToLower)
    //去除style
    re,_=regexp.Compile(`\<style[\S\s]+?\</style\>`)
    src = re.ReplaceAllString(src,"")
    //去除script
    re,_=regexp.Compile(`\<script[\S\s]+?\</script\>`)
    src = re.ReplaceAllString(src,"")
    //去掉所有尖括号内的HTML代码，并替换为换行符
    re,_ =regexp.Compile(`\<[\S\s]+?\>`)
    src = re.ReplaceAllString(src,"\n")
    //去除连续的换行符
    re,_ =regexp.Compile(`\s{2,}`)
    src = re.ReplaceAllString(src,"\n")

    fmt.Println(strings.TrimSpace(src))

}


func FindInString() {

    a:= "I am learning Go Language"
    re,_ := regexp.Compile(`[a-z]{2,4}`)

    one:=re.Find([]byte(a))
    fmt.Println("Find:",string(one))

    fmt.Println("FindString:",re.FindString(a))

    fmt.Println("FindAllString:",re.FindAllString(a,-1))


    fmt.Println("FindStringIndex",re.FindStringIndex(a))
    fmt.Println("FindAllStringIndex",re.FindAllStringIndex(a,-1))

    re,_ = regexp.Compile(`(\w+)\W?`)
    submatch := re.FindAllStringSubmatch(a,-1)
    for _,v := range submatch{
        fmt.Println("len",len(v))
        fmt.Println(v[1])
    }


    src:=`
    call hello alice
    hello bob
    call hello eve
    `
    pat := regexp.MustCompile(`(?m)(call)\s+(?P<cmd>\w+)\s+(?P<arg>.+)\s*$`)

    res := []byte{}

    for _,s:= range pat.FindAllStringSubmatchIndex(src,-1){

        fmt.Println("s=",s)
        res = pat.Expand(res,[]byte(`$cmd('$arg')\n`),[]byte(src),s)
    }

    fmt.Println(string(res))

}

func checkError(err error) {
    if err!=nil {
        fmt.Println(err)
        return
    }
}
