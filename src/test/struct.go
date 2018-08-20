package main
import "fmt"

type Books struct {

title string
author string
subject string
book_id int
}

func main(){
fmt.Println(Books{"Go language","www.runoob.com","Go language manuals",6495407})


fmt.Println(Books{title:"Go language",author:"www.runoob.com",subject:"Go language manuals",book_id:6495407})


fmt.Println(Books{title:"Go language",author:"www.runoob.com"})


}
