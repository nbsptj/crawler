package engine

import (
	"testing"
)

func TestFetch(t *testing.T) {

	cases := []struct {
		url string
		err bool
	}{
		{
			url: "https://www.baidu.com",
			err: false,
		},
		{
			url: "https://123",
			err: true,
		},
	}

	for i, c := range cases {
		_, err := Fetch(c.url)

		if c.err && err == nil {
			t.Errorf("#%d: want err but don't got err", i)
			continue
		}

		if !c.err && err != nil {
			t.Errorf("#%d: don't want err but got err: %v", i, err)
			continue
		}
	}

}
