package parser

import (
	"crawl/types"
	"regexp"
)

const cityReg = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParserCity(contents []byte,url string) types.ParseResult {
	re, _ := regexp.Compile(cityReg)
	matches := re.FindAllSubmatch(contents, -1)
	result := types.ParseResult{}
	for _, m := range matches {
		name:=string(m[2])
		result.Items = append(result.Items, "User"+string(m[2]))
		result.Requests = append(result.Requests,
			types.Request{Url: url, ParseFunc: func(content []byte) types.ParseResult {
				return ParseProfile(content, name)
			}})
	}
	return result
}
