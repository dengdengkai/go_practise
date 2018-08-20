package main
import (
//	    "runtime"
        "fmt"
		"time"
		"strconv"    //导入strconv包，strconv.Itoa(i)转换为字符串
)
/*主动出让或发生阻塞会跳转协程*/
var ch chan int //设置缓存
func main() {
ch = make(chan int)
	//协程1
	 go func() {
		 for i := 1; i < 100; i++ {
			 if i==10 {
//			 runtime.Gosched()//主动要求出让cpu
              <-ch    //读取chan但是因为chan初始化为空，没有读取的所以会阻塞
			 }
		   fmt.Println("routine 1:" +strconv.Itoa(i))
		 }
	 }()

	 //协程2   函数外加括号代表匿名
       go func() {
           for i := 100 ; i < 200; i++ {
             fmt.Println("routine 2:" +strconv.Itoa(i))
           }

		   ch <- 1 //协程2运行完后往chan缓存中写入1
       }()
  time.Sleep(time.Second)
//返回一个time.C这个管道，1秒(time.Second)后会在此管道中放入一个时间点(time.Now())
}
