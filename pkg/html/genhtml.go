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

//-- генерация таблицы в html: первый параметр это заголовок таблицы, второй параметр [[],[],...] - строки таблицы, keys - массив указывающий в каком порядке выводить в таблицу
//func genhtmltable0(datas map[string]DataTelMans, zagol string, keys []string) string {
//	res := ""
//	//res = map gentablestroka str

//	titletab := []string{"ФИО РГ",
//		"номер телефона",
//		"ФИО менеджера",
//		"продол-ть",
//		"кол-во звонков",
//		"кол-во уник. тел.",
//		"кол-во результ. звонков",
//		"продол-ть уник.",
//		"ср. время звонка"}
//	tabletitle := gentablestroka(titletab)

//	tabledata := ""
//	//for key, _ := range datas {
//	for i := 0; i < len(keys); i++ {
//		key := keys[i]
//		str := []string{
//			datas[key].fio_rg,
//			key,
//			datas[key].fio_man,
//			sec_to_s(datas[key].totalsec),
//			strconv.Itoa(datas[key].totalzv),
//			strconv.Itoa(datas[key].kolunik),
//			strconv.Itoa(datas[key].kolresult),
//			sec_to_s(datas[key].secresult),
//			sec_to_s(devidezero(datas[key].totalsec, datas[key].totalzv))}

//		tabledata += gentablestroka(str)
//	}

//	zagolovok := "<CAPTION>" + zagol + "</CAPTION>\n"
//	tablehtml := zagolovok + tabletitle + tabledata
//	return "<TABLE>" + "\n" + "<TABLE BORDER>\n" + tablehtml + res + "</TABLE>"
//}

//--------------------
