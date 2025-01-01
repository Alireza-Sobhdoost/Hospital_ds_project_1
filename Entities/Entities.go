package Entities

import (
	"fmt"
	"project_1/DataStructures"

	// "log"
	"golang.org/x/crypto/bcrypt"
)

// Base User struct
type User struct {
	ID        string
	FirstName string
	LastName  string
	Age       int
	Role      string
	Password  string // Store the hashed password
}

// SetPassword hashes the plain text password and stores it in the User struct
func (u *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// ValidatePassword compares the stored hashed password with a plain text password
func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// Doctor struct (inherits from User)
type Doctor struct {
	User
	Department  string
	PatientList DataStructures.LinkedList
	VisitQueue  *DataStructures.PriorityQueue // Change to pointer
}

// Patient struct (inherits from User)
type Patient struct {
	User
	MedicalHistory string
	DrugAllergies  *DataStructures.Stack
	DoctorList     *DataStructures.LinkedList
	PriorityToVsit int
}

// Manager struct (inherits from User)
type Manager struct {
	User
	Department string
	ToAddStack *DataStructures.Stack
}

type DrugMan struct {
	User
}
type Triage struct {
	User
}

func DisplayDocs(hm *DataStructures.HashMap) (DataStructures.LinkedList, int) {

	doclist := DataStructures.LinkedList{}
	count := 0

	// Iterate over the buckets in the hash map
	for _, bucket := range hm.Buckets {
		if len(bucket) > 0 {
			// fmt.Printf("Bucket %d: ", i)
			for _, kv := range bucket {
				// Safely assert that the value is a *Doctor
				doc, ok := kv.Value.(*Doctor)
				if ok {
					fmt.Printf("[%d] Dr.%s %s ", count+1, doc.FirstName, doc.LastName)
					doclist.AddToEnd(doc)
					count += 1

				} else {
					fmt.Printf("[Unknown value type] ")
				}
				fmt.Println()

			}
			fmt.Println()
		}
	}

	return doclist, count
}

func DisplayDocsList(list DataStructures.LinkedList) int {
	current := list.Head
	counter := 0
	for current != nil {
		fmt.Printf("[%d] ", counter+1)
		fmt.Printf(" " + current.Data.(*Doctor).FirstName + " " + current.Data.(*Doctor).LastName + " \n")
		current = current.Next
		counter += 1
	}

	return counter
}

func DisplayPatList(list DataStructures.LinkedList) int {
	current := list.Head
	counter := 0
	for current != nil {
		fmt.Printf("[%d] ", counter+1)
		fmt.Printf(" " + current.Data.(Patient).FirstName + " " + current.Data.(Patient).LastName + current.Data.(Patient).ID + " \n")
		current = current.Next
		counter += 1
	}

	return counter
}

func DisplayPatList2(list DataStructures.LinkedList) int {
	current := list.Head
	counter := 0
	for current != nil {
		fmt.Printf("[%d] ", counter+1)
		fmt.Printf(" " + current.Data.(*Patient).FirstName + " " + current.Data.(*Patient).LastName + current.Data.(*Patient).ID + " \n")
		current = current.Next
		counter += 1
	}

	return counter
}
