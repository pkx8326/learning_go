package main

import (
	"fmt"

	"example.com/struct-embedding/ops"
)

func main() {
	var newStudent ops.Student
	var newTeacher ops.Teacher
	var nextopt string

	for {
		c := ops.DataEntrySelector()
		if c == "1" {
			newStudent = ops.StudentStructConstructor()
			nextopt = ops.Nextoperation()
		} else if c == "2" {
			newTeacher = ops.TeacherStructConstructor()
			nextopt = ops.Nextoperation()
		} else if c == "3" {
			ops.DisplayStudentData(&newStudent)
			nextopt = ops.Nextoperation()
		} else if c == "4" {
			ops.DisplayTeacherData(&newTeacher)
			nextopt = ops.Nextoperation()
		} else if c == "5" {
			ops.ClearStudentData(&newStudent)
			nextopt = ops.Nextoperation()
		} else if c == "6" {
			ops.ClearTeacherData(&newTeacher)
			nextopt = ops.Nextoperation()
		} else {
			nextopt = "n"
		}
		if nextopt == "N" || nextopt == "n" {
			fmt.Println("Goodbye.")
			break
		} else {
			continue
		}
	}
}
