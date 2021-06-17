package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score int
	left  *Student
	right *Student
}

func trans(root *Student) {
	if root == nil {
		return
	}
	fmt.Println(root)

	trans(root.left)
	trans(root.right)

}

func main() {
	var root *Student = new(Student)
	root.Name = "Stu01"
	root.Age = 1
	root.Score = 11

	var left *Student = new(Student)
	left.Name = "Stu02"
	left.Age = 12
	left.Score = 112

	root.left = left

	var right *Student = new(Student)
	right.Name = "Stu03"
	right.Age = 13
	right.Score = 113

	root.right = right

	var left1 *Student = new(Student)
	left1.Name = "Stu04"
	left1.Age = 14
	left1.Score = 114

	left.left = left1

	var right2 *Student = new(Student)
	right2.Name = "Stu05"
	right2.Age = 15
	right2.Score = 115

	left1.right = right2

	trans(root)

}
