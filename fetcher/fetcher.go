package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func Fetcher(url string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln("NewRequest is err ", err)
		return nil, fmt.Errorf("NewRequest is err %v\n", err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")

	//返送请求获取返回结果
	resp, err := client.Do(req)

	//直接用http.Get(url)进行获取信息，爬取时可能返回403，禁止访问
	//resp, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("Error: http Get, err is %v\n", err)
	}
	defer resp.Body.Close()
	//if resp.StatusCode != http.StatusOK {
	//	fmt.Println("error code,", resp.StatusCode)
	//	fmt.Println(resp.Body)
	//	return nil, fmt.Errorf("error code  %d", resp.StatusCode)
	//}
	e := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

//根据网站编码类型来编码网站
func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
