package main

import "fmt"

func main() {
	var var1 int = 10
	var var2 *int = &var1
	fmt.Println(var1)  //=> 10
	fmt.Println(var2)  //=> 0x10414020
	fmt.Println(*var2) //=> 10
}
