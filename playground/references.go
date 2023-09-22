package main

import "fmt"

func ChangingString(myString *string) {
	*myString = "changed!"
}

func main() {

	someString := "hello sir"

	ChangingString(&someString)

	fmt.Println(someString)

}
