package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

//定义一个书的结构体
type BookInformationStruct struct {
	Id string
	Name string
	Publisher string
	Price float64
	Author string

}

//定义一个图书管理的接口，提供各种方法
type BookManagement interface {
	findBookByPublisher()          //通过出版社查询图书
	findBookByAuthor()			//通过作者查询书籍
	findBookByName()			//通过书名查询信息
}

//定义一个实现 管理接口 的类
type BookManagementStruct struct {
	//定义一个map用于储存所有的书的信息，书的id是k，书的信息也就是结构体是v
	BookInformationMapAll map[string]BookInformationStruct

}
//
//对类进行初始化
//func NewBookManagentStrct(BookInformationPublisher map[string]string,BookInformationAuthor map[string]string,BookInformationName map[string]BookInformationStruct) *BookManagementStruct{
//	return &BookManagementStruct{BookInformationPublisher,BookInformationAuthor,BookInformationName}
//}

//为 管理类 添加方法
//------------------------------------------------------------------------|----------------------------------------------------------\\
//通过出版社查找图书信息得到ID映射
func (t BookManagementStruct) findBookByPublisher(str string) []BookInformationStruct {
	//a := make(map[string]BookInformationStruct)
	a := make([]BookInformationStruct,0,50)
	for _, v := range t.BookInformationMapAll{
		if strings.TrimSpace(v.Publisher) == str{
			//a[str] = v
			a = append(a,v)
		}
	}
	fmt.Println(a)
	return a
}

//通过作者查找图书信息得到ID映射
func (t BookManagementStruct) findBookByAuthor(str string) []BookInformationStruct {
	//a := make(map[string]BookInformationStruct)
	a := make([]BookInformationStruct,0,50)
	for _, v := range t.BookInformationMapAll{
		if strings.TrimSpace(v.Author) == str{
			//a[str] = v
			a = append(a,v)
		}
	}
	return a
}

//通过书名查找图书信息ID到详细信息的映射
func (t BookManagementStruct) findBookByName(str string) []BookInformationStruct{
	//a := make(map[string]BookInformationStruct)
	a:=make([]BookInformationStruct,0,50)

	for _, v := range t.BookInformationMapAll{
		if v.Name == str{
			//a[str] = v
			a = append(a,v)
		}
	}
	return a
}
/*

-----                                     -这里可以添加想要的查询方法比如通过其他的方式进行查询-                         --------

*/

//------------------------------------------------------------------------^|^----------------------------------------------------------\\




//添加传入参数的对findBookByname的调用
func (a BookManagementStruct) Byname(c *gin.Context) {
	//c.ShouldBindJSON()
	name := c.Param("name")
	//c.Query()

	//fmt.Println(a.BookInformationName)
	c.JSON(http.StatusOK,a.findBookByName(name))
}


//为findbookbypublisher添加参数输入的方法
func (a BookManagementStruct) ByPublisher(c *gin.Context)  {
	//name := c.Param("publisher")
	name := c.Query("publisher")
	c.JSON(http.StatusOK,a.findBookByPublisher(name))
}


//添加findBookByAuthor的参数传入的方法
func (a BookManagementStruct) ByAuthor(c *gin.Context)  {
	name := c.Param("author")
	c.JSON(http.StatusOK,a.findBookByAuthor(name))
}

