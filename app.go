package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func init() {
	fmt.Println("Init now...")

}

type Phone interface {
	call()
}
type IPhone struct {
}
type OPhone struct {
}

func (iPhone IPhone) call() {
	println("iPhone call")
}
func (iPhone OPhone) call() {
	println("OPhone call")
}

func newer() {
	// Init Log
	log.SetPrefix("App:")
	log.SetFlags(0)

	fmt.Println("Start build ...")
	msg := builder.Build("test msg")
	fmt.Println(msg)
	msg, err := builder.BuildWithError("a")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(msg)
	}
	fmt.Println("End build ...")

	var x, _, y = "xxx", "___", "yyy"
	println(x)
	println(y)

	const Male = "Male"
	println(Male)
	ptr := &x
	println(ptr)
	println(*ptr)

	const (
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		i        //8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)

	var add = func(a int, b int) int {
		return a + b
	}

	println(add(1, 3))
	/*
		type Circle struct {
			radius int
		}
		func(c Circle) getArea() int {
			return c.radius*c.radius
		}
	*/
	var o = [5]int{1, 2, 3}
	var p = []int{1, 2, 3}
	p[2] = 1
	var j int
	for j = 0; j < 3; j++ {
		println(o[j])
	}
	type Books struct {
		title string
	}

	var book Books
	var ptrBook *Books
	book.title = "123"
	ptrBook = &book
	println(book.title)
	println(ptrBook.title)

	var s = []int{1, 2, 3, 4, 5}
	var s1 = s[:]
	var s2 = s[0:]
	var s3 = s[:5]
	var s4 = s[1:2]
	fmt.Printf("%v", s)
	fmt.Printf("%v", s1)
	fmt.Printf("%v", s2)
	fmt.Printf("%v", s3)
	fmt.Printf("%v", s4)
	for i, v := range s {
		println(i, v)
	}
	var m map[string]string = make(map[string]string)
	m["k"] = "value"
	m["k2"] = "value2"
	for i, v := range m {
		println(i, v)
	}
	delete(m, "k")
	for k, v := range m {
		println(k, v)
	}
	// i = int("1")
	// println(a)
	var phone Phone
	phone = new(IPhone)
	phone.call()
	phone = new(OPhone)
	phone.call()
}

//WriteFile test to write a file
func WriteFile() {
	var text = "text content"
	var data = []byte(text)
	ioutil.WriteFile("tmp", data, 0600)
}

//ReadFile to read a file
func ReadFile() {
	body, _ := ioutil.ReadFile("tmp")
	var text = string(body)
	println(text)
}
func main() {
	println("Start wirte file ....")
	WriteFile()
	println("Start read file ....")
	ReadFile()

}
