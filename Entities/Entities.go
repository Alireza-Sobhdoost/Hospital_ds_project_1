package Entities


import (
	"project_1/DataStructures"
	// "fmt"
	// "log"
	"golang.org/x/crypto/bcrypt"
)

// Base User struct
type User struct {
	ID        string
	FirstName string
	LastName  string
	Age int
	Role string
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
	PriorityToVsit int
}

// Manager struct (inherits from User)
type Manager struct {
	User
	Department string
}

