// tsetMysql project main.go
package main
 
import (
   "strings"
	"time"
 
	"github.com/astaxie/beego/orm"
	. "github.com/soekchl/myUtils"
	_ "github.com/go-sql-driver/mysql"
)
 
type Student struct {
	Id         int64     `orm:"auto"`
	Name       string    `orm:"size(10)"`
	CreateTime time.Time `orm:"type(timestamp)"`
	Note       string    `orm:"type(text);null"`
}
 
func (this *Student) TableName() string {
	return "student" // 数据库创建表名
}
 
func (this *Student) Read(cols ...string) (err error) {
	o := orm.NewOrm()
	if err = o.Read(this, cols...); err == nil {
		return nil
	}
	return err
}
 
func (this *Student) Insert() (int64, error) {
	o := orm.NewOrm()
	return o.Insert(this)
}
 
func (this *Student) Update(cols ...string) (int64, error) {
	o := orm.NewOrm()
	return o.Update(this, cols...)
}
 
func (this *Student) Delete() (int64, error) {
	o := orm.NewOrm()
	return o.Delete(this)
}


func init() {

	dbhost := "localhost"
	dbport := "3306"
	dbname := "test"
	dbuser := "root"
	dbpass := "123456"
 
	dsn := strings.Join([]string{dbuser, ":", dbpass, "@tcp(", dbhost, ":", dbport, ")/", dbname, "?charset=utf8&loc=Asia%2FShanghai"}, "")

		
orm.RegisterDataBase("default", "mysql", dsn, 30)


orm.RegisterModel(new(Student))
/*orm.RegisterDataBase("default", "mysql", "root:123456@127.0.0.1:3306/test?charset=utf8",30)
*/

orm.RunSyncdb("default", false, true) // 表存在就不再创建
 }
 
func main() {
 
	var stu Student = Student{Name: "Luke"}
 
	_, err := stu.Insert() // 插入
	if err != nil {
		Error(err)
		return
	}
	Notice(stu.Name, " Id is :", stu.Id)
 
	stu.Name = "Test"
	_, err = stu.Update("name") // 只更新名字
	if err != nil {
		Error(err)
		return
	}
	Notice("Name is change: ", stu.Name)
 
	var stu1 Student
	stu1.Id = 1
	err = stu1.Read() // 读取id为1 的数据
	if err != nil {
		Error(err)
		return
	}
 
	Notice(stu1)
 
}

