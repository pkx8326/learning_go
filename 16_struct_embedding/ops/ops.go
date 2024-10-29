package ops

import (
	//"bufio"
	"fmt"
	//"os"
	//"strings"
)

type Student struct {
	firstname string
	lastname  string
	email     string
}

type Teacher struct {
	subject        string
	homeroom       string
	teacherAddInfo Student //add type student annonymously
	//(the teacherAddInfo field is the whole student struct)
}

func DataEntrySelector() (c string) {
	//reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Choose action :")
		fmt.Println("1. Edit student data")
		fmt.Println("2. Edit teacher data")
		fmt.Println("3. Display student data")
		fmt.Println("4. Display teacher data")
		fmt.Println("5. Clear student data")
		fmt.Println("6. Clear teacher data")
		fmt.Println("7. Exit")
		fmt.Print("Your choice :")
		fmt.Scanln(&c)
		//c, _ = reader.ReadString('\n')
		//c = strings.TrimSpace(c)
		if c == "1" || c == "2" || c == "3" || c == "4" || c == "5" || c == "6" || c == "7" {
			return
		} else {
			fmt.Println("Invalid input. Please try again.")
			continue
		}
	}
}

func StudentStructConstructor() (newStudent Student) {
	var userFirstName, userLastName, userEmail string

	for {
		fmt.Print("Please input the student's first name :")
		fmt.Scanln(&userFirstName)
		if userFirstName == "" {
			fmt.Println("The student's first name cannot be blank.")
			continue
		} else {
			break
		}
	}

	for {
		fmt.Print("Please input the student's last name :")
		fmt.Scanln(&userLastName)
		if userLastName == "" {
			fmt.Println("The student's last name cannot be blank.")
			continue
		} else {
			break
		}
	}

	for {
		fmt.Print("Please input the student's email :")
		fmt.Scanln(&userEmail)
		if userEmail == "" {
			fmt.Println("The student's emial cannot be blank.")
			continue
		} else {
			break
		}
	}

	newStudent = Student{
		firstname: userFirstName,
		lastname:  userLastName,
		email:     userEmail,
	}

	return
}

func TeacherStructConstructor() (newTeacher Teacher) {
	var userFirstName, userLastName, userEmail, teacherSubject, teacherHomeroom string

	for {
		fmt.Print("Please input the teacher's first name :")
		fmt.Scanln(&userFirstName)
		if userFirstName == "" {
			fmt.Println("The teacher's first name cannot be blank.")
			continue
		} else {
			break
		}
	}

	for {
		fmt.Print("Please input the teacher's last name :")
		fmt.Scanln(&userLastName)
		if userLastName == "" {
			fmt.Println("The teacher's last name cannot be blank.")
			continue
		} else {
			break
		}
	}

	for {
		fmt.Print("Please input the teacher's email :")
		fmt.Scanln(&userEmail)
		if userEmail == "" {
			fmt.Println("The teacher's email cannot be blank.")
			continue
		} else {
			break
		}
	}

	for {
		fmt.Print("Please input the teacher's subject :")
		fmt.Scanln(&teacherSubject)
		if teacherSubject == "" {
			fmt.Println("The teacher's subject cannot be blank.")
			continue
		} else {
			break
		}
	}

	for {
		fmt.Print("Please input the teacher's homeroom :")
		fmt.Scanln(&teacherHomeroom)
		if teacherHomeroom == "" {
			fmt.Println("The teacher's homeroom cannot be blank.")
			continue
		} else {
			break
		}
	}

	newTeacher = Teacher{
		subject:  teacherSubject,
		homeroom: teacherHomeroom,
		teacherAddInfo: Student{ //this is how we can assign values to a nested struct
			firstname: userFirstName,
			lastname:  userLastName,
			email:     userEmail,
		},
	}
	return
}

func DisplayStudentData(studentData *Student) {
	fmt.Println("The student's first name is", studentData.firstname)
	fmt.Println("The student's last name is", studentData.lastname)
	fmt.Println("The student's email is", studentData.email)
}

func DisplayTeacherData(teacherData *Teacher) {
	fmt.Println("The teacher's first name is", teacherData.teacherAddInfo.firstname)
	fmt.Println("The teacher's last name is", teacherData.teacherAddInfo.lastname)
	fmt.Println("The teacher's email is", teacherData.teacherAddInfo.email)
	fmt.Println("The teacher's subject is", teacherData.subject)
	fmt.Println("The teacher's homeroom is", teacherData.homeroom)
}

func ClearStudentData(studentData *Student) {
	fmt.Println("Clearing student data...")
	(*studentData).firstname = ""
	(*studentData).lastname = ""
	(*studentData).email = ""
	fmt.Println("Done.")
}

func ClearTeacherData(teacherData *Teacher) {
	fmt.Println("Clearing teacher data...")
	(*teacherData).teacherAddInfo.firstname = ""
	(*teacherData).teacherAddInfo.lastname = ""
	(*teacherData).teacherAddInfo.email = ""
	(*teacherData).subject = ""
	(*teacherData).homeroom = ""
	fmt.Println("Done.")
}

func Nextoperation() (c string) {
	for {
		fmt.Print("Next operation? [Y/N] :")
		fmt.Scanln(&c)
		if c == "Y" || c == "y" || c == "N" || c == "n" {
			return
		} else {
			fmt.Println("Invalid input. Please try again.")
		}
	}
}
