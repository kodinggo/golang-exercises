package main

import (
	"fmt"

	"golang-second-project/subject"
	"golang-second-project/user"
)

func main() {
	math := subject.NewMath("Matematika")
	english := subject.NewEnglish("B. Inggris")

	student := user.NewStudent(math, english)
	student.SetName("Doni")

	fmt.Println(student.GetName())
	fmt.Println(student.GetInfo())
}
