package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	foo()
	fmt.Println("after foo")
	for i:=0; i< 10000000; i++ {
		if i %2==0 {
			fmt.Println(i)
		}
	}
	bar()
}

func foo(){
	fmt.Println("I am in foo")
}

func bar(){
	fmt.Println("end... bye")
}
