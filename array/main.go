package main
import "fmt"

func main(){
	var a [3]int
	var b [4]int
	//1、a和b长度不一样，不能赋值
	fmt.Println(a)
	fmt.Println(b)
	//数组的初始化
	//1、数组定义是使用初始值列表的方式初始化
	var cityArray=[4]string{"aa","bb","cc"}
	fmt.Println(cityArray)
	//3、使用索引值方式赋值
	cityArray[2]="dd"
	fmt.Println(cityArray)
	//2、编译器推导数组的长度
	var boolArray=[...]bool{true,false,false,false,true,true,true}
	fmt.Println(boolArray[0])
	fmt.Println(boolArray)
	//使用索引值的方式初始化
	var langArray =[...]string{1:"Golang",3:"Java",7:"Python"}
	fmt.Println(langArray)
	fmt.Printf("%T\n",langArray)

	//数组遍历
	for i:=0;i<len(cityArray);i++{
		fmt.Printf("%s\t",cityArray[i])
	}
	fmt.Println()

	//range遍历数组的下标和值
	for index,value:=range cityArray{
		fmt.Printf("%d:%s\t",index,value)
	}
	fmt.Println()

	//range遍历数组下标
	for index:=range cityArray{
		fmt.Printf("%d\t",index)
	}
	fmt.Println()

	//range遍历数组值
	for _, value:=range cityArray{
		fmt.Printf("%s\t",value)
	}
	fmt.Println()


	//二维数组
	//TArray:=[3][2]string{
	TArray:=[...][2]string{
		{"aa","AA"},
		{"bb","BB"},
	//	{"cc","CC"}}
		{"cc","CC"},
	}
	fmt.Println(TArray)
	fmt.Println(cityArray[2][1])
	fmt.Println()

	for _,row:=range TArray{
		for _,col:=range row{
			fmt.Printf("%s\t",col)
		}
		fmt.Println()
	}

	for _,row:=range TArray{
		fmt.Println(row)
	}

	
	//数组是值类型，就是赋值的时候是完整拷贝
	//和java和C++不同
	x:=[3]int{1,2,3}
	fmt.Println("func1前")
	fmt.Println(x)
	fmt.Println("func1后")
	f1(x)
	fmt.Println(x)

}
//鲁迅曾经说过，the best way to code，
func f1(a [3]int){
	a[0]=100
	fmt.Println(a[0])
}

//数组长度是固定的，数组的长度属于类型部分，也就是不同长度的数组不能赋值
//切片底层是数组，是一个拥有相同类型元素的可变长度的序列，
//是引用类型，内部结构包含指向底层数组的指针，长度，容量