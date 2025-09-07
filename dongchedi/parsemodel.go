package dongchedi

import (
	"bytes"
	"crawler/engine"
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

const LINK_SELECTOR = "div.car-list_model__aT0i_  > a"

func ParseModel(req *engine.Request) *engine.Result {
	result := &engine.Result{}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(req.Data))
	if err != nil {
		fmt.Printf("dongchedi: new document from reader error: %v\n", err)
		return result
	}

	sel := doc.Find(LINK_SELECTOR)
	sel.Each(func(i int, s *goquery.Selection) {
		if href, exist := s.Attr("href"); !exist {
			fmt.Printf("dongchedi: parse model can not get serial href, selection is %v\n", s.Text())
		} else {
			result.Requests = append(result.Requests, engine.Request{
				Url:    DONG_CHE_DI_URL + href,
				Parser: engine.NilParser{},
			})

			if title, exist := s.Attr("title"); exist {
				result.Items = append(result.Items, title)
			}
		}
	})

	return result
}
