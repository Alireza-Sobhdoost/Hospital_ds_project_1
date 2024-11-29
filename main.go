package main

import (
	"fmt"
	"log"
	"project_1/DataStructures" 
	"project_1/Entities"
	"project_1/Auth"

)

func main() {
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
	DataBase := DataStructures.NewHashMap(100)
	args := []string{"Doctor", "Cardiology"}
	err := Auth.Signup("1", "Alice", "Smith", "password", args ,  20 , *DataBase)
	
	if err != nil {
		log.Fatalf("Error setting password for doctor: %v", err)
	}


	doc , _:= Auth.Login(*DataBase, "1", "password")
	doctor := doc.(*Entities.Doctor)
	fmt.Printf("Doctor: %v %v, Department: %v\n", doctor.FirstName, doctor.LastName, doctor.Department)


	doctor.VisitQueue.Push(Entities.Patient{User: Entities.User{
		ID:        "10",
		FirstName: "John",
		LastName:  "Doe",
	}, PriorityToVsit: 3})

	doctor.VisitQueue.Push(Entities.Patient{User: Entities.User{
		ID:        "20",
		FirstName: "Jane",
		LastName:  "Smith",
	}, PriorityToVsit: 1})

	p1 := &Entities.Patient{User: Entities.User{
		ID:        "30",
		FirstName: "Emily",
		LastName:  "Davis",
	}, PriorityToVsit: 2}
	p1.MedicalHistory = "High blood pressure"
	fmt.Printf("Patient: %v %v, Priority: %v, Medical History: %v\n", p1.FirstName, p1.LastName, p1.PriorityToVsit, p1.MedicalHistory)

	doctor.VisitQueue.Push(*p1)
	
	fmt.Println("Patients in visit order:")

	for !doctor.VisitQueue.IsEmpty() {
		patient, err := doctor.VisitQueue.Pop()
		Patient := patient.(Entities.Patient)
		if err != nil {
			fmt.Println("Error:", err)
			break
		}
		fmt.Printf("%v\n", Patient.FirstName)
	}

	doc2 , _:= Auth.Login(*DataBase, "1", "password")
	doctor2 := doc2.(*Entities.Doctor)

	for !doctor2.VisitQueue.IsEmpty() {
		patient, err := doctor2.VisitQueue.Pop()
		Patient := patient.(Entities.Patient)
		if err != nil {
			fmt.Println("Error:", err)
			break
		}
		fmt.Printf("%v\n", Patient.FirstName)
	}



}