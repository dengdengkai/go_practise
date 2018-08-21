// regex2.go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	url := "https://movie.douban.com/subject/27605698/"
	resp, err := http.Get(url)
	if err != nil {

		panic(err)
	}
	defer resp.Body.Close()
	sHtml, _ := ioutil.ReadAll(resp.Body)

	//匹配电影名
	//  \s*代表0到多个空格符
	reg := regexp.MustCompile(`<span\s*property="v:itemreviewed">(.*)</span>`)
	result := reg.FindAllStringSubmatch(string(sHtml), -1)
	fmt.Printf("v\n", result)
	fmt.Println()
	fmt.Println(result[0][1])
	for _, v := range result {
		fmt.Println(v[0])
		fmt.Println(v[1])

	}
	//匹配评分
	reg1 := regexp.MustCompile(`<strong class="ll rating_num" property="v:average">(.*)</strong>`)
	result1 := reg1.FindAllStringSubmatch(string(sHtml), -1)
	fmt.Printf("v\n", result1)
	fmt.Println()
	fmt.Println(result[0][1])

	for _, v := range result1 {
		fmt.Println(v[0])
		fmt.Println(v[1])

	}

}
