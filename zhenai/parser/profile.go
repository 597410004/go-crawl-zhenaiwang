package parser

import (
	"crawl/types"
	"crawl/zhenai/model"
	"regexp"
	"strconv"
)

var genderRegex = regexp.MustCompile(`<td width="180"><span class="grayL">性别：</span>([^<]+)</td>`)
var ageRegex = regexp.MustCompile(`<td width="180"><span class="grayL">年龄：</span>([^<]+)</td>`)
var heightRegex = regexp.MustCompile(`<td width="180"><span class="grayL">身   高：</span>([^<]+)</td>`)
var moneyRegex = regexp.MustCompile(`<td><span class="grayL">月   薪：</span>([^<]+)</td>`)

func ParseProfile(contents []byte, name string) types.ParseResult {
	profile := model.Profile{}
	//fmt.Println(string(contents))
	if age, err := strconv.Atoi(extractString(ageRegex, contents)); err == nil {
		profile.Age = age
	}
	if gender := extractString(genderRegex, contents); gender != "" {
		profile.Gender = gender
	}
	if height, err := strconv.Atoi(extractString(heightRegex, contents)); err == nil {
		profile.Height = height
	}
	if money := extractString(moneyRegex, contents); money != "" {
		profile.Money = money
	}
	profile.Name = name
	result := types.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}
func extractString(re *regexp.Regexp, content []byte) string {
	match := re.FindSubmatch(content)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}
