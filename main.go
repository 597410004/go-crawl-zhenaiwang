package main

import (
	"crawl/engine"
	"crawl/types"
	"crawl/zhenai/parser"
)

func main() {
	engine.Run(types.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParserCityList,
	})

}
