package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const url = "https://www.dongchedi.com"

func main() {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}

	sel := doc.Find("div.car-list_model__aT0i_  > a")
	sel.Each(func(i int, s *goquery.Selection) {
		if title, exist := s.Attr("title"); exist {
			href, _ := s.Attr("href")
			fmt.Printf("got car %q url %q\n", title, href)
			getSeries(url + href)
		} else {
			fmt.Printf("attr not exist for %d: %v\n", i, s)
		}
	})
}

func getSeries(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("got series from %s error: %v\n", url, err)
		return
	}
	if resp.StatusCode != 200 {
		fmt.Printf("got error statuscode %d for %q\n", resp.StatusCode, url)
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Printf("parse series page error: %v", err)
		return
	}

	sel := doc.Find("div.jsx-1226022265 > a")
	sel.Each(func(i int, s *goquery.Selection) {
		if title, exist := s.Attr("title"); exist {
			fmt.Printf("got series %d: %q\n", i, title)
		} else {
			fmt.Printf("title not exist for %v", s)
		}
	})
}
