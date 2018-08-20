package main
import(
	"fmt"
	"time"
    "strconv"
)

func test_Add(x,y int) {
	z:= x  + y
      fmt.Println(z)
}
func Read(ch chan int){
     value := <-ch
	 fmt.Println("value:" + strconv.Itoa(value))
}

func Write(ch chan int){
ch <-10
}
func  main() {
	ch := make(chan int)
	go Read(ch)
	go Write(ch)
    time.Sleep(time.Second)
}

