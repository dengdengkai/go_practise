package main 
import (
          "fmt"
 )

 /*   定义结构体*/  
 type Circle struct {
  radius float64
 
 }

 func main() {
 var c1 Circle
 c1.radius = 10.00
 fmt.Println("Area of Circle(c1) = ", c1.getArea())

 }
 
 //  该methd属于circle类型对象中的方法
 func (c Circle) getArea() float64 {
 
 
 return  3.14 * c.radius  *  c.radius
 }

