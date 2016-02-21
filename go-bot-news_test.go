package main

import (
	"fmt"
	//	"strings"
	"testing"
)

func TestGetNewsRBC(t *testing.T) {
	ln := ListNews{name: "RBC_RT", url: "http://rt.rbc.ru/"}

	n := GetNews(ln)
	fmt.Println(n)
	//		n = DelNullNews(n)
	//		str += Htmlpage(ln[i], n) + "<br><br>"

}
