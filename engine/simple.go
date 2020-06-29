package engine

import (
	"crawl/fetcher"
	"crawl/types"
	"log"
)

func (e *SimpleEngine) Run(seed ...types.Request) {
	var requests []types.Request
	for _, r := range seed {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		result, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, result.Requests...)
		itemCount:=0
		for _, item := range result.Items {
			log.Printf("Got item %v,Count is %d", item,itemCount)
			itemCount++
		}
	}
}
func worker(r types.Request) (types.ParseResult, error) {
	body, err := fetcher.Fetcher(r.Url)
	log.Printf("Fetching %s", r.Url)
	if err != nil {
		log.Printf("Fetcher error :"+"fetching url is %s :%v", r.Url, err)
		return types.ParseResult{}, err
	}
	return r.ParseFunc(body), nil
}

type SimpleEngine struct {
}
