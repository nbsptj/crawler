package engine

import "fmt"

type SimpleEngine struct{}

func (*SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	requests = append(requests, seeds...)

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		result := Work(&r)

		requests = append(requests, result.Requests...)

		for _, item := range result.Items {
			fmt.Printf("Got item: %v\n", item)
		}
	}
}
