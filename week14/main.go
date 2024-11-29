package main

import (
	"fmt"
)

type student struct {
	id   int
	name string
	gpa  float32
}

func main() {
	var student1 student
	student1.id = 202444024
	student1.name = "주영훈"
	student1.gpa = 4.5
	fmt.Println(student1.gpa)
	var student2 student
	student2.id = 202444000
	student2.name = "joo"
	student2.gpa = 4.4
	fmt.Println(student2.id)
}
