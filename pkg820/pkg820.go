package pkg820

import "fmt"

//Slice is a demo of slice
func Slice() {
	fmt.Println("--------slice test--------")
/*
	var arr = [...]int{3, 1, 4, 1, 5, 9}
	slice := arr[3:4]
	fmt.Println("arr=", arr)
	fmt.Println("slice 内容=", slice)
	fmt.Println("slice nums=", len(slice))
	fmt.Println("slice cap", cap(slice))
	fmt.Println("arr地址=", &arr[2])
	fmt.Println("slice地址=", &slice[0])
*/
/*
	slice2 := make([]int, 2)
	fmt.Println("slice2 地址=",&slice2[0])
	fmt.Println("slice2 内容=", slice2)
	fmt.Println("slice2 nums=", len(slice2))
	fmt.Println("slice2 cap", cap(slice2))
	slice2 = append(slice2,3)
	fmt.Println("slice2 地址=",&slice2[0])
	fmt.Println("slice2 内容=", slice2)
	fmt.Println("slice2 nums=", len(slice2))
	fmt.Println("slice2 cap", cap(slice2))
	slice2 = append(slice2,slice2...)
	fmt.Println("slice2 地址=",&slice2[0])
	fmt.Println("slice2 内容=", slice2)
	fmt.Println("slice2 nums=", len(slice2))
	fmt.Println("slice2 cap", cap(slice2))
*/
	fmt.Println("cap分配规律：")
	s := make([]int, 0, 1)
    c := cap(s)

    for i := 0; i < 50; i++ {
        s = append(s, i)
        if n := cap(s); n > c {
            fmt.Printf("cap: %d -> %d\n", c, n)
            c = n
        }
	}
	fmt.Println("slice cpoy:")
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    fmt.Println("array data : ", data)
    s1 := data[7:]
    s2 := data[2:6]
    fmt.Printf("slice s1 : %v\n", s1)
    fmt.Printf("slice s2 : %v\n", s2)
    copy(s1, s2)
    fmt.Printf("copied slice s1 : %v\n", s1)
    fmt.Printf("copied slice s2 : %v\n", s2)
	fmt.Println("last array data : ", data)
	
	str := "Hello world"
    st := []rune(str) //中文字符需要用[]rune(str)
    st[6] = 'G'
    st = st[:8]
    st = append(st, '!')
    str = string(st)
    fmt.Println(str)
}

//PointerDemo is a demo of pointer
func PointerDemo(){
	fmt.Println("--------pointer test--------")
	a := 10
    b := &a // 取变量a的地址，将指针保存到b中
    fmt.Printf("type of b:%T\n", b)
    c := *b // 指针取值（根据指针去内存取值）
    fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)
	
	var a1 *int
	a1 = new(int) //不初始分配内存会报错
    *a1 = 100
    fmt.Println(*a1)

}
//MapDemo is a demo of map
func MapDemo(){//key-value
	fmt.Println("--------map test--------")
	var m map[string]int
	m = make(map[string]int,5)//make 给map分配空间
    m["测试"] = 100
    fmt.Println("m=",m)

	//方法二
	m2 := make(map[string]string)
	m2["city1"] = "北京"
	m2["c2"] = "天津"
	delete(m2,"c2")
	fmt.Println("m2=",m2)
	//方法三
	m3 := map[string]string{
		"师傅" : "zqr",
		"徒弟" : "lxq",
		"课时" : "100years",
		"学费" : "-999999",
	}
	for k,v := range m3{
        fmt.Println(k, v)
	}
	
	
	mm := make(map[string]map[string]string)

	mm["stu1"] = make(map[string]string)
	mm["stu1"]["name"] = "tom"
	mm["stu1"]["sex"] = "man"
	mm["stu2"] = make(map[string]string)
	mm["stu2"]["name"] = "jerry"
	mm["stu2"]["sex"] = "woman"
	//fmt.Println("值为map的map\nmm=",mm)
	for k1,v1 := range mm{
		fmt.Printf("k1=%v",k1)
		for k2,v2 := range v1{
			fmt.Printf("\t k2=%v,v2=%v\n",k2,v2)
		}
	}
	//值为slice的map
	ms := make(map[string][]int,3)
	fmt.Println("值为slice的map\nms=",ms)
	ms["arr1"] = make([]int,0)
	ms["arr1"] = append(ms["arr1"],0,1,2,3,4)
	fmt.Println("值为slice的map\nms=",ms)

	//值为map的slice
	sm := make([]map[string]string,0)
	hero := map[string]string{
		"name" : "孙悟空",
		"age" : "500",
		"skill" : "72",
	} 
	sm = append(sm,hero)
	hero = map[string]string{
		"name" : "猪八戒",
		"age" : "1000",
		"skill" : "36",
	} 
	sm = append(sm,hero)
	// if sm[0] ==nil{
	// 	sm[0] =make(map[string]string)
	// 	sm[0]["name"] = "孙悟空"
	// 	sm[0]["age"] = "500"
	// 	sm[0]["skill"] = "72"
	// }
	// if sm[1] ==nil{
	// 	sm[1] =make(map[string]string)
	// 	sm[1]["name"] = "猪八戒"
	// 	sm[1]["age"] = "1000"
	// 	sm[1]["skill"] = "36"
	// }
	fmt.Println("值为map的slice\n",sm)
	for i := range sm{
		fmt.Printf("hero%v:\n",i)
		for k,v := range sm[i]{
			fmt.Printf("k=%v,v=%v\n",k,v)
		}

	}
}
