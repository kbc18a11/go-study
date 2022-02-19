package main

import (
	"fmt"
	"time"
)

type Person struct{}  //Person構造体
type Person2 struct{} //Person2構造体

type People interface {
	intro()
}

func IntroForPerson(arg People) {
	arg.intro()
}

//Person構造体のメソッドintro()
func (p *Person) intro() {
	fmt.Println("Hello World")
}

//Person2構造体のメソッドintro()
func (p *Person2) intro() {
	fmt.Println("Hello World")
}

func main() {
	ch1 := make(chan bool)
	ch2 := make(chan bool)

	fmt.Println("Start!")

	go func() {
		process(2, "A")
		ch1 <- true
	}()

	go func() {
		process(2, "B")
		ch2 <- true
	}()

	<-ch1
	<-ch2

	fmt.Println("Finish!")
}

func process(num int, str string) {
	for i := 0; i <= num; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(i, str)
	}
}
