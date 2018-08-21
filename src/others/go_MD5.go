// go_MD5
package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	Md5Inst := md5.New()
	Md5Inst.Write([]byte("zhangsan"))
	Result := Md5Inst.Sum([]byte(""))
	fmt.Printf("%v", Result)

	fmt.Printf("%v", string(Result))

	fmt.Printf("%x\n\n", Result)

}
