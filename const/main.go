package main
import "fmt"
//常量
// const PI=3.14
// const e=2.7

// const(
// 	PI=3.14
// 	e=2.7
// )

const(
	n1=10
	//如果常量声明某一行没有表达式
	//他默认和上一行一样
	//n2=10
	n2
	n3
)

const (
	//iota遇到const被重置为0
	a1=iota
	//const中每新增一行常量声明将使iota计数一次
	a2
	a3
	a4
)

//_跳过某些值
const (
	b1=iota //0 
	b2  //1
	_
	b4  //3
)

//iota声明中间插队
const (
	//iota遇到const被重置为0
	c1=iota //0 
	c2  //1
	c3=100
	c4=iota  //3
)
const c5=iota

//const定义数量级
const(
	_=iota
	KB=1<<(10*iota)//1<<10
	MB=1<<(10*iota)//1<<20
	GB=1<<(10*iota)//1<<30
	TB=1<<(10*iota)//1<<40
	PB=1<<(10*iota)//1<<50
)

//多个iota定义在一行
const(
	a,b=iota+1,iota+2//iota=0,a=1,2
	c,d//iota=1,c=2,d=3
	e,f//iota=2,e=3,f=4
)

//iota是go语言的常量计数器，只能在常量的表达式中使用
//iota在const关键字出现时将被重置为0.
//const中每新增一行常量声明将使iota计数一次（iota可理解为const语句块中的行索引）
//使用iota能简化定义，在定义枚举时很有用
func main(){
	fmt.Println(n1,n2,n3)
	const haha=15
	fmt.Println(haha)

	fmt.Println(a1,a2,a3,a4)

	fmt.Println(b1,b2,b4)

	fmt.Println(c1,c2,c3,c4,c5)

	fmt.Println(KB,MB,GB,TB,PB)

	fmt.Println(a,b,c,d,e,f)
}