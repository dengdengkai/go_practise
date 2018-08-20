 package  main
import "fmt"
func main(){

sliceTest()
twoDimenAr()

}
func sliceTest(){
arr  :=[]int{1,2,3,4,5}
 s :=arr[:]
for e := range s {
fmt.Println(s[e])
}

s1 := make([]int,3)
for e := range s1{

   fmt.Println(s1[e])

  }

}

func twoDimenAr(){


/*五行二列非规则数组*/  

var a = [][]int{{0,0},{1,2},{2},{4,0}}
var i,j int
for i = 0; i < len(a); i++ {

    for j = 0; j< len(a[i]); j++ {
	
	
     	fmt.Printf("a[%d][%d] = %d \n", i , j ,a[i][j])
	}


}




}
