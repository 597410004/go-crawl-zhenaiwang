package parser

import (
	"crawl/types"
	"regexp"
)

const urlReg = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParserCityList(contents []byte) types.ParseResult {
	re, _ := regexp.Compile(urlReg)
	matches := re.FindAllSubmatch(contents, -1)
	result := types.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, "City"+string(m[2]))
		url:=string(m[1])
		result.Requests = append(result.Requests,
			types.Request{Url: url, ParseFunc: func(bytes []byte) types.ParseResult {
				return ParserCity(bytes,url)
			}})
	}
	return result
}
