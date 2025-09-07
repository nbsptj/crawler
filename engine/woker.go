package engine

import "log"

// Work fetch content from Request's Url, then Parse by Request's Parser.
func Work(r *Request) *Result {
	b, err := Fetch(r.Url)
	if err != nil {
		log.Printf("Worker: error fetching url %s:%v", r.Url, err)
		return &Result{}
	}

	r.Data = b

	return r.Parser.Parse(r)
}
