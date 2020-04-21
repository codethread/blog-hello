package main

import "fmt"

func main() {
	fmt.Println(Greet("Blog"))
}

func Greet(entity string) string {
	return fmt.Sprintf("Hello, %s!", entity)
}
