package main
import (
    "testing"
    "time"
    "net/http"
    "net/http/httptest"
    "io/ioutil"
    "fmt"
    "encoding/json"
)

func TestGetMd5(t *testing.T){
    t.Log(GetMd5("dsad123123"))
    t.Log(GetMd5("dsad123123"))
}

func TestTimeFormat(t *testing.T){
    const layout = "20060102030405"
    now := time.Now()
    t.Log(now.Format(layout))
    t.Log(now.UTC().Format(layout))
}

func TestGetVKey(t *testing.T){
    GetVKey()
    t.Log(Parameter)
}

func TestHttp(t *testing.T){
    ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        r.ParseForm()
        t.Log(r.FormValue("mobiletype"))
    }))
    defer ts.Close()

    res, err := http.PostForm(ts.URL,Parameter)
    if err != nil {
        t.Fatal(err)
    }
    greeting, err := ioutil.ReadAll(res.Body)
    defer res.Body.Close()
    if err != nil {
        t.Fatal(err)
    }

    t.Log(fmt.Sprintf("%s",greeting))
}

func TestGetVersion (t *testing.T){
    resp,err:=GetVersion()
    if err!=nil {
        t.Fatal(err)
    }else {
        ReadResp(resp,t)
    }
}


func TestGetUserSimpleInfo (t *testing.T){
    resp,err:=GetUserSimpleInfo("469c6750-4ba2-49ce-8baa-43e8640229ac")
    if err!=nil {
        t.Fatal(err)
    }else {
        ReadResp(resp,t)
    }
}


func TestGetNearbyBookingRideList (t *testing.T){
    resp,err:=GetNearbyBookingRideList(BaiDuDaSha)
    if err!=nil {
        t.Fatal(err)
    }else {
        cr := new(CommResp)
        json.Unmarshal([]byte(ReadResp(resp,t)),cr)
        t.Log(cr)
    }
}

func ReadResp(resp *http.Response,t *testing.T)string{
    data,err:=ioutil.ReadAll(resp.Body)
    if err!=nil {
        t.Fatal(err)

        return err.Error()
    }else {
        str:=string(data)
        t.Log(str)
        return str
    }
}