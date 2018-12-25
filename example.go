package main

import "fmt"

type Describer interface {
	Describe()
}
type Person struct {
	name string
	age  int
}

func (p Person) Describe() { //implemented using value receiver
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

type Address struct {
	state   string
	country string
}

func (a *Address) Describe() { //implemented using pointer receiver
	fmt.Printf("State %s Country %s", a.state, a.country)
}

func main() {
	var d1 Describer // 接口类型变量
	p1 := Person{"Sam", 25}
	d1 = p1 // 值类型
	d1.Describe()

	p2 := Person{"James", 32}
	d1 = &p2 // 指针类型
	d1.Describe()

	var d2 Describer
	a := Address{"Washington", "USA"}

	//d2 = a // 不能使用值类型（引发panic）①
	//
	d2 = &a
	d2.Describe()
	a.Describe() // 直接使用值类型调用②

	a1 := []int{1, 2, 3}
	fmt.Println(a1[:10])

}
