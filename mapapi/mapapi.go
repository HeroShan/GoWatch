package mapapi

import(
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"regexp"
)
var Ip string

type IpJson struct{
	Address  string			`json:"address"`
	Content  JsContent 		`json:"content"`
	Status	 int 			`json:"status"`
}

type JsContent struct{
	Address  		string 				`json:"address"`
	Address_detail  JsAddress_detail 	`json:"address_detail"`
	Point 			JsPoint 			`json:"point"`
}

type JsAddress_detail struct{
	City			string 		`json:"city"`
	City_code		int 		`json:"city_code"`
	District		string		`json:"district"`
	Province		string		`json:"province"`
	Street			string		`json:"street"`
	Street_number	string		`json:"street_number"`
}

type JsPoint struct{
	X	string 		`json:"x"`
	Y	string		`json:"y"`
}


func Getpoint(Ip string) (JsPoint,string) {
	
	url := "http://api.map.baidu.com/location/ip?ip="+Ip+"&ak=vKgTlhoXlrEFFmLGBD4E04Gdv29ZYMxX&coor=bd09ll"
	resp, err := http.Get(url); if err != nil{
		fmt.Println(err)
	}
	body , errred := ioutil.ReadAll(resp.Body); if errred != nil{
		fmt.Println(errred)
	}
	//fmt.Println(string(body))
	jsonip := IpJson{}
	jserr := json.Unmarshal(body,&jsonip); if jserr != nil{
		fmt.Println("jserr is error::",jserr)
	}
	//fmt.Println(jsonip)
	return jsonip.Content.Point,jsonip.Content.Address
}

func Getarea(Point JsPoint) string {
	
	url := "http://api.map.baidu.com/geocoder?location="+Point.X+","+Point.Y+""
	resp, err := http.Get(url); if err != nil{
		fmt.Println(err)
	}
	body , errred := ioutil.ReadAll(resp.Body); if errred != nil{
		fmt.Println(errred)
	}
	reg := regexp.MustCompile(`<formatted_address>(.*?)</formatted_address>`)

	result := reg.FindAllStringSubmatch(string(body), -1)
	return result[0][1]
}