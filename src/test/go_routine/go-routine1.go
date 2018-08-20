package main
import ( 
	"fmt"
	"time"
)
func test_Routine() {
      fmt.Println("This is one routine!!")
}
func  main() {
   

	go test_Routine()
    time.Sleep(1000)
}

