package main

import "fmt"

type Student struct {
	name string
	age int
}

func main() {
	var stu1 Student
	stu1.name = "apple"
	stu1.age = 20
	fmt.Println(stu1)  // {apple 20}
	fmt.Println(stu1.name)  // apple
	fmt.Println("-------------")

	var stu2 = new(Student)
	stu2.name = "apple"
	stu2.age = 20
	fmt.Println(stu2)  // &{apple 20}
	fmt.Println(stu2.name) // apple
	fmt.Println((*stu2).name) // apple
	fmt.Println("-------------")

	var stu3 = Student{"apple", 20}
	fmt.Println(stu3)  // {apple 20}
	fmt.Println(stu3.name) //apple
	fmt.Println("-------------")

	var stu4 = &Student{"apple", 20}
	fmt.Println(stu4)  // &{apple 20}
	fmt.Println(stu4.name) // apple
	fmt.Println((*stu4).name) // apple
	fmt.Println("-------------")

	stu5 := new(Student)
	stu5.name = "apple"
	stu5.age = 20
	fmt.Println(stu5)  // &{apple 20}
	fmt.Println(stu5.name) // apple
	fmt.Println((*stu5).name) // apple
	fmt.Println("-------------")
}