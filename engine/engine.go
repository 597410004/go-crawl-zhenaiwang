package engine

import (
	"crawl/fetcher"
	"crawl/types"
	"log"
)

func Run(seed ...types.Request) {
	var requests []types.Request
	for _, r := range seed {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		body, err := fetcher.Fetcher(r.Url)
		log.Printf("Fetching %s", r.Url)
		if err != nil {
			log.Printf("Fetcher error :"+"fetching url is %s :%v", r.Url, err)
			continue
		}
		result := r.ParseFunc(body)
		requests = append(requests, result.Requests...)
		for _, item := range result.Items {
			log.Printf("Got item %v", item)
		}
	}
}
