package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
//	doc, err := goquery.NewDocument("http://www.biquge.com.tw/17_17275/")
	doc, err := goquery.NewDocument("http://book.qidian.com/info/1009704712")
	if err != nil {
		log.Fatal(err)
	}

	t := doc.Find("div.volume").Find("ul.cf")
	fmt.Println("length:", t.Length())
	for i := 0; i < t.Length(); i++ {
		fmt.Println(t.Eq(i).Text())
		d, exi := t.Eq(i).Attr("href")
		fmt.Println(d, exi)
	}


}

func main() {
	ExampleScrape()
}
