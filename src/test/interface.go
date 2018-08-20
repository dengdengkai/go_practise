package main
 import (
     "fmt"
)
 type Phone interface {
    call()
 }
 type NokiaPhone struct {
 }
 func (nokiaPhone NokiaPhone) call() {
 
 
 fmt.Println("I am Nokia,I can call you!")
 }
 type IPhone struct{
 }
 func (iPhone IPhone) call() {
 
 fmt.Println("I am iPhone,I can call you")
 }
 func main() {
 
 var phone Phone
    phone = new (NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()  
 }

 /*
在上面的例子中，我们定义了一个接口Phone，接口里面有一个方法call()。然后我们在main函数里面定义了一个Phone类型变量，并分别为之赋值为NokiaPhone和IPhone。然后调用call()方法，输出结果如下：*/
