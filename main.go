package main

import (
	"fmt"
	"math"
	"strconv"
	"unsafe"

	testpkg "./testpkg"
	pkg820 "./pkg820"

)

func traversalString() {
	fmt.Println("--------traversalString test--------")
	s := "pprof.cn博客"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}
func changeString() {
	fmt.Println("--------changeString test--------")
	s1 := "hello"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'H'
	fmt.Println(string(byteS1))

	s2 := "博客"
	runeS2 := []rune(s2)
	runeS2[0] = '狗'
	fmt.Println(string(runeS2))
}
func sqrtDemo() {
	var a, b = 3, 4
	var c int
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}
func arrTest() {
	fmt.Println("--------arr test--------")
	//var arr0 [5]int = [5]int{1, 2, 3} //左边[5]int可省略
	var arr1 = [5]int{1, 2, 3, 4, 5}
	var arr2 = [...]int{1, 2, 3, 4, 5, 6}
	var str = [5]string{3: "hello world", 4: "tom"}
	a := [3]int{1, 2}           // 未初始化元素值为 0。
	b := [...]int{1, 2, 3, 4}   // 通过初始化值确定数组长度。
	c := [5]int{2: 100, 4: 200} // 使用引号初始化元素。
	d := [...]struct {
		name string
		age  uint8
	}{
		{"user1", 10}, // 可省略元素类型。
		{"user2", 20}, // 别忘了最后一行的逗号。
	}
	fmt.Println( arr1, arr2, str)
	fmt.Println(a, b, c, d)
	var arr3 = [...][3]int{{1, 2, 3}, {7, 8, 9}}
	fmt.Println(arr3)
}
func printArr(arr *[5]int) {
	arr[0] = 10
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
func myTest(a [5]int, target int) {
	fmt.Println("--------遍历数组 test--------")
	// 遍历数组
	for i := 0; i < len(a); i++ {
		other := target - a[i]
		// 继续遍历
		for j := i + 1; j < len(a); j++ {
			if a[j] == other {
				fmt.Printf("(%d,%d)\n", i, j)
			}
		}
	}
}
func demo() {
	var (
		a1 string
		b1 int
	)
	fmt.Println(a1 + strconv.Itoa(b1))

	var (
		a2 = "abc"
		b2 = len(a2)
		c2 = unsafe.Sizeof(b2)
	)
	println(a2, b2, c2)

	const (
		a, b = iota + 1, iota + 2 //1,2
		c, d                      //2,3
		e, f                      //3,4
	)
	println(a, b, c, d, e, f)

	s1 := `hello	//多行字符串
		world
		!`
	println(s1)

	arr2 := [...]int{2, 4, 6, 8, 10}
	printArr(&arr2)
	fmt.Println(arr2)
	myTest(arr2, 10)
	/*
		type mystruct struct {
			name string
			m    int
		}
		var my = mystruct{"lxq", 50525}
		fmt.Println(my)
	*/
}

func main() {
	
	fmt.Println("hello world!")
	testpkg.Print()
	//demo()
	//traversalString()
	//changeString()
	//sqrtDemo()
	//arrTest()
	pkg820.Slice()
	pkg820.PointerDemo()
	pkg820.MapDemo()
}
