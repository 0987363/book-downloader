package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
//	"golang.org/x/net/html/charset"
)

func ExampleScrape() {
//	doc, err := goquery.NewDocument("http://www.biquge.com.tw/17_17275/")
//	doc, err := goquery.NewDocument("http://book.qidian.com/info/1009704712")
	doc, err := goquery.NewDocument("http://www.999wx.com/article/76/29041/Default.shtml")

//	doc, err := goquery.NewDocument("http://read.qidian.com/chapter/rCo0cx0UxW08kjk6dUsm_A2/8yNI9L7ye4tMs5iq0oQwLQ2")
//	doc, err := goquery.NewDocument("http://www.biquge.com.tw/17_17275/7633258.html")
//	doc, err := goquery.NewDocument("http://book.zongheng.com/chapter/672340/37010929.html")
//	doc, err := goquery.NewDocument("http://www.999wx.com/article/19/32372/11226626.shtml")
	if err != nil {
		log.Fatal(err)
	}

	
	// request dir
	t := doc.Find("div.volume").Find("ul.cf")
	fmt.Println("length:", t.Length())
	for i := 0; i < t.Length(); i++ {
		fmt.Println(t.Eq(i).Text())
		d, exi := t.Eq(i).Attr("href")
		fmt.Println(d, exi)
	}
	

	/*
	// request text
	ttt := doc.Find("head").Find("meta")
	fmt.Println("length:", ttt.Length())
	charset := "utf-8"
	for i := 0; i < ttt.Length(); i++ {
		s := ttt.Get(i).Attr
		for _, a := range s {
			key := strings.ToLower(a.Key)
			val := strings.ToLower(a.Val)
			if strings.Contains(key, "charset") {
				charset = val
				fmt.Println(charset)
				break
			}
			if strings.Contains(val, "charset") {
				result := strings.Split(val, "charset=")

				charset = result[1]
				fmt.Println("charset:", charset)
			}
		}
	}

	t := doc.Find("div")
	fmt.Println("length:", t.Length())
	c := 0
	s := ""
	for i := 0; i < t.Length(); i++ {
		if (len(t.Eq(i).Text()) < 6000) {
			continue
		}
		if c == 0 || c > len(t.Eq(i).Text()) {
			fmt.Println(len(t.Eq(i).Text()))
			c = len(t.Eq(i).Text())
			s = t.Eq(i).Text()
		}
	}
	fmt.Println("text len:", len(s))
//	fmt.Println(s)
		*/


}

func main() {
	ExampleScrape()
}
