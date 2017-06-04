package main
//demo:-> https://gobyexample.com/
// golang开源库:mgo nsq etcd consul
import (
	"fmt"
	"strconv"
)

var( //全局变量
	goo string = "daaa"
	doo int = 233
)

const A string  = "abd" //常量字符串
const B   = "阿打算打算" //常量字符串 隐式

func main(){
    	//局部变量
	var name  = "danxingxi"
	var name1  = "go"
	var name2  = "233 go"
	var aa, ab, ac  = "adad","adcsw", "asdasd" //这种赋值方式牛逼

	fmt.Println( name )

	fmt.Println("hello golang " + name , name1, name2)

	fmt.Println(aa , ab, ac)

	fmt.Println(goo, doo)

	var numb  = 22
	var point *int /* 声明指针变量 */

	point = &numb  /* 指针变量的存储地址 */

	fmt.Println("变量地址：", &numb)

	fmt.Println("point 变量储存的指针地址:：", point)

	fmt.Println("使用指针访问值:", *point)

	fmt.Println("point不是空指针：",point!=nil)

	test(23)
}

type Book struct { //结构体(全局变量)
	name string
	year string
	id   int

}

/**
 * 定义方法 参数  返回值
 */
func test(num int) (string, string) {
	fmt.Println("传入参数：", num)
	var book1 Book
	book1.name = "最牛逼的golang教程"
	book1.year = "2017"
	book1.id = 12214
	printBook(book1)
	return "sss","sdsds"  //方法可以返回两个参数
}
/**
 *参数为结构体
 */
func printBook(book Book)  {
	fmt.Printf( "Book name : %s\n", book.name)
	fmt.Printf( "Book year : %s\n", book.year)
	fmt.Printf( "Book id : %d\n", book.id)
	printPBook(&book)
}
/**
 * 参数为结构体指针
 */
func printPBook(bookP *Book)  {
	fmt.Printf( "Book name : %s\n", bookP.name)
	fmt.Printf( "Book year : %s\n", bookP.year)
	fmt.Printf( "Book id : %d\n", bookP.id)
	byteTest()
}

func byteTest()  {
	var byte1  = [4]float32{12.1, 32.2, 2.0, 21}  //数组
	var byte2  = [4]string{}

	fmt.Println("使用数组：", byte1[1])

	for  i:=0;i<len(byte2);i++{ // for循环数组赋值
		var temp  = "数组item值：" + strconv.Itoa(i)  //类型转换
		byte2[i] = temp
	}

	for  j:=0;j<len(byte2);j++{ // for循环
		fmt.Println("数组item值--->", byte2[j])
	}

}