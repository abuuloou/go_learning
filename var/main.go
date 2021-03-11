package main
import "fmt"

//全局变量
var x=88

func foo()(int,string){
	return 10,"岳云鹏"
}

func main(){
	var name string
	var age int
	var isOk bool

	fmt.Println(name,age,isOk)

	var(
		a string
		b int
		c bool
		d float32
	)
	//声明变量的时候会自动初始化
	fmt.Println(a,b,c,d)

	//声明变量同时初始值
	var name1 string ="郭德纲"
	var age1 int =55
	fmt.Println(name1,age1)
	
	var(
		A string="aa"
		B string="bb"
	)
	fmt.Println(A,B)

	//类型推导：根据赋值的初始值推断出var的类型
	var name2,age2 ="于谦",50
	fmt.Println(name2,age2)

	//短变量声明：函数内部，可以使用  :=
	short:=100
	fmt.Println(short)

	fmt.Println(x)

	//匿名变量，匿名变量用_表示
	//匿名变量不分配命名空间，也不会分配内存，所以匿名变量之间不存在重复声明
	//比如一个函数有两个返回值，我只想要其中一个，另一个就可以用匿名变量去接收
	x,_ :=foo()
	_,y :=foo()
	fmt.Println("x=",x)
	fmt.Println("y =",y)

}