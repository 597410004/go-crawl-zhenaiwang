package main

import (
	"crawl/engine"
	"crawl/scheduler"
	"crawl/types"
	"crawl/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleSchduler{},
		WorkerCount: 100,
	}
	e.Run(
		types.Request{
			Url:       "http://www.zhenai.com/zhenghun",
			ParseFunc: parser.ParserCityList,
		})

}
