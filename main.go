package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"project_1/Auth"
	"project_1/DataStructures"
	"project_1/Entities"
	"reflect"
	"runtime"
	"strconv"
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

func Patient_menu()(int) {
	fmt.Println("==Patient Menu==")
	fmt.Println("Please enter your choice:")
	fmt.Println("[1] Book an appointment")
	fmt.Println("[2] Cancel appointments")
	fmt.Println("[3] Edit account")
	fmt.Println("[4] Exit")

	reader := bufio.NewReader(os.Stdin)
	cmd, _ := reader.ReadString('\n')
	cmd = cmd[:len(cmd)-1] // Remove the trailing newline character
	Intcmd , _:= strconv.Atoi(cmd)
	clear()
	return Intcmd
}

func Book_appointment()(string) {
	fmt.Println("==Patient Menu==")
	fmt.Println("Please enter your choice:")
	fmt.Println("[1] Cardiology")
	fmt.Println("[2] Emergency")
	fmt.Println("[3] Back")

	reader := bufio.NewReader(os.Stdin)
	cmd, _ := reader.ReadString('\n')
	cmd = cmd[:len(cmd)-1] // Remove the trailing newline character
	Intcmd , _:= strconv.Atoi(cmd)
	clear()

	if Intcmd == 1 {
		return "Cardiology"
	} else if Intcmd == 2 {
		return "Emergency"
	} else {
		return "Back"
	}
}

func choose_doc(caller Entities.Patient , DB *DataStructures.HashMap) () {
	fmt.Println("==Choose a doctor==\n")
	DocsList , lenght := Entities.DisplayDocs(DB)
	DocsList.Display()
	fmt.Println("[e] back")
	reader := bufio.NewReader(os.Stdin)
	cmd, _ := reader.ReadString('\n')
	cmd = cmd[:len(cmd)-1] // Remove the trailing newline character
	if cmd == "e" {
		return 
	}
	Intcmd , _:= strconv.Atoi(cmd)
	clear()
	doc_internal_pointer_var := DocsList.Find_by_index(Intcmd-1 , lenght)
	// fmt.Println(doc_internal_pointer_var)
	doc_internal_pointer_var.Data.(*Entities.Doctor).VisitQueue.Push(caller)
	caller.DoctorList.AddToStart(doc_internal_pointer_var.Data.(*Entities.Doctor))
	caller.DoctorList.Display()
	fmt.Println("You have been added to the queue")
	
}

func cancel_appointment(caller Entities.Patient) () {
	fmt.Println("==Choose a appointment==\n")
	caller.DoctorList.Display()
	
	lenght := Entities.DisplayDocsList(*caller.DoctorList)
	fmt.Println("[e] back")

	reader := bufio.NewReader(os.Stdin)
	cmd, _ := reader.ReadString('\n')
	cmd = cmd[:len(cmd)-1] // Remove the trailing newline character
	if cmd == "e" {
		return 
	}
	Intcmd , _:= strconv.Atoi(cmd)
	clear()
	doc_internal_pointer_var := caller.DoctorList.Find_by_index(Intcmd-1 , lenght)
	// fmt.Println(doc_internal_pointer_var)
	doc_internal_pointer_var.Data.(*Entities.Doctor).VisitQueue.Remove(caller)
	caller.DoctorList.Remove(doc_internal_pointer_var.Data.(*Entities.Doctor))
	fmt.Println("You have been added to the queue")
	
}

func Doctor_menu(doc Entities.Doctor)(int) {


	fmt.Println("==Doctor Menu==")
	fmt.Println("Hello Dr. " , doc.FirstName , " " , doc.LastName)
	fmt.Println("----------------to visit list----------------")
	fmt.Println("Patient	Firstname	Lastname	Age		ID")
	count := 0
	for i, value := range doc.VisitQueue.Heap {
		patient := value.(Entities.Patient)
		fmt.Printf("[%d] %s %s %d %s\n", i+1, patient.FirstName, patient.LastName, patient.Age, patient.ID) // Use %v to handle generic types
		count += 1
	}
	if count == 0 {
		fmt.Println("There is no one to visit")
	} else {
		fmt.Println("There are " , count , " patients to visit")
	}

	fmt.Println("---------------------------------------------")

	fmt.Println("[1] Start to visit")
	fmt.Println("[2] Edit account")
	fmt.Println("[3] Exit")

	reader := bufio.NewReader(os.Stdin)
	cmd, _ := reader.ReadString('\n')
	cmd = cmd[:len(cmd)-1] // Remove the trailing newline character
	Intcmd , _:= strconv.Atoi(cmd)
	return Intcmd
	
}

func visit(patient Entities.Patient)() {
	fmt.Println("==Visit==")
	fmt.Printf("\n++Patient Information++\nFirstname: %s\nLastname:%s\nAge: %d\nID: %s\n", patient.FirstName, patient.LastName, patient.Age, patient.ID) // Use %v to handle generic types
	fmt.Println("Does the patient need Drug ?")
	fmt.Println("[1] Yes")
	fmt.Println("[2] No")

	reader := bufio.NewReader(os.Stdin)
	cmd, _ := reader.ReadString('\n')
	cmd = cmd[:len(cmd)-1] // Remove the trailing newline character
	if cmd == "1" {
		fmt.Println("Please enter the drugs name or enter 0 to finish the visit")
		for true{
			Drugs, _ := reader.ReadString('\n')
			Drugs = Drugs[:len(Drugs)-1]
			if Drugs == "0" {
				break
			}
			patient.DrugAllergies.Push(Drugs)
		}
		
	}

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
	CardiologyDB := DataStructures.NewHashMap(100)
	EmergencyDB := DataStructures.NewHashMap(100)
	DoctorsDB.Insert("Cardiology" , CardiologyDB)
	DoctorsDB.Insert("Emergency" , EmergencyDB)
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
					choice := Doctor_menu(*currentUser)
					for true {
						if choice == 1 {
							p , _:= currentUser.VisitQueue.Pop()
							patient := p.(Entities.Patient)
							visit(patient)
							currentUser.PatientList.AddToStart(patient)
							d , _:= patient.DrugAllergies.Peek()
							fmt.Println("Drugs : " , d)
							fmt.Println(patient.DrugAllergies)
							currentUser.PatientList.Display()

						} else if choice == 2 {
							break
						} else if choice == 3 {
							break
						}
						choice = Doctor_menu(*currentUser)
					}
				} else if our_type == reflect.TypeOf(&Entities.Patient{}) {
					currentUser := user.(*Entities.Patient)
					inner_cmd := Patient_menu()
					for true {
						if inner_cmd == 1 {
							clinic := Book_appointment()
							if clinic == "Back" {
								// break
								continue
							}
							clinicDBInterface , _ := DoctorsDB.GetRecursive(clinic)
							clinicDB := clinicDBInterface.(*DataStructures.HashMap)
							choose_doc(*currentUser , clinicDB)
							

						} else if inner_cmd == 2 {
							currentUser.DoctorList.Display()
							cancel_appointment(*currentUser)

						} else if inner_cmd == 3 {
							break
						} else if inner_cmd == 4 {
							// back = true
							break
						}
						inner_cmd = Patient_menu()
					}

				} else if our_type == reflect.TypeOf(&Entities.Manager{}) {
					currentUser := user.(*Entities.Manager)
					fmt.Printf("Doctor: %v %v, Department: %v\n", currentUser.FirstName, currentUser.LastName, currentUser.Department)

				}

				// break
				// continue
			}	
		} else if cmd == 3 {

			break
		}
		cmd = greet()
		clear()
	}
	
	CardiologyDB.Display()
	


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