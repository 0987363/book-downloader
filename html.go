package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func HtmlScrape(rule *Rule, url string) {
//	doc, err := goquery.NewDocument("http://www.biquge.com.tw/17_17275/")
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	if len(rule.Class) 

	t := doc.Find("div.volume").Find("ul.cf").Find("a")
	fmt.Println("length:", t.Length())
	for i := 0; i < t.Length(); i++ {
		fmt.Println(t.Eq(i).Text())
		d, exi := t.Eq(i).Attr("href")
		fmt.Println(d, exi)
	}
}

