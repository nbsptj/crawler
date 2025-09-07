package dongchedi

import (
	"crawler/engine"
	"crawler/shared/cssselector"
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

const SERIEL_SELECTOR = "div.jsx-1226022265 > a"

func ParseSeriel(req *engine.Request) *engine.Result {
	res := &engine.Result{}

	if doc, suc := cssselector.ParseToDoc(req.Data); suc {
		sel := doc.Find(SERIEL_SELECTOR)
		sel.Each(func(i int, s *goquery.Selection) {
			if href, exist := s.Attr("href"); !exist {
				fmt.Printf("dongchedi: parse seriel can not get information href, selection is %v\n", s.Text())
			} else {
				var carseries string
				if title, exist := s.Attr("title"); exist {
					res.Items = append(res.Items, title)
					carseries = title
				} else {
					carseries = "unknown"
				}

				payload := make(map[string]any)
				payload["carseries"] = carseries
				res.Requests = append(res.Requests, engine.Request{
					Url:     DONG_CHE_DI_URL + href,
					Payload: payload,
					Parser:  engine.NilParser{},
				})
			}
		})
	}

	return res
}
