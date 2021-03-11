package main 
import "fmt"
func main(){
	a:=10
	b:=46
	fmt.Println(a+b)
	fmt.Println(a-b)
	fmt.Println(a/b)
	fmt.Println(a%b)

	//go只有a++ a--
	//就是+1，-1操作
	//是一个语句，++，--不是运算符
	a++
	fmt.Println(a)
	fmt.Println()

	fmt.Println(a>b)
	fmt.Println(a<b)

	//这句是错的，不能强转
	//fmt.Println(a&&b)

	fmt.Println()
	fmt.Println(10>5 && 3>6)

	//位运算符
	//& 两数对应位想与
	//| 两数对应位相或
	fmt.Println()
	c:=12
	d:=14
	//e:=c&d
	fmt.Printf("c:%d\t",c)
	fmt.Printf("d:%d\n",d)
	fmt.Printf("c:%b\t",c)
	fmt.Printf("d:%b\n",d)
	fmt.Printf("c&d:%d\t",c & d)
	fmt.Printf("c&d:%b\n",c & d)
	fmt.Printf("c|d:%d\t",c | d)
	fmt.Printf("c|d:%b\n",c | d)
	fmt.Printf("1<<2:%d\t",1<<2)
	fmt.Println()

	//赋值运算
	a+=a
	fmt.Println(a)
	a>>=1
	fmt.Println(a)
	
}
