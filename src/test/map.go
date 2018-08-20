package main
import  "fmt"
func  main(){
      var countryCapitalMap  map[string]string  /*map集合创建 var 变量 map[键类型]值类型      */
      countryCapitalMap = make(map[string]string) //初始化map,不初始化则为空nil
      /*map 插入key - value 对,各个国家对应首都*/
	  countryCapitalMap ["France"] = "Paris"

	  countryCapitalMap ["Italy"] = "罗马"
	  countryCapitalMap ["Japan"] = "东京"
	  countryCapitalMap ["India"] = "新德里"
	  for country := range  countryCapitalMap {
		  fmt.Println(country,"首都是",countryCapitalMap[country])
	  }
/*查看元素在集合中是否存在*/
captial, ok := countryCapitalMap ["美国"] 
fmt.Println(captial)
fmt.Println(ok)
if(ok){
   fmt.Println("美国的首都是",captial)
}else {
   fmt.Println("美国的首都不存在")
 }




}
