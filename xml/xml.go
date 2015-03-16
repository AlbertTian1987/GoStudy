package xml

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Recurlyservers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version,attr"`
	Svs         []server `xml:"server"`
	Description string   `xml:",innerxml"`
}

type server struct {
	XMLName    xml.Name   `xml:"server"`
	ServerName serverName `xml:"ServerName"`
	ServerIP   string     `xml:"serverIP"`
}

type serverName struct {
	Key   string `xml:"opt,attr"`
	Value string `xml:",chardata"`
}

func Decode_Servers() *Recurlyservers {
	file, err := os.Open("servers.xml")
	checkError(err)

	defer file.Close()
	data, err := ioutil.ReadAll(file)
	checkError(err)

	v := Recurlyservers{}
	err = xml.Unmarshal(data, &v)
	checkError(err)
	return &v
}

func checkError(err error) {
	if err != nil {
		fmt.Printf("error :v%", err)
		return
	}
}

func Encode_Servers() {

	v := new(Recurlyservers)
	v.Version = "1"
	v.Svs = append(v.Svs, server{ServerName: serverName{"经济首都", "上海"}, ServerIP: "127.0.0.2"})
	v.Svs = append(v.Svs, server{ServerName: serverName{"政治首都", "北京"}, ServerIP: "127.0.0.1"})
	output, err := xml.MarshalIndent(&v, " ", "  ")
	checkError(err)
	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)
}
