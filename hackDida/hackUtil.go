package main
import (
    "crypto/md5"
    "encoding/hex"
    "time"
    "strings"
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "bytes"
)
const (
    key1 = "BOOKING_APP"
    key2 = "kx123456789012345678901234567890"
    host = "211.151.134.222:9010"
    token = "YOUR_TOKEN"
    user_cid = "YOUR_USER_CID"
)

var Parameter = map[string][]string{
    "user_cid": {"f31c32a4-bda0-48c3-b96d-104b28c8fe6a"},
    "token": {"40790825-1939-4019-8ea9-be9b13abfae7"},
    "actid":{"booking_app"},
    "mobiletype":{"2"},
    "version":{"2.2.0.didapinche_taxi_tengcent"},
}

func GetMd5(str string) string {
    m := md5.New()
    m.Write([]byte(str))
    return hex.EncodeToString(m.Sum(nil))
}

func GetVKey() {
    const layout = "20060102030405"
    now := time.Now()
    ts := now.Format(layout)

    temp := strings.ToUpper(GetMd5(key1+ts))
    vkey := strings.ToUpper(GetMd5(temp+key2))

    Parameter["vkey"] = []string{vkey}
    Parameter["ts"] = []string{ts}
}

func GetVersion() (*http.Response, error) {
    GetVKey()
    version_url := fmt.Sprintf("http://%s/V3/Common/getVersion", host)
    return http.PostForm(version_url, Parameter)
}

func GetUserSimpleInfo(cid string) (*http.Response, error) {
    GetVKey()
    version_url := fmt.Sprintf("http://%s/V3/User/getUserSimpleInfo", host)
    Parameter["queried_user_cid"] = []string{cid}
    defer delete(Parameter, "queried_user_cid")
    return http.PostForm(version_url, Parameter)
}

type Geo struct {
    Latitude float64
    Longitude float64
}

var WangJingSoHo Geo = Geo{39.997362, 116.479959}
var BaiDuDaSha Geo = Geo{40.051158, 116.301085}

func GetNearbyBookingRideList(geo Geo) (*http.Response, error) {
    GetVKey()
    version_url := fmt.Sprintf("http://%s/V3/BookingDriver/getNearbyBookingRideList", host)
    Parameter["center_latitude"] = []string{fmt.Sprintf("%f", geo.Latitude)}
    Parameter["center_longitude"] = []string{fmt.Sprintf("%f", geo.Longitude)}
    Parameter["filter_by"] = []string{"all"}
    Parameter["order_by"] = []string{"plan_start_time_asc"}
    Parameter["page"] = []string{"1"}
    Parameter["page_size"] = []string{"10"}

    defer resetMap()
    return http.PostForm(version_url, Parameter)
}


func resetMap() {
    Parameter = map[string][]string{
        "user_cid": {"f31c32a4-bda0-48c3-b96d-104b28c8fe6a"},
        "token": {"40790825-1939-4019-8ea9-be9b13abfae7"},
        "actid":{"booking_app"},
        "mobiletype":{"2"},
        "version":{"2.1.1"},
    }
}

type CommResp struct {
    Code int `json:"code"`
    Message string `json:"message"`
}

func (this *CommResp) String() string {
    return fmt.Sprintf("code:%d,message:%s", this.Code, this.Message)
}

type UserInfo struct {
    Cid string `json:"cid"`
    Name string `json:"name"`
    Gender string `json:"gender"`
    Phone string `json:"phone"`
    Role int `json:"role"`
    CurrentRole int `json:"currentrole"`
    DriverInfo json.RawMessage `json:"driverinfo"`
}

func (this UserInfo) String() string{
    var buf bytes.Buffer
    buf.WriteString("{")
    buf.WriteString(this.Name+" ")
    buf.WriteString(this.Phone+" ")
    if this.DriverInfo!=nil {
        buf.WriteString(string(this.DriverInfo)+"\n")
    }else {
        buf.WriteString("\n")
    }
    return buf.String()
}

type NearbyResultItem struct {
    UserInfo UserInfo `json:"passenger_user_info"`
}

type NearbyResult struct {
    List []NearbyResultItem `json:"list"`
}


func main() {
    //"4796af72-af6a-11e4-9c1d-782bcb4cf8d8"
    resp, err := GetNearbyBookingRideList(WangJingSoHo)
    if err != nil {
        fmt.Println(err)
        return
    }
    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
        return
    }

    result := new(NearbyResult)
    json.Unmarshal(data, result)

    if len(result.List) >0 {
        for _, item := range result.List {

            fmt.Println(item.UserInfo.String())
        }
    }

}