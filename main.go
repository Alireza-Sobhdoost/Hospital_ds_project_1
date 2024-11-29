package main

import (
	"fmt"
	"log"
	"project_1/DataStructures" 
	"project_1/Entities"
	"project_1/Auth"
	"bufio"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"reflect"


)


func greet()(int) {
	fmt.Println("==Hospital==")
	fmt.Println("Wellcome to our Hospital. How can we help you?")
	fmt.Println("[1] Sign up")
	fmt.Println("[2] Login")
	fmt.Println("[3] Exit")

	reader := bufio.NewReader(os.Stdin)
	cmd, _ := reader.ReadString('\n')
	cmd = cmd[:len(cmd)-1] // Remove the trailing newline character
	Intcmd , _:= strconv.Atoi(cmd)
	return Intcmd

}
func signup_form() (string, string, string, string, []string, int) {
	fmt.Println("==Signup==")
	fmt.Println("Please enter your information in order below:")
	fmt.Println("First name, Last name, National ID, password, age, role")
	fmt.Println("If you're a doctor, please enter your department too")

	reader := bufio.NewReader(os.Stdin)

	// Read user inputs
	fmt.Print("First name: ")
	firstName, _ := reader.ReadString('\n')
	firstName = firstName[:len(firstName)-1] // Remove the trailing newline character

	fmt.Print("Last name: ")
	lastName, _ := reader.ReadString('\n')
	lastName = lastName[:len(lastName)-1]

	fmt.Print("National ID: ")
	nationalID, _ := reader.ReadString('\n')
	nationalID = nationalID[:len(nationalID)-1]

	fmt.Print("Password: ")
	password, _ := reader.ReadString('\n')
	password = password[:len(password)-1]

	fmt.Print("Age: ")
	ageStr, _ := reader.ReadString('\n')
	ageStr = ageStr[:len(ageStr)-1]
	age, err := strconv.Atoi(ageStr) // Convert age to int
	if err != nil {
		fmt.Println("Error: Invalid age entered.")
		return "", "", "", "", nil, 0
	}

	fmt.Print("Role: ")
	role, _ := reader.ReadString('\n')
	role = role[:len(role)-1]

	// If the user is a doctor, ask for the department
	var department string
	if role == "Doctor" {
		fmt.Print("Department: ")
		department, _ = reader.ReadString('\n')
		department = department[:len(department)-1]
	}

	// // Collect the information
	// fmt.Println("\nCollected Information:")
	// fmt.Println("First Name:", firstName)
	// fmt.Println("Last Name:", lastName)
	// fmt.Println("National ID:", nationalID)
	// fmt.Println("Password:", password)
	// fmt.Println("Age:", age)
	// fmt.Println("Role:", role)

	if role == "Doctor" {
		fmt.Println("Department:", department)
		args := []string{role, department}
		return nationalID, firstName, lastName, password, args, age
	} else {
		args := []string{role}
		return nationalID, firstName, lastName, password, args, age
	}
}


func login_form() (string, string) {
	fmt.Println("==Login==")
	fmt.Println("Please enter your information in order below:")
	fmt.Println("National ID, password")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("National ID: ")
	nationalID, _ := reader.ReadString('\n')
	nationalID = nationalID[:len(nationalID)-1]

	fmt.Print("Password: ")
	password, _ := reader.ReadString('\n')
	password = password[:len(password)-1]

	return nationalID, password

}

func clear() {
	var cmd *exec.Cmd

	// Check the operating system and run the appropriate command
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}

	// Execute the command
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func main() {
	
	DataBase := DataStructures.NewHashMap(100)
	DoctorsDB := DataStructures.NewHashMap(100)
	PatientsDB := DataStructures.NewHashMap(100)
	ManagerDB  := DataStructures.NewHashMap(100)
	DataBase.Insert("Doctors" , DoctorsDB)
	DataBase.Insert("Patients", PatientsDB)
	DataBase.Insert("Managers" , ManagerDB)
	
	cmd := greet()
	clear()
	for true {
		
		if cmd == 1 {

			NID , FirstName , Lastname , password , args , age := signup_form()
			err := Auth.Signup( NID , FirstName , Lastname , password , args , age,*DataBase)
			clear()
			if err != nil {
				log.Fatalf("Error setting password for doctor: %v", err)
			}
			
		} else if cmd == 2 {

			NID , password := login_form()
			user , err:= Auth.Login(*DataBase, NID, password)
			clear()
			if err != nil {
				log.Fatalf("Error setting password for doctor: %v", err)
			} else {
				our_type := reflect.TypeOf(user)
				if our_type == reflect.TypeOf(&Entities.Doctor{}) {
					currentUser := user.(*Entities.Doctor)
					fmt.Printf("Doctor: %v %v, Department: %v\n", currentUser.FirstName, currentUser.LastName, currentUser.Department)

				} else if our_type == reflect.TypeOf(&Entities.Patient{}) {
					currentUser := user.(*Entities.Patient)
					fmt.Printf("Doctor: %v %v, Department: %v\n", currentUser.FirstName, currentUser.LastName)

				} else if our_type == reflect.TypeOf(&Entities.Manager{}) {
					currentUser := user.(*Entities.Manager)
					fmt.Printf("Doctor: %v %v, Department: %v\n", currentUser.FirstName, currentUser.LastName, currentUser.Department)

				}

				break
			}	
		} else if cmd == 3 {

			break
		}
		cmd = greet()
		clear()
	}
	


	// doctor := &Entities.Doctor{
	// 	User: Entities.User{
	// 		ID:        "1",
	// 		FirstName: "Alice",
	// 		LastName:  "Smith",
	// 	},
	// 	Department: "Cardiology",
	// 	PatientList: DataStructures.LinkedList{},
	// 	VisitQueue: DataStructures.NewPriorityQueue(func(a, b interface{}) bool {
	// 		patientA := a.(Entities.Patient)
	// 		patientB := b.(Entities.Patient)
	// 		return patientA.PriorityToVsit < patientB.PriorityToVsit
	// 	}),
	// }
	// args := []string{"Doctor", "Cardiology"}
	// err := Auth.Signup("1", "Alice", "Smith", "password", args ,  20 , *DataBase)
	
	// if err != nil {
	// 	log.Fatalf("Error setting password for doctor: %v", err)
	// }


	// doc , _:= Auth.Login(*DataBase, "1", "password")
	// doctor := doc.(*Entities.Doctor)
	// fmt.Printf("Doctor: %v %v, Department: %v\n", doctor.FirstName, doctor.LastName, doctor.Department)


	// doctor.VisitQueue.Push(Entities.Patient{User: Entities.User{
	// 	ID:        "10",
	// 	FirstName: "John",
	// 	LastName:  "Doe",
	// }, PriorityToVsit: 3})

	// doctor.VisitQueue.Push(Entities.Patient{User: Entities.User{
	// 	ID:        "20",
	// 	FirstName: "Jane",
	// 	LastName:  "Smith",
	// }, PriorityToVsit: 1})

	// p1 := &Entities.Patient{User: Entities.User{
	// 	ID:        "30",
	// 	FirstName: "Emily",
	// 	LastName:  "Davis",
	// }, PriorityToVsit: 2}
	// p1.MedicalHistory = "High blood pressure"
	// fmt.Printf("Patient: %v %v, Priority: %v, Medical History: %v\n", p1.FirstName, p1.LastName, p1.PriorityToVsit, p1.MedicalHistory)

	// doctor.VisitQueue.Push(*p1)
	
	// fmt.Println("Patients in visit order:")

	// for !doctor.VisitQueue.IsEmpty() {
	// 	patient, err := doctor.VisitQueue.Pop()
	// 	Patient := patient.(Entities.Patient)
	// 	if err != nil {
	// 		fmt.Println("Error:", err)
	// 		break
	// 	}
	// 	fmt.Printf("%v\n", Patient.FirstName)
	// }

	// doc2 , _:= Auth.Login(*DataBase, "1", "password")
	// doctor2 := doc2.(*Entities.Doctor)

	// for !doctor2.VisitQueue.IsEmpty() {
	// 	patient, err := doctor2.VisitQueue.Pop()
	// 	Patient := patient.(Entities.Patient)
	// 	if err != nil {
	// 		fmt.Println("Error:", err)
	// 		break
	// 	}
	// 	fmt.Printf("%v\n", Patient.FirstName)
	// }



}