// go-bot-news
package main

import (
	"fmt"
	"go-bot-news/pkg"
	"go-bot-news/pkg/html"
	"golang.org/x/net/html/charset"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type News struct {
	url     string //урл новости
	title   string // заголовок новости
	content string // содержимое новости
}

//инициализация лог файла
func InitLogFile(namef string) *log.Logger {
	file, err := os.OpenFile(namef, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", os.Stderr, ":", err)
	}
	multi := io.MultiWriter(file, os.Stdout)
	LFile := log.New(multi, "Info: ", log.Ldate|log.Ltime|log.Lshortfile)
	return LFile
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

//удаление повторных элементов в массиве
func delpovtor(s []string) []string {
	fl := false
	st := make([]string, 0)
	st = append(st, s[0])
	for i := 0; i < len(s); i++ {
		fl = true
		for j := 0; j < len(st); j++ {
			if s[i] == st[j] {
				fl = false
			}
		}
		if fl {
			st = append(st, s[i])
		}
	}
	return st
}

func printarray(s []string) {
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
	return
}

//получение новостей
func GetNewsUrl(url string) []string {
	//	var ss []string
	if url == "" {
		return make([]string, 0)
	}
	body := gethtmlpage(url)
	shtml := string(body)

	// <a rel="nofollow" href="/likes/e1678230/" class="share" data-url="http://echo.msk.ru/news/1678230-echo.html" data-title="Новое уголовное дело о ремонте кораблей Северного флота поступило в суд">
	snewsmusor, _ := pick.PickAttr(&pick.Option{&shtml, "a", nil}, "data-url")
	snews := make([]string, 0)
	for i := 0; i < len(snewsmusor); i++ {
		if strings.Contains(snewsmusor[i], "-echo.htm") && (strings.Contains(snewsmusor[i], "/news/")) {
			snews = append(snews, snewsmusor[i])
		}
	}

	//	printarray(delpovtor(snews))

	return delpovtor(snews)
}

func (this *News) GetNews() {

	if this.url == "" {
		return
	}
	body := gethtmlpage(this.url)
	shtml := string(body)

	//	<meta property="og:title" content="Новости / 17 декабря, 16:31 | Путин утверждает, что  никогда  не  обсуждал  с  региональными  лидерами расследование конкретных  уголовных  дел" />

	stitle, _ := pick.PickAttr(&pick.Option{&shtml, "meta", &pick.Attr{"property", "og:title"}}, "content")
	if len(stitle) > 0 {
		this.title = stitle[0]
	}

	//	<meta property="og:description" content="
	//В   том числе дела об убийстве    Бориса  Немцова. «Следствие должно установить, как бы долго оно ни продолжалось. Это преступление должно быть расследовано и участники должны быть наказаны, кто бы это ни был, — сказал глава государства." />
	scont, _ := pick.PickAttr(&pick.Option{&shtml, "meta", &pick.Attr{"property", "og:description"}}, "content")
	this.content = scont[0]

	return
}

func Htmlpage(sn []News) string {
	zagol := "ГРАББЕР НОВОСТЕЙ"
	begstr := "<html>\n <head>\n <meta charset='utf-8'>\n <title>" + zagol + "</title>\n </head>\n <body>\n"
	//	<h3 id=”Razdel2”> Раздел2 </h3>
	bodystr := "<h1 align=\"center\"><a name=\"MainPage\"> ГРАББЕР НОВОСТЕЙ </a></h1><br>" + "<TABLE align=\"center\" border=\"1\">"
	for i := 0; i < len(sn); i++ {
		bodystr += "<TR> <TD width=\"350\"> <b>" + genhtml.Link(sn[i].title, sn[i].url) + "</b></TD>" + "<TD width=\"550\">" + sn[i].content + "" + "<a href=\"#MainPage\"> В начало </a>" + "</TD> </TR>"
	}
	bodystr += "</TABLE>"
	endstr := "</body>\n" + "</html>"
	return begstr + bodystr + endstr
}

func main() {
	fmt.Println("Starting программы")
	url := "http://echo.msk.ru/"
	n := make([]News, 0)

	ss := GetNewsUrl(url)
	for i := 0; i < len(ss); i++ {
		n = append(n, News{url: ss[i]})
	}

	for i := 0; i < len(n); i++ {
		n[i].GetNews()
	}

	str := Htmlpage(n)
	genhtml.Savestrtofile("news.html", str)

	fmt.Println("Ending программы")
}
