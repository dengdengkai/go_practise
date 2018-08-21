// regex   go语言标准库内提供了regexp包
package main

import (
	"fmt"
	"regexp"
)

func main() {
	//isok, _ := regexp.Match("[a-zA-Z]{3}", []byte("zhl"))
	//isok, _ := regexp.MatchString("[a-zA-Z]{3}", "z0hl2")
	/*
		reg := regexp.MustCompile(`[a-zA-Z]{2}`)   //注意``是1前面的符号
		result := reg.FindAllString("zhl", -1)
	*/
	reg := regexp.MustCompile(`^z.*l$`) //匹配z开始，l结尾的字符串
	result := reg.FindAllString("zhhmghmfgml03l", -1)
	fmt.Printf("%v\n", result)

	//捕获去掉了正则中匹配的
	reg2 := regexp.MustCompile(`^z(.{1})(.{1})(.*)l$`)
	result2 := reg2.FindAllStringSubmatch("zhangsanl", -1)
	fmt.Printf("%v\n", result2)

}
