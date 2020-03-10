package main

import "fmt"

func main()  {
	s1 := [] int {1,2,3 }
	s2 := [5] int {1,2,3 }
	fmt.Println(s1) // [1 2 3]
	fmt.Println(s2) // [1 2 3 0 0]
}