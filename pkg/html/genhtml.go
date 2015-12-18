// genhtml
package genhtml

import (
	//	"fmt"
	"os"
	"strconv"
)

func Savestrtofile(namef string, str string) int {
	file, err := os.Create(namef)
	if err != nil {
		// handle the error here
		return -1
	}
	defer file.Close()

	file.WriteString(str)
	return 0
}

// ----------------- функции генерации html page
//-- генерация ячейки таблицы в html
func tablecell(str string) string {
	return "<TD>" + str + "</TD>" + "\n"
}

//-- генерация ссылки в html
func Link(str string, url string) string {
	return "<a href=\"" + url + "\" >" + str + "</a> <br>" + "\n"
}

//-- генерация строки таблицы в html
func gentablestroka(str []string) string {
	res0 := ""
	for i := 0; i < len(str); i++ {
		res0 += tablecell(str[i])
	}
	return "<TR>" + "\n" + res0 + "</TR>" + "\n"
}

func Htmlpage(surl []string) string {
	zagol := "НОВОСТИ"
	begstr := "<html>\n <head>\n <meta charset='utf-8'>\n <title>" + zagol + "</title>\n </head>\n <body>\n"
	bodystr := ""
	for i := 0; i < len(surl); i++ {
		ch := strconv.Itoa(i)
		bodystr += Link("новость "+ch, surl[i])
	}
	//	bodystr := genhtmltable0(datas, zagol, keys)
	endstr := "</body>\n" + "</html>"
	return begstr + bodystr + endstr
}

//--------------------
