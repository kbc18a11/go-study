package main

import "fmt"

type Person struct {
	firstName string
	age       int
}

func (a Person) name() string { //Personのメソッド
	return a.firstName + " Personのやつ"
}

func (p Person) getAge() int {
	return p.age
}

type User struct {
	Person
}

func (a User) name() string { //Userのメソッド
	return a.firstName + " Userのやつ"
}

func main() {
	bob := Person{"Bob", 10}
	mike := User{}
	mike.firstName = "Mike"
	mike.age = 20

	fmt.Println(bob.name())    //=> Bob
	fmt.Println(mike.name())   //=> Mike
	fmt.Println(mike.getAge()) //=> Mike
}
