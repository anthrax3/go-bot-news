// go-bot-news
package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"golang.org/x/net/html/charset"
//	"go-bot-price/pkg"
//	"strings"
	
)

type News struct {
	url string  //урл новости
	content string // содержимое новости
}

//получение страницы из урла url
func gethtmlpage(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("HTTP error:", err)
		panic("HTTP error")
	}
	defer resp.Body.Close()
	// вот здесь и начинается самое интересное
	utf8, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	if err != nil {
		fmt.Println("Encoding error:", err)
		panic("Encoding error")
	}
	body, err := ioutil.ReadAll(utf8)
	if err != nil {
		fmt.Println("IO error:", err)
		panic("IO error")
	}
	return body
}

//получение данных товара из магазина Эльдорадо по урлу url
func (this *News) GetNews(url string) {
//	var ss []string
	if url == "" {
		return
	}
	body := gethtmlpage(url)
	shtml := string(body)
	
	fmt.Println(shtml)

//	sname, _ := pick.PickText(&pick.Option{ // текст цены книги
//		&shtml,
//		"div",
//		&pick.Attr{
//			"class",
//			"q-fixed-name no-mobile",
//		},
//	})

//	for i := 0; i < len(sname); i++ {
//		if strings.TrimSpace(sname[i]) != "" { // удаление пробелов
//			ss = append(ss, sname[i])
//		}
//	}

//	this.name = ss[0]

//	sprice, _ := pick.PickText(&pick.Option{&shtml, "span", &pick.Attr{"itemprop", "price"}})

//	ss = make([]string, 0)
//	for i := 0; i < len(sprice); i++ {
//		if strings.TrimSpace(sprice[i]) != "" { // удаление пробелов
//			ss = append(ss, sprice[i])
//		}
//	}

//	if len(ss) > 0 {
//		this.price, _ = strconv.Atoi(ss[0])
//	}

	return
}



func main() {
	fmt.Println("Starting программы")
	url:="http://echo.msk.ru/"
	var n News
	
//	fmt.Println(n.GetNews(url))
	n.GetNews(url)
	fmt.Println(n)
	
	fmt.Println("Ending программы")
}
