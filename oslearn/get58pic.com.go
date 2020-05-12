package main

import(
	"net/http"
	"net/url"
	"io/ioutil"
	"strconv"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"time"
)

type Jar struct {
	cookies []*http.Cookie
}

func (jar *Jar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	jar.cookies = cookies
}

func (jar *Jar) Cookies(u *url.URL) []*http.Cookie {
	return jar.cookies
}

var over = make(chan bool)

func getcatepage(i int, page chan int){
	var url string
	jar := new(Jar)

	client := &http.Client{nil, nil, jar, 99999999999992}

	url = "https://www.58pic.com/piccate/3-0-0-ty1-p"+strconv.Itoa(i)+".html"

	//cookie set
	_, err := client.PostForm(url, map[string][]string{
		"last_view_picid": {"%2210633350%7C5Yay5rOw5Lyg5aqS%7C1589848864%7C67f3dd336cd7157bec77b8440aefa607%22"}, 
		"m_uid": {"a044a022-02de-4358-868e-e23ef405184f"},
		"source_lookp" : {"%2235970806-%22"},
		"showAd:826fb21dea6cf579abf0ad6efa28d8b0" : {"%22w6SIEgLKiJOIC5HVD3fKoJGYnMzImJfKzwe5y5y4nZLHyMyWywq5zwzHmJHKogiWiIWIywr5zxj3AxnLCL2Pzci9iJeIlcj3DxjUiJOXlcjZAg26x6rPBwvZiJOInciSiMXHC6rFC5HVD423Aw4LiJOXntG8mJq3nZqZFv3%3D%22"},
		"qt_visitor_id" : {"%22826fb21dea6cf579abf0ad6efa28d8b0%22"},
		"is_pay1589212800" : {"is_pay1589212800"},
		"_auth_dl_" : {"MTA2MzMzNTB8MTU4OTg0ODg2NHxjMGZkZThmMjA2MGQ1NGNkZDdjM2Y3MWRiYTQwZmI0Mg%3D%3D"},
		"WEBPARAMS" : {"is_pay=1"},
		"ssid" : {"%225eb9f0a0575007.42402924%22"},
		"sns" : {"%7B%22token%22%3A%7B%22access_token%22%3A%22C0DCD74C25D5D81A57CBE5A6FDC48E0B%22%2C%22expires_in%22%3A%227776000%22%2C%22refresh_token%22%3A%2298C52D92BECC1935EE0BA983F80697AC%22%2C%22openid%22%3A%22F60CE2918B966DADDB18CEA9023B01B7%22%7D%2C%22type%22%3A%22qq%22%7D"},
		"qiantudata2018jssdkcross" : {"%7B%22distinct_id%22%3A%22172065004c965f-00b3ec23fd168a-3c3f590d-2073600-172065004caab0%22%2C%22props%22%3A%7B%22latest_traffic_source_type%22%3A%22%E7%9B%B4%E6%8E%A5%E6%B5%81%E9%87%8F%22%2C%22latest_referrer%22%3A%22%22%2C%22latest_referrer_host%22%3A%22%22%2C%22latest_search_keyword%22%3A%22%E6%9C%AA%E5%8F%96%E5%88%B0%E5%80%BC_%E7%9B%B4%E6%8E%A5%E6%89%93%E5%BC%80%22%7D%7D"},
	})
	if err != nil {
		panic(err.Error())
	}

	resp, geterr := client.Get(url)
	if geterr != nil{
		panic(geterr)
	}
	document, goqueryerr := goquery.NewDocumentFromResponse(resp)
	if goqueryerr != nil{
		panic(goqueryerr)
	}
	document.Find(".qtw-card").Find("a").Each(func(_ int, item *goquery.Selection) {
		linkTag := item
		link, _ := linkTag.Attr("href")
		 reg, _ := regexp.Compile(`//(.*?)newpic\/(.*?).html`)
		 //www.58pic.com/newpic/35777526.html
		 imgurls := reg.FindStringSubmatch(link)
		 if imgurls != nil{
			//寻找素材下载点
			imgurl := "https:"+imgurls[0]
			resp, imgerr := client.Get(imgurl)
			if imgerr != nil{
				panic(imgerr)
			}
			body,_ := ioutil.ReadAll(resp.Body)
			fmt.Println(string(body))
			select {

			case force := <-page:
				if force < 13300 {
					fmt.Println(force)
				}
			case force := <-page:
				if force >= 13300 {
					over <- true
				}
			}
		 }
    	
	})



}

func main(){
	var i int
	page := make(chan int, 5)
	for i = 1; i < 270; i++ {
		select {
		case page <- i:
		case <-time.After(time.Second * 10):
		case tu := <-over :
		if tu == true{
			return 
			}
		}
	 go getcatepage(i, page)
	}
	
	defer fmt.Println("Over")	
}