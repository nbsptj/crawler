package main

import (
	"fmt"
	"io/fs"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const baseUrl = "https://www.dongchedi.com"

func main() {
	resp, err := http.Get(baseUrl)
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
			getSeries(baseUrl + href)
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
			href, _ := s.Attr("href")
			getInformationHtml(baseUrl + href)
		} else {
			fmt.Printf("title not exist for %v", s)
		}
	})
}

func getInformationHtml(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("get from url: %s, error: %v", url, err)
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Printf("parse information page error: %v", err)
		return
	}

	sel := doc.Find("a.tw-float-right.tw-clear-both")
	firstSel := sel.First()
	if firstSel == nil {
		fmt.Printf("can't find 'look up complete param config' link!")
		return
	}

	href, exist := sel.Attr("href")
	if !exist {
		fmt.Printf("can't find 'look up complete param config' element's href attribute!")
		return
	}

	detailUrl := baseUrl + href
	resp, err = http.Get(detailUrl)
	if err != nil {
		fmt.Printf("get from url: %s, err: %v", detailUrl, err)
		return
	}
	defer resp.Body.Close()

	bytes, err := httputil.DumpResponse(resp, true)
	if err != nil {
		fmt.Printf("dump response error: %v", err)
		return
	}

	filename := fmt.Sprintf("html/%s.html", url[strings.LastIndex(url, "/")+1:])
	_, err = os.Create(filename)
	if err != nil {
		fmt.Printf("create file error: %v\n", err)
		return
	}

	err = os.WriteFile(filename, bytes, fs.ModeAppend)
	if err != nil {
		fmt.Printf("write file error: %v\n", err)
		return
	}
}
