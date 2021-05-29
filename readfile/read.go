package readfile

import (
	"Project07/controller"
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

//定义一个字符串转结构体的函数
//将读取的信息转化为bookInformationMapAll类型
func toStruct(str string) controller.BookInformationStruct{
	var b controller.BookInformationStruct
	//首先分割掉分号
	a :=strings.Split(str,";") //a1 的类型是切片
	s1 := strings.Replace(a[len(a)-1], "\n", "", -1)
	s2 := strings.Replace(s1, "\r", "", -1)
	a[len(a)-1] = s2
	a1 :=a

	//a3 := strings.Replace(a2, "\r", "", -1)
	var bMap = make(map[string]string)
	for i := 0;i<len(a1);i++{
		bMap[(strings.Split(string(a1[i]),":"))[0]] = (strings.Split(string(a1[i]),":"))[1]
	}

	f,err:=strconv.ParseFloat(bMap["单价"],64)
	if err!=nil{
		fmt.Println("string to float64 error")
		f = 0
	}
	b =controller.BookInformationStruct{Id: bMap["ID"], Name: bMap["书名"], Publisher: bMap["出版社"], Author: bMap["作者"], Price:f }
	//fmt.Println(b)
	return b
}





//定义一个对文件的基本读取
func ReadFile(filePath string)map[string]controller.BookInformationStruct {
	var d controller.BookInformationStruct
	c := make(map[string]controller.BookInformationStruct)
	fmt.Println("start read the specified file:\n-\n-")
	f , err := os.Open(filePath)
	if err != nil {
		fmt.Printf("open file err = %v",err)
	}
	defer fmt.Println("read complete , have closed")
	defer f.Close()    //完成后要关闭文件
	reader := bufio.NewReader(f)    //创建一个缓冲
	//读文件
	for{
		str , err := reader.ReadString('\n')
		if err == io.EOF{     //io.EOF表示文件的最后一个值，每次读取挨个字符，
			// 都会有一个err，读得到err=nil，或者其他，
			break
		}
		if 0 == len(str) || str == "\r\n" {
			continue
		}
		d = toStruct(str)
		c[d.Id] = d
	}
	fmt.Println("读取基本结果如下：")
	return c
}