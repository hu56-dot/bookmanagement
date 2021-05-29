package main

import (
	"Project07/controller"
	"Project07/readfile"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("----------------------")
	//读取文件开始处理
	a := controller.BookManagementStruct{BookInformationMapAll: readfile.ReadFile("C:\\Users\\wfl19\\GOworkspace\\src\\Project07\\book.txt")}

	//开始搭建web框架
	r := gin.Default() //返回默认的路由引擎
	r.GET("/byPublisher", a.ByPublisher)
	r.GET("/byName/:name", a.Byname) //通过出版社进行查询
	r.GET("/byAuthor/:author", a.ByAuthor)

	err := r.Run(":9090")
	if err != nil {
		return
	}
}
