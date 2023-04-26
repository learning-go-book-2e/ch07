package main

import (
	"fmt"
)

func main() {
	var s *string
	fmt.Println(s == nil) // prints true
	var i interface{}
	fmt.Println(i == nil) // prints true
	i = s
	fmt.Println(i == nil) // prints false
}
