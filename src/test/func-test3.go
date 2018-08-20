package main
import "fmt"

func getSequence()  func() int {
	i := 0
	return  func() int {
	
		i+=1
		return i
	}
}
    func main() {

	    /* nextNumber 为一个函数，函数i为0 */
	    nextNumber := getSequence()

	    /*调用 nextNumber函数，1     变量自增1 返回 */
	    fmt.Println(nextNumber())
	    fmt.Println(nextNumber())
	    fmt.Println(nextNumber())

    
    }
