package parser

import (
	"io/ioutil"
	"testing"
)

func TestParserCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParserCityList(contents)
	const resultSize = 470
	if len(result.Items) != resultSize {
		t.Errorf("result should hava %d requests; but had %d",
			resultSize, len(result.Requests))
	}
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedCities := []string{
		"City阿坝", "City阿克苏", "City阿拉善盟",
	}
	for i, v := range expectedUrls {
		if result.Requests[i].Url != v {
			t.Errorf("expected url #%d: %s; but was %s",
				i, v, result.Requests[i].Url)
		}
	}
	for i, v := range expectedCities {
		if result.Items[i].(string) != v {
			t.Errorf("expected city #%d: %s,but was %s",
				i, v, result.Items[i].(string))
		}
	}
}
