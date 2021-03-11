package main

//数组长度是固定的，数组的长度属于类型部分，也就是不同长度的数组不能赋值
//切片底层是数组，是一个拥有相同类型元素的可变长度的序列，
//是引用类型，内部结构包含指向底层数组的指针，长度，容量
import (
	"fmt"	
	"sort"
)

func main(){
	var a []string
	var b []int
	c := []bool{true,false,true}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println()

	// //切片的拷贝，因为切片是引用类型，如果直接赋值，会指向同一个元素
	// //copy相当于深拷贝

	// var d []bool
	// copy(d,c)
	
	// fmt.Printf("c的地址：%p\n",c)
	// fmt.Printf("d的地址：%p\n",d)
	// e:=d
	// fmt.Printf("e的地址：%p\n",e)

	//construct()
	//testCapAndLen()
	//compare()

	// set()
	// foreach()
	//growSize()
	del()


	//使用内置sort对数组进行排序
	//sort.Ints需要一个切片类型，所以直接切ss
	//ss是指向arr的引用，所以就直接改变ss了
	var ss =[...]int{3,7,8,8,1}
	sort.Ints(ss[:])
	fmt.Println(ss)

}
func del(){
	a:=[]string{"aaa","bbb","ccc","ddd"}
	a=append(a[0:2],a[3:]...)
	fmt.Println(a)
}

func growSize(){
	//切片要初始化之后才能使用
	//这是一个nil切片，底层没有数组
	var a []int
	//编译不会出错，但运行会报错
	//panic: runtime error: index out of range [0] with length 0
	// a[0]=100
	// fmt.Println(a)

	// a=append(a, 10)
	// //[10]
	// fmt.Println(a)

	for i:=0;i<10;i++{
		a=append(a, i)
		//a的值 长度 容量 地址
		//切片是引用类型所以不用取址符
		fmt.Printf("%v \t\t\tlen:%d \t\t\tcap%d \t\t\tptr:%p\n",a,len(a),cap(a),a)
	}

	//append函数将数组元素追加到数组的最后并返回该数组
	//append可以一次追加多个元素
	b:=[]int{11,12,13,14}
	a=append(a,b...)
}

func foreach(){
	a:=[]int{1,2,3,4,5,6,7,8,9}
	for i:=0;i<len(a);i++{
		fmt.Printf("%d\t",a[i])
	}
	fmt.Println()

	for _,value:=range a{
		fmt.Printf("%d\t",value)
	}
}
func testCapAndLen(){
	//如果切片没有初始值，cap=size=0
	var a []string
	//初始化之后，cap=size=len
	c := []bool{true,false,true}
	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Println()
	fmt.Println(len(c))
	fmt.Println(cap(c))
}

func construct(){
	//基于数组定义切片
	d := [5]int{1,2,3,4,5}
	//左闭右开
	e := d[1:4]
	fmt.Println(e)
	fmt.Printf("%T\n",e)
	fmt.Println()

	//切片再次切片
	//两个语句等价
	f :=e[:]
	//f:=e[0:len(e)]
	fmt.Println(f)
	fmt.Println()

	//make函数构造切片
	//arg1:切片类型   arg2:切片里数据的个数  arg3:切片指向底层数组的容量
	//arg3可以不写，不写默认capacity=size
	g:=make([]int,5,10)
	fmt.Println(g)
	fmt.Printf("%T\t",g)
	fmt.Println()

	//通过len函数获取切片长度
	fmt.Println(len(g))
	//通过cap函数获取切片的len
	fmt.Println(cap(g))
}

func compare(){
	//切片不能用==进行比较
	//切片只能和 nil比较
	//一个nil值的切片并没有底层数组，一个nil的切片长度和容量都是0
	//但一个长度容量为0的切片一定是nil

	//nil
	var s1 []int
	fmt.Println(s1,len(s1),cap(s1))
	if s1==nil{
		fmt.Println("s1是一个nil切片，底层是没有数组的")
	}

	var s2 =[]int{}
	fmt.Println(s2,len(s2),cap(s2))
	if s2==nil {
		fmt.Println("s2是一个nil切片，底层是没有数组的")
	}else{
		fmt.Println("s2不是一个nil切片，底层是有数组的")
	}

	c:=make([]int,0)
	fmt.Println(c,len(c),cap(c))
	if c==nil {
		fmt.Println("c是一个nil切片，底层是没有数组的")
	}else{
		fmt.Println("c不是一个nil切片，底层是有数组的")
	}

}

func set(){
	//切片是赋值拷贝
	//a b是共用一个底层数组的
	a:=make([]int,3)
	b:=a
	b[0]=100
	fmt.Println(a[0])
	fmt.Println(b[0])

}