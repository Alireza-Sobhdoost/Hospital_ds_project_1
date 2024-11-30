package Auth

import (
	"project_1/Entities"
	"project_1/DataStructures"
	"fmt"
	"reflect"
)



func Signup(NID, firstName, lastName, password string, args []string , age int, DataBase DataStructures.HashMap) (error) {
	// Create a base User entity
	user := &Entities.User{
		ID:        NID,
		FirstName: firstName,
		LastName:  lastName,
		Role: args[0],
		Age : age,
		Password:  password,
	}

	// Set the password (error handling)
	err := user.SetPassword(password)
	if err != nil {
		return  fmt.Errorf("error setting password: %w", err)
	}

	// Create role-specific entities
	switch args[0] {
		case "Patient":
			patient := &Entities.Patient{
				User: *user,
				PriorityToVsit: 5,
				DrugAllergies: DataStructures.NewStack(),
				DoctorList: DataStructures.NewLinkedList(),
				MedicalHistory: "",
			}
			DataBasePatientsInterface, _ := DataBase.Get("Patients")
			DataBasePatients := DataBasePatientsInterface.(*DataStructures.HashMap)
			DataBasePatients.Insert(NID, patient)
			return nil

		case "Doctor":
			doctor := &Entities.Doctor{
				User: *user,
				Department: args[1],
				PatientList: DataStructures.LinkedList{},
				VisitQueue: DataStructures.NewPriorityQueue(func(a, b interface{}) bool {
					patientA := a.(Entities.Patient)
					patientB := b.(Entities.Patient)
					return patientA.PriorityToVsit < patientB.PriorityToVsit
				}),
			}
			DataBaseDoctorsInterface, _ := DataBase.Get("Doctors")
			DataBaseDoctors := DataBaseDoctorsInterface.(*DataStructures.HashMap)
			DepartmentInterface , _ := DataBaseDoctors.Get(args[1])
			Department := DepartmentInterface.(*DataStructures.HashMap)
			Department.Insert(NID, doctor)
			return nil

		case "Manager":
			manager := &Entities.Manager{
				User: *user,
				ToAddStack: DataStructures.NewStack() ,
			}
			DataBaseManagersInterface, _ := DataBase.Get("Managers")
			DataBaseManagers := DataBaseManagersInterface.(*DataStructures.HashMap)
			DataBaseManagers.Insert(NID, manager)
			return nil

		case "DrugMan":
			DrugMan := &Entities.DrugMan{
				User: *user,
			}
			DataBaseDrugMansInterface, _ := DataBase.Get("DrugMans")
			DataBaseDrugMans := DataBaseDrugMansInterface.(*DataStructures.HashMap)
			DataBaseDrugMans.Insert(NID, DrugMan)
			return nil

		case "Triage":
			DrugMan := &Entities.Triage{
				User: *user,
			}
			DataBaseTriagesInterface, _ := DataBase.Get("Triages")
			DataBaseTriages := DataBaseTriagesInterface.(*DataStructures.HashMap)
			DataBaseTriages.Insert(NID, DrugMan)
			return nil

		default:
			return fmt.Errorf("invalid role: %s", args[0])
		}
}

func SignupEntity(entity interface{}, DataBase DataStructures.HashMap) (error) {
	// Create a base User entity

	our_type := reflect.TypeOf(entity)
	if our_type == reflect.TypeOf(&Entities.Doctor{}) {
		userInterface := entity.(*Entities.Doctor)
		DataBaseDoctorsInterface, _ := DataBase.Get("Doctors")
		DataBaseDoctors := DataBaseDoctorsInterface.(*DataStructures.HashMap)
		DepartmentInterface , _ := DataBaseDoctors.Get(userInterface.Department)
		Department := DepartmentInterface.(*DataStructures.HashMap)
		Department.Insert(userInterface.ID, userInterface)
		fmt.Println("register successfull")
		return nil
	} else if our_type == reflect.TypeOf(&Entities.Patient{}) {
		userInterface := entity.(*Entities.Patient)
		DataBasePatientsInterface, _ := DataBase.Get("Patients")
		DataBasePatients := DataBasePatientsInterface.(*DataStructures.HashMap)
		DataBasePatients.Insert(userInterface.ID, userInterface)
		return nil
	} else if our_type == reflect.TypeOf(&Entities.DrugMan{}) {
		userInterface := entity.(*Entities.DrugMan)
		DataBaseDrugMansInterface, _ := DataBase.Get("DrugMans")
			DataBaseDrugMans := DataBaseDrugMansInterface.(*DataStructures.HashMap)
			DataBaseDrugMans.Insert(userInterface.ID, userInterface)
			return nil
	} else if our_type == reflect.TypeOf(&Entities.Triage{}) {
		userInterface := entity.(*Entities.Triage)
		DataBaseTriagesInterface, _ := DataBase.Get("Triages")
		DataBaseTriages := DataBaseTriagesInterface.(*DataStructures.HashMap)
		DataBaseTriages.Insert(userInterface.ID, userInterface)
		return nil
	}
	return nil
}
func Login(DataBase DataStructures.HashMap ,NID, password string) (interface{}, error) {

	user, ok := DataBase.GetRecursive(NID)
	if !ok {
		return nil, fmt.Errorf("user not found")
	}
	our_type := reflect.TypeOf(user)
	if our_type == reflect.TypeOf(&Entities.Doctor{}) {
		userInterface := user.(*Entities.Doctor)
		if !userInterface.ValidatePassword(password) {
			return nil, fmt.Errorf("invalid password")
		}
		return userInterface, nil
	} else if our_type == reflect.TypeOf(&Entities.Patient{}) {
		userInterface := user.(*Entities.Patient)
		if !userInterface.ValidatePassword(password) {
			return nil, fmt.Errorf("invalid password")
		}
		return userInterface, nil
	} else if our_type == reflect.TypeOf(&Entities.Manager{}) {
		userInterface := user.(*Entities.Manager)
		if !userInterface.ValidatePassword(password) {
			return nil, fmt.Errorf("invalid password")
		}
		return userInterface, nil
	} else if our_type == reflect.TypeOf(&Entities.DrugMan{}) {
		userInterface := user.(*Entities.DrugMan)
		if !userInterface.ValidatePassword(password) {
			return nil, fmt.Errorf("invalid password")
		}
		return userInterface, nil
	} else if our_type == reflect.TypeOf(&Entities.Triage{}) {
		userInterface := user.(*Entities.Triage)
		if !userInterface.ValidatePassword(password) {
			return nil, fmt.Errorf("invalid password")
		}
		return userInterface, nil
	}

	return nil , fmt.Errorf("user not found")
}