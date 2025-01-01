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

func greet() int {
	fmt.Println("==Hospital==")
	fmt.Println("Wellcome to our Hospital. How can we help you?")
	fmt.Println("[1] Sign up")
	fmt.Println("[2] Login")
	fmt.Println("[3] Emergency")
	fmt.Println("[4] Exit")

	reader := bufio.NewReader(os.Stdin)
	cmd, _ := reader.ReadString('\n')
	cmd = cmd[:len(cmd)-1] // Remove the trailing newline character
	Intcmd, _ := strconv.Atoi(cmd)
	return Intcmd

}
func signup_form() (interface{}, string, string, string, string, []string, int, bool) {
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
		return nil, "", "", "", "", nil, 0, false
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
	if role == "Patient" {
		args := []string{role}
		return nil, nationalID, firstName, lastName, password, args, age, true
	}
	user := &Entities.User{
		ID:        nationalID,
		FirstName: firstName,
		LastName:  lastName,
		Role:      role,
		Age:       age,
		Password:  password,
	}
	user.SetPassword(password)

	switch role {

	case "Doctor":
		doctor := &Entities.Doctor{
			User:        *user,
			Department:  department,
			PatientList: DataStructures.LinkedList{},
			VisitQueue: DataStructures.NewPriorityQueue(func(a, b interface{}) bool {
				patientA := a.(Entities.Patient)
				patientB := b.(Entities.Patient)
				return patientA.PriorityToVsit < patientB.PriorityToVsit
			}),
		}

		return doctor, "", "", "", "", []string{}, 0, false

	case "DrugMan":
		DrugMan := &Entities.DrugMan{
			User: *user,
		}
		return DrugMan, "", "", "", "", []string{}, 0, false

	case "Triage":
		triage := &Entities.Triage{
			User: *user,
		}
		return triage, "", "", "", "", []string{}, 0, false

	default:
		return nil, "", "", "", "", []string{}, 0, false
	}
}

func edit_form() []string {
	fmt.Println("==edit==")
	fmt.Println("Please enter your information in order below:")
	fmt.Println("First name, Last name, National ID, password, age, role")
	fmt.Println("If you dont want to change a part fill it with / ")

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

	// // Collect the information
	// fmt.Println("\nCollected Information:")
	// fmt.Println("First Name:", firstName)
	// fmt.Println("Last Name:", lastName)
	// fmt.Println("National ID:", nationalID)
	// fmt.Println("Password:", password)
	// fmt.Println("Age:", age)
	// fmt.Println("Role:", role)

	args := []string{firstName, lastName, password, nationalID, ageStr}
	return args

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

func Patient_menu() int {
	fmt.Println("==Patient Menu==")
	fmt.Println("Please enter your choice:")
	fmt.Println("[1] Book an appointment")
	fmt.Println("[2] Cancel appointments")
	fmt.Println("[3] Go to Drugstore")
	fmt.Println("[4] Edit account")
	fmt.Println("[5] Exit")

	reader := bufio.NewReader(os.Stdin)
	cmd, _ := reader.ReadString('\n')
	cmd = cmd[:len(cmd)-1] // Remove the trailing newline character
	Intcmd, _ := strconv.Atoi(cmd)
	clear()
	return Intcmd
}

func Book_appointment() string {
	fmt.Println("==Patient Menu==")
	fmt.Println("Please enter your choice:")
	fmt.Println("[1] Cardiology")
	fmt.Println("[2] Emergency")
	fmt.Println("[3] Back")

	reader := bufio.NewReader(os.Stdin)
	cmd, _ := reader.ReadString('\n')
	cmd = cmd[:len(cmd)-1] // Remove the trailing newline character
	Intcmd, _ := strconv.Atoi(cmd)
	clear()

	if Intcmd == 1 {
		return "Cardiology"
	} else if Intcmd == 2 {
		return "Emergency"
	} else {
		return "Back"
	}
}

func choose_doc(caller Entities.Patient, DB *DataStructures.HashMap) {
	fmt.Println("==Choose a doctor==\n")
	DocsList, lenght := Entities.DisplayDocs(DB)
	// DocsList.Display()
	fmt.Println("[e] back")
	reader := bufio.NewReader(os.Stdin)
	cmd, _ := reader.ReadString('\n')
	cmd = cmd[:len(cmd)-1] // Remove the trailing newline character
	if cmd == "e" {
		return
	}
	Intcmd, _ := strconv.Atoi(cmd)
	clear()
	doc_internal_pointer_var := DocsList.Find_by_index(Intcmd-1, lenght)
	// fmt.Println(doc_internal_pointer_var)
	doc_internal_pointer_var.Data.(*Entities.Doctor).VisitQueue.Push(caller)
	caller.DoctorList.AddToStart(doc_internal_pointer_var.Data.(*Entities.Doctor))
	caller.DoctorList.Display()
	fmt.Println("You have been added to the queue")

}

func cancel_appointment(caller Entities.Patient) {
	fmt.Println("==Choose a appointment==\n")
	// caller.DoctorList.Display()

	lenght := Entities.DisplayDocsList(*caller.DoctorList)
	fmt.Println("[e] back")

	reader := bufio.NewReader(os.Stdin)
	cmd, _ := reader.ReadString('\n')
	cmd = cmd[:len(cmd)-1] // Remove the trailing newline character
	if cmd == "e" {
		return
	}
	Intcmd, _ := strconv.Atoi(cmd)
	clear()
	doc_internal_pointer_var := caller.DoctorList.Find_by_index(Intcmd-1, lenght)
	// fmt.Println(doc_internal_pointer_var)
	doc_internal_pointer_var.Data.(*Entities.Doctor).VisitQueue.Remove(caller)
	caller.DoctorList.Remove(doc_internal_pointer_var.Data.(*Entities.Doctor))
	fmt.Println("You have been added to the queue")

}

func Doctor_menu(doc Entities.Doctor) int {

	clear()
	fmt.Println("==Doctor Menu==")
	fmt.Println("Hello Dr. ", doc.FirstName, " ", doc.LastName)
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
		fmt.Println("There are ", count, " patients to visit")
	}

	fmt.Println("---------------------------------------------")

	fmt.Println("[1] Start to visit")
	fmt.Println("[2] See Patient list history")
	fmt.Println("[3] Edit account")
	fmt.Println("[4] Exit")

	reader := bufio.NewReader(os.Stdin)
	cmd, _ := reader.ReadString('\n')
	cmd = cmd[:len(cmd)-1] // Remove the trailing newline character
	Intcmd, _ := strconv.Atoi(cmd)
	return Intcmd

}

func visit(patient Entities.Patient) {

	clear()
	fmt.Println("==Visit==")
	fmt.Printf("\nPatient Information: \nFirstname: %s\nLastname:%s\nAge: %d\nID: %s\n", patient.FirstName, patient.LastName, patient.Age, patient.ID) // Use %v to handle generic types
	fmt.Println("Does the patient need Drug ?")
	fmt.Println("[1] Yes")
	fmt.Println("[2] No")

	reader := bufio.NewReader(os.Stdin)
	cmd, _ := reader.ReadString('\n')
	cmd = cmd[:len(cmd)-1] // Remove the trailing newline character
	if cmd == "1" {
		fmt.Println("Please enter the drugs name or enter 0 to finish the visit")
		for true {
			Drugs, _ := reader.ReadString('\n')
			Drugs = Drugs[:len(Drugs)-1]
			if Drugs == "0" {
				break
			}
			patient.DrugAllergies.Push(Drugs)
		}

	}

}

func DrugStore_menu() int {

	fmt.Println("==Drug Store Menu==")
	fmt.Println("Please enter your choice:")
	fmt.Println("[1] Customer Service")
	fmt.Println("[2] Add New Drug")
	fmt.Println("[3] Search Drug by ID")
	fmt.Println("[4] Search Drug by Name")
	fmt.Println("[5] Search Drug by Type")
	fmt.Println("[6] Search Drug by Price range")
	fmt.Println("[7] Get Cheapest Drug")
	fmt.Println("[8] Get the most expencive Drug")
	fmt.Println("[9] Delete Drug")
	fmt.Println("[10] See All Drugs (Sorted by ID)")
	fmt.Println("[11] Edit account")
	fmt.Println("[12] Exit")
	fmt.Println("[13] Get Total Drug Count")
	fmt.Println("[14] Get BST Depth")

	reader := bufio.NewReader(os.Stdin)
	cmd, _ := reader.ReadString('\n')
	cmd = cmd[:len(cmd)-1]
	Intcmd, _ := strconv.Atoi(cmd)
	clear()
	return Intcmd

}

func DrugStore_menu_Patient() int {

	fmt.Println("==Drug Store Menu==")
	fmt.Println("Please enter your choice:")
	fmt.Println("[1] Search Drug by ID")
	fmt.Println("[2] Search Drug by Name")
	fmt.Println("[3] Search Drug by Type")
	fmt.Println("[4] Search Drug by Price range")
	fmt.Println("[5] Exit")

	reader := bufio.NewReader(os.Stdin)
	cmd, _ := reader.ReadString('\n')
	cmd = cmd[:len(cmd)-1]
	Intcmd, _ := strconv.Atoi(cmd)
	clear()
	return Intcmd

}

func addNewDrug(drugBST *DataStructures.DrugBST) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter Drug Details:")
	fmt.Print("ID: ")
	id, _ := reader.ReadString('\n')
	id = id[:len(id)-1]

	fmt.Print("Name: ")
	name, _ := reader.ReadString('\n')
	name = name[:len(name)-1]

	fmt.Print("Price: ")
	priceStr, _ := reader.ReadString('\n')
	priceStr = priceStr[:len(priceStr)-1]
	price, _ := strconv.ParseFloat(priceStr, 64)

	fmt.Print("Type: ")
	drugType, _ := reader.ReadString('\n')
	drugType = drugType[:len(drugType)-1]

	fmt.Print("Dose: ")
	dose, _ := reader.ReadString('\n')
	dose = dose[:len(dose)-1]

	err := drugBST.Insert(id, name, price, drugType , dose)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("Drug added successfully!")
}

func searchDrugByID(drugBST *DataStructures.DrugBST) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Drug ID: ")
	id, _ := reader.ReadString('\n')
	id = id[:len(id)-1]

	drug, err := drugBST.SearchByID(id)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Found Drug: ID=%s, Name=%s, Price=%.2f, Type=%s Dose= %s, Count= %d\n",
	drug.ID, drug.Name, drug.Price, drug.Type , drug.Dose , drug.Count)
}

func searchDrugByName(drugBST *DataStructures.DrugBST) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Drug Name or partial name: ")
	name, _ := reader.ReadString('\n')
	name = name[:len(name)-1]

	// Use auto complete to find matching drugs
	results, err := drugBST.Trie.SearchByName(name)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if results.Head == nil {
		fmt.Println("No matching drugs found")
		return
	}

	fmt.Println("Found Drugs:")
	results.DisplayDrugs()
	
}

func searchDrugByType(drugBST *DataStructures.DrugBST) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Drug Type or partial type: ")
	drugType, _ := reader.ReadString('\n')
	drugType = drugType[:len(drugType)-1]

	// Use auto complete to find matching drug types
	results, err := drugBST.SearchByType(drugType)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if results.Head == nil {
		fmt.Println("No matching drug types found")
		return
	}

	fmt.Println("Found Drugs:")
	results.DisplayDrugs()
}

// func autoCompleteDrug(drugBST *DataStructures.DrugBST) {
// 	reader := bufio.NewReader(os.Stdin)

// 	fmt.Print("Enter Drug Name Prefix: ")
// 	prefix, _ := reader.ReadString('\n')
// 	prefix = prefix[:len(prefix)-1]

// 	results, err := drugBST.Trie.AutoComplete(prefix)
// 	if err != nil {
// 		fmt.Printf("Error: %v\n", err)
// 		return
// 	}

// 	fmt.Println("Suggestions:")
// 	for _, drug := range results {
// 		fmt.Printf("ID=%s, Name=%s, Price=%.2f, Type=%s\n",
// 			drug.ID, drug.Name, drug.Price, drug.Type)
// 	}
// }

func deleteDrug(drugBST *DataStructures.DrugBST) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Drug ID to delete: ")
	id, _ := reader.ReadString('\n') 
	id = id[:len(id)-1]

	// First get the drug name before deleting from BST
	drug, err := drugBST.SearchByID(id)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Delete from BST
	err = drugBST.Delete(id)
	if err != nil {
		fmt.Printf("Error: %v\n", err) 
		return
	}

	// Also delete from Trie
	drugBST.Trie.Delete(drug.Name)

	fmt.Println("Drug deleted successfully from BST and Trie!")
}

// func DrugStore_workflow(drugman *Entities.DrugMan, DB *DataStructures.HashMap, drugBST *DataStructures.DrugBST) {
// 	for {
// 		cmd := DrugStore_menu()
// 		switch cmd {
// 		case 1:
// 			DrugStore_Csevent(*drugman, DB)
// 		case 2:
// 			addNewDrug(drugBST)
// 		case 3:
// 			searchDrugByID(drugBST)
// 		case 4:
// 			searchDrugByName(drugBST)
// 		case 5:
// 			searchDrugByType(drugBST)
// 		case 6:
// 			cheapest := drugBST.GetCheapestDrug()
// 			if cheapest != nil {
// 				fmt.Printf("Cheapest Drug: ID=%s, Name=%s, Price=%.2f, Type=%s\n",
// 					cheapest.ID, cheapest.Name, cheapest.Price, cheapest.Type)
// 			} else {
// 				fmt.Println("No drugs in inventory")
// 			}
// 		case 7:
// 			deleteDrug(drugBST)
// 		case 8:
// 			edit_menu(drugman)
// 		case 9:
// 			return
// 		}
// 	}
// }

func DrugStore_Csevent(drugman Entities.DrugMan, DB *DataStructures.HashMap , drugBST *DataStructures.DrugBST) {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("==Drug Store Menu==")
	fmt.Println("Please enter your Customer ID :")

	nationalID, _ := reader.ReadString('\n')
	nationalID = nationalID[:len(nationalID)-1]

	custumer, _ := DB.GetRecursive(nationalID)
	petient := custumer.(*Entities.Patient)
	clear()
	fmt.Println("==Customer Drugs==")
	fmt.Println("--------------------------------------------------------------")
	petient.DrugAllergies.PrintByPoppingCopy()
	fmt.Println("--------------------------------------------------------------")

	fmt.Println("Please enter your choice:")
	fmt.Println("[1] Start making Drugs from Doctor's prescription")
	fmt.Println("[2] Start making Drugs without Doctor's prescription")
	fmt.Println("[3] Exit")

	cmd, _ := reader.ReadString('\n')
	cmd = cmd[:len(cmd)-1] // Remove the trailing newline character
	Intcmd, _ := strconv.Atoi(cmd)
	clear()
	if Intcmd == 1 {
		fmt.Println("Please verify the drugs name by enter 1 or enter 0 to finish")
		for !petient.DrugAllergies.IsEmpty() {
			drug, _ := petient.DrugAllergies.Peek()
			fmt.Println(drug.(string))
			Continue, _ := reader.ReadString('\n')
			Continue = Continue[:len(Continue)-1]
			if Continue == "0" {
				break
			}
			if Continue == "1" {
				results , err := drugBST.Trie.SearchByName(drug.(string))
				if err != nil {
					fmt.Printf("Error: %v\n", err)
					return
				}
			
				if results.Head == nil {
					fmt.Println("No matching drugs found")
					return
				}
			
				fmt.Println("Found Drugs:")
				len_list := results.DisplayDrugs()
				drug_cmd, _ := reader.ReadString('\n')
				drug_cmd = drug_cmd[:len(drug_cmd)-1] // Remove the trailing newline character
				Intdrugcmd, _ := strconv.Atoi(drug_cmd)
				if Intdrugcmd != 0 {
					given_drug := results.Find_by_index(Intdrugcmd - 1 ,len_list).Data.(*DataStructures.DrugNode)
					fmt.Printf("The given drug : ID=%s, Name=%s, Price=%.2f, Type=%s, Dose=%s, Count=%d\n" ,
					given_drug.ID, given_drug.Name, given_drug.Price, given_drug.Type, given_drug.Dose, given_drug.Count)
					petient.DrugAllergies.Pop()
					given_drug.Count -= 1
					if given_drug.Count == 0 {
						drugBST.Delete(given_drug.ID)
					}
				}
			}

		}
	} else if Intcmd == 2 {
		innercmd := DrugStore_menu_Patient()
		for true {
			if innercmd == 1 {
				reader := bufio.NewReader(os.Stdin)

				fmt.Print("Enter Drug ID: ")
				id, _ := reader.ReadString('\n')
				id = id[:len(id)-1]
			
				drug, err := drugBST.SearchByID(id)
				if err != nil {
					fmt.Printf("Error: %v\n", err)
					return
				}
				fmt.Printf("Found Drug: ID=%s, Name=%s, Price=%.2f, Type=%s Dose= %s, Count= %d\n",
					drug.ID, drug.Name, drug.Price, drug.Type , drug.Dose , drug.Count)
				verfy, _ := reader.ReadString('\n')
				verfy = verfy[:len(verfy)-1]
				if verfy == "1"{
					drug.Count -= 1 
					if drug.Count == 0 {
						drugBST.Delete(drug.ID)
					}
				}
			} else if innercmd == 2 {
				drugname, _ := reader.ReadString('\n')
				drugname = drugname[:len(drugname)-1]
				results , err := drugBST.Trie.SearchByName(drugname)
				if err != nil {
					fmt.Printf("Error: %v\n", err)
					return
				}
			
				if results.Head == nil {
					fmt.Println("No matching drugs found")
					return
				}
			
				fmt.Println("Found Drugs:")
				len_list := results.DisplayDrugs()
				drug_cmd, _ := reader.ReadString('\n')
				drug_cmd = drug_cmd[:len(drug_cmd)-1] // Remove the trailing newline character
				Intdrugcmd, _ := strconv.Atoi(drug_cmd)
				if Intdrugcmd != 0 {
					given_drug := results.Find_by_index(Intdrugcmd - 1 ,len_list).Data.(*DataStructures.DrugNode)
					fmt.Printf("The given drug : ID=%s, Name=%s, Price=%.2f, Type=%s, Dose=%s, Count=%d\n" ,
					given_drug.ID, given_drug.Name, given_drug.Price, given_drug.Type, given_drug.Dose, given_drug.Count)
					given_drug.Count -= 1
					if given_drug.Count == 0 {
						drugBST.Delete(given_drug.ID)
					}
				}
			} else if innercmd == 3 {
				Typeame, _ := reader.ReadString('\n')
				Typeame = Typeame[:len(Typeame)-1]
				results , err := drugBST.SearchByType(Typeame)
				if err != nil {
					fmt.Printf("Error: %v\n", err)
					return
				}
			
				if results.Head == nil {
					fmt.Println("No matching drugs found")
					return
				}
			
				fmt.Println("Found Drugs:")
				len_list := results.DisplayDrugs()
				drug_cmd, _ := reader.ReadString('\n')
				drug_cmd = drug_cmd[:len(drug_cmd)-1] // Remove the trailing newline character
				Intdrugcmd, _ := strconv.Atoi(drug_cmd)
				if Intdrugcmd != 0 {
					given_drug := results.Find_by_index(Intdrugcmd - 1 ,len_list).Data.(*DataStructures.DrugNode)
					fmt.Printf("The given drug : ID=%s, Name=%s, Price=%.2f, Type=%s, Dose=%s, Count=%d\n" ,
					given_drug.ID, given_drug.Name, given_drug.Price, given_drug.Type, given_drug.Dose, given_drug.Count)
					given_drug.Count -= 1
					if given_drug.Count == 0 {
						drugBST.Delete(given_drug.ID)
					}
				}
			} else if innercmd == 4 {
				fmt.Print("Enter minimum price: ")
				minStr, _ := reader.ReadString('\n')
				minStr = minStr[:len(minStr)-1]
				minPrice, err := strconv.ParseFloat(minStr, 64)
				if err != nil {
					fmt.Println("Invalid minimum price")
					return
				}
				
				fmt.Print("Enter maximum price: ")
				maxStr, _ := reader.ReadString('\n')
				maxStr = maxStr[:len(maxStr)-1]
				maxPrice, err := strconv.ParseFloat(maxStr, 64)
				if err != nil {
					fmt.Println("Invalid maximum price")
					return
				}
				
				results, err := drugBST.SearchByPriceRange(minPrice, maxPrice)
				if err != nil {
					fmt.Println("No drugs found in this price range")
					return
				}
				
				fmt.Printf("\nDrugs between %.2f and %.2f:\n", minPrice, maxPrice)
				len_list := results.DisplayDrugs()
				drug_cmd, _ := reader.ReadString('\n')
				drug_cmd = drug_cmd[:len(drug_cmd)-1] // Remove the trailing newline character
				Intdrugcmd, _ := strconv.Atoi(drug_cmd)
				if Intdrugcmd != 0 {
					given_drug := results.Find_by_index(Intdrugcmd - 1 ,len_list).Data.(*DataStructures.DrugNode)
					fmt.Printf("The given drug : ID=%s, Name=%s, Price=%.2f, Type=%s, Dose=%s, Count=%d\n" ,
					given_drug.ID, given_drug.Name, given_drug.Price, given_drug.Type, given_drug.Dose, given_drug.Count)
					given_drug.Count -= 1
					if given_drug.Count == 0 {
						drugBST.Delete(given_drug.ID)
					}
				}

			} else if innercmd == 5 {
				break
			}
			innercmd = DrugStore_menu_Patient()
		}	
	} else if Intcmd == 3 {
		return
	}

}

func searchByPriceRange(drugBST *DataStructures.DrugBST) {
    reader := bufio.NewReader(os.Stdin)
    
    fmt.Print("Enter minimum price: ")
    minStr, _ := reader.ReadString('\n')
    minStr = minStr[:len(minStr)-1]
    minPrice, err := strconv.ParseFloat(minStr, 64)
    if err != nil {
        fmt.Println("Invalid minimum price")
        return
    }
    
    fmt.Print("Enter maximum price: ")
    maxStr, _ := reader.ReadString('\n')
    maxStr = maxStr[:len(maxStr)-1]
    maxPrice, err := strconv.ParseFloat(maxStr, 64)
    if err != nil {
        fmt.Println("Invalid maximum price")
        return
    }
    
    results, err := drugBST.SearchByPriceRange(minPrice, maxPrice)
    if err != nil {
        fmt.Println("No drugs found in this price range")
        return
    }
    
    fmt.Printf("\nDrugs between %.2f and %.2f:\n", minPrice, maxPrice)
	results.DisplayDrugs()
}

func seeAllDrugs(drugBST *DataStructures.DrugBST) {
	drugs := drugBST.InOrderTraversalByID()
	if len(drugs) == 0 {
		fmt.Println("No drugs in inventory")
		return
	}

	fmt.Println("All Drugs (Sorted by ID):")
	for _, drug := range drugs {
		fmt.Printf("ID=%s, Name=%s, Price=%.2f, Type=%s, Dose=%s, Count=%d\n",
			drug.ID, drug.Name, drug.Price, drug.Type, drug.Dose, drug.Count)
	}
}

func getTotalDrugCount(drugBST *DataStructures.DrugBST) {
	count := drugBST.CountAllDrugs()
	fmt.Printf("Total number of drugs in the BST: %d\n", count)
}

func getBSTDepth(drugBST *DataStructures.DrugBST) {
	depth := drugBST.GetBSTDepth()
	fmt.Printf("Depth of the BST: %d\n", depth)
}

func Triage_entry(DB DataStructures.HashMap, list *DataStructures.LinkedList) {
	fmt.Println("==Triage==")
	fmt.Println("Please enter your NID:")

	reader := bufio.NewReader(os.Stdin)
	nationalID, _ := reader.ReadString('\n')
	nationalID = nationalID[:len(nationalID)-1]

	custumer, _ := DB.GetRecursive(nationalID)
	if custumer == nil {
		args := []string{"Patient"}
		Auth.Signup(nationalID, "", "", nationalID, args, 0, DB)
		custumer, _ = DB.GetRecursive(nationalID)
	}
	petient := custumer.(*Entities.Patient)
	fmt.Println(petient.ID)
	list.AddToStart(petient)

}

func Triage_menu() int {
	fmt.Println("==Triage==")
	fmt.Println("[1]  See Patients")
	fmt.Println("[2]  Edit profile")
	fmt.Println("[3]  Exit")

	reader := bufio.NewReader(os.Stdin)
	cmd, _ := reader.ReadString('\n')
	cmd = cmd[:len(cmd)-1] // Remove the trailing newline character
	Intcmd, _ := strconv.Atoi(cmd)

	return Intcmd
}

func Triage_agent(EmergencyDB DataStructures.HashMap, list DataStructures.LinkedList) {
	fmt.Println("==Triage==")
	fmt.Println("See Petaints : ")
	fmt.Println("--------------------------------------------------------------")
	lenght := Entities.DisplayPatList2(list)
	fmt.Println("--------------------------------------------------------------")

	reader := bufio.NewReader(os.Stdin)
	cmd, _ := reader.ReadString('\n')
	cmd = cmd[:len(cmd)-1] // Remove the trailing newline character

	if cmd == "e" {
		return
	}
	Intcmd, _ := strconv.Atoi(cmd)

	custumer := list.Find_by_index(Intcmd-1, lenght)

	fmt.Println("Please set a proitery :")

	pr, _ := reader.ReadString('\n')
	pr = pr[:len(pr)-1] // Remove the trailing newline character
	Intpr, _ := strconv.Atoi(pr)
	custumer.Data.(*Entities.Patient).PriorityToVsit = Intpr
	choose_doc(*custumer.Data.(*Entities.Patient), &EmergencyDB)
	list.Remove(custumer.Data)

}

func edit_menu(user interface{}) {
	fmt.Println("==Edit Menu==")
	fmt.Println("this is your current data")
	reader := bufio.NewReader(os.Stdin)
	// Remove the trailing newline character
	our_type := reflect.TypeOf(user)
	if our_type == reflect.TypeOf(&Entities.Doctor{}) {
		currentUser := user.(*Entities.Doctor)
		fmt.Println("Fristname : %s , Lastname : %s , Password : %s , ID : %s , Age : %d", currentUser.FirstName, currentUser.LastName, currentUser.Password, currentUser.ID, currentUser.Age)
		fmt.Println("Do you want to change your data ?")
		fmt.Println("[1] Yes")
		fmt.Println("[2] No")
		cmd, _ := reader.ReadString('\n')
		cmd = cmd[:len(cmd)-1]
		if cmd == "2" {
			return
		}
		args := edit_form()
		if args[0] != "/" {
			currentUser.FirstName = args[0]
		} else if args[1] != "/" {
			currentUser.LastName = args[1]
		} else if args[2] != "/" {
			currentUser.Password = args[2]
		} else if args[3] != "/" {
			currentUser.ID = args[3]
		} else if args[4] != "/" {
			age, _ := strconv.Atoi(args[4])
			currentUser.Age = age
		}

	} else if our_type == reflect.TypeOf(&Entities.Manager{}) {
		currentUser := user.(*Entities.Manager)
		fmt.Println("Fristname : %s , Lastname : %s , Password : %s , ID : %s , Age : %d", currentUser.FirstName, currentUser.LastName, currentUser.Password, currentUser.ID, currentUser.Age)
		fmt.Println("Do you want to change your data ?")
		fmt.Println("[1] Yes")
		fmt.Println("[2] No")
		cmd, _ := reader.ReadString('\n')
		cmd = cmd[:len(cmd)-1]
		if cmd == "2" {
			return
		}
		args := edit_form()
		if args[0] != "/" {
			currentUser.FirstName = args[0]
		} else if args[1] != "/" {
			currentUser.LastName = args[1]
		} else if args[2] != "/" {
			currentUser.Password = args[2]
		} else if args[3] != "/" {
			currentUser.ID = args[3]
		} else if args[4] != "/" {
			age, _ := strconv.Atoi(args[4])
			currentUser.Age = age
		}

	} else if our_type == reflect.TypeOf(&Entities.DrugMan{}) {
		currentUser := user.(*Entities.DrugMan)
		fmt.Println("Fristname : %s , Lastname : %s , Password : %s , ID : %s , Age : %d", currentUser.FirstName, currentUser.LastName, currentUser.Password, currentUser.ID, currentUser.Age)
		fmt.Println("Do you want to change your data ?")
		fmt.Println("[1] Yes")
		fmt.Println("[2] No")
		cmd, _ := reader.ReadString('\n')
		cmd = cmd[:len(cmd)-1]
		if cmd == "2" {
			return
		}
		args := edit_form()
		if args[0] != "/" {
			currentUser.FirstName = args[0]
		} else if args[1] != "/" {
			currentUser.LastName = args[1]
		} else if args[2] != "/" {
			currentUser.Password = args[2]
		} else if args[3] != "/" {
			currentUser.ID = args[3]
		} else if args[4] != "/" {
			age, _ := strconv.Atoi(args[4])
			currentUser.Age = age
		}

	} else if our_type == reflect.TypeOf(Entities.Triage{}) {
		currentUser := user.(*Entities.Triage)
		fmt.Println("Fristname : %s , Lastname : %s , Password : %s , ID : %s , Age : %d", currentUser.FirstName, currentUser.LastName, currentUser.Password, currentUser.ID, currentUser.Age)
		fmt.Println("Do you want to change your data ?")
		fmt.Println("[1] Yes")
		fmt.Println("[2] No")
		cmd, _ := reader.ReadString('\n')
		cmd = cmd[:len(cmd)-1]
		if cmd == "2" {
			return
		}
		args := edit_form()
		if args[0] != "/" {
			currentUser.FirstName = args[0]
		} else if args[1] != "/" {
			currentUser.LastName = args[1]
		} else if args[2] != "/" {
			currentUser.Password = args[2]
		} else if args[3] != "/" {
			currentUser.ID = args[3]
		} else if args[4] != "/" {
			age, _ := strconv.Atoi(args[4])
			currentUser.Age = age
		}

	} else if our_type == reflect.TypeOf(&Entities.Patient{}) {
		currentUser := user.(*Entities.Patient)
		fmt.Println("Fristname : %s , Lastname : %s , Password : %s , ID : %s , Age : %d", currentUser.FirstName, currentUser.LastName, currentUser.Password, currentUser.ID, currentUser.Age)
		fmt.Println("Do you want to change your data ?")
		fmt.Println("[1] Yes")
		fmt.Println("[2] No")
		cmd, _ := reader.ReadString('\n')
		cmd = cmd[:len(cmd)-1]
		if cmd == "2" {
			return
		}
		args := edit_form()
		if args[0] != "/" {
			currentUser.FirstName = args[0]
		} else if args[1] != "/" {
			currentUser.LastName = args[1]
		} else if args[2] != "/" {
			currentUser.Password = args[2]
		} else if args[3] != "/" {
			currentUser.ID = args[3]
		} else if args[4] != "/" {
			age, _ := strconv.Atoi(args[4])
			currentUser.Age = age
		}

	}
}

func Manager_menu() int {
	// clear()
	fmt.Println("Manager Menu")
	fmt.Println("[1] Add employee")
	fmt.Println("[2] See signup Queue")
	fmt.Println("[3] Delete employee")
	fmt.Println("[4] Exit")

	reader := bufio.NewReader(os.Stdin)
	cmd, _ := reader.ReadString('\n')
	cmd = cmd[:len(cmd)-1] // Remove the trailing newline character
	Intcmd, _ := strconv.Atoi(cmd)
	clear()
	return Intcmd
}

func Manager_add_employee(manager Entities.Manager, DB *DataStructures.HashMap) {
	user, _, _, _, _, _, _, _ := signup_form()
	Auth.SignupEntity(user, *DB)

}

func Manager_delete_employee(manager Entities.Manager, DB *DataStructures.HashMap) {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("==Delete employee==")
	fmt.Print("Please enter employee National ID: ")

	nationalID, _ := reader.ReadString('\n')
	nationalID = nationalID[:len(nationalID)-1]

	DB.DeleteRecursive(nationalID)

}

func Manager_add_entity(manager Entities.Manager, DB *DataStructures.HashMap) {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("==Wait to add an entity==")
	fmt.Println("--------------------------------------------------------------")
	manager.ToAddStack.PrintByPoppingCopy()
	fmt.Println("--------------------------------------------------------------")

	fmt.Println("Please enter your choice:")
	fmt.Println("[1] Start checking emplyees")
	fmt.Println("[2] Exit")

	cmd, _ := reader.ReadString('\n')
	cmd = cmd[:len(cmd)-1] // Remove the trailing newline character
	Intcmd, _ := strconv.Atoi(cmd)
	clear()
	if Intcmd == 1 {
		fmt.Println("Please verify the entitys name by enter 1 or enter 0 to finish")
		for !manager.ToAddStack.IsEmpty() {
			entity, _ := manager.ToAddStack.Pop()
			fmt.Println(entity)
			Continue, _ := reader.ReadString('\n')
			Continue = Continue[:len(Continue)-1]
			if Continue == "0" {
				continue
			}
			if Continue == "1" {
				Auth.SignupEntity(entity, *DB)
			}

		}
	} else if Intcmd == 2 {
		return
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

	TriageList := DataStructures.LinkedList{}
	// ToAddStack := DataStructures.Stack{}
	ToAddStack := DataStructures.Stack{}

	DataBase := DataStructures.NewHashMap(100)
	DoctorsDB := DataStructures.NewHashMap(100)
	CardiologyDB := DataStructures.NewHashMap(100)
	TriageDB := DataStructures.NewHashMap(100)
	EmergencyDB := DataStructures.NewHashMap(100)
	DoctorsDB.Insert("Cardiology", CardiologyDB)
	DoctorsDB.Insert("Emergency", EmergencyDB)
	PatientsDB := DataStructures.NewHashMap(100)
	ManagerDB := DataStructures.NewHashMap(100)
	DrugManDB := DataStructures.NewHashMap(100)
	DataBase.Insert("Doctors", DoctorsDB)
	DataBase.Insert("Patients", PatientsDB)
	DataBase.Insert("Managers", ManagerDB)
	DataBase.Insert("DrugMans", DrugManDB)
	DataBase.Insert("Triages", TriageDB)

	drugBST := DataStructures.NewDrugBST()
	// testDrugSearch(drugBST)
	Auth.Signup("121", "Alireza", "Sobhdoost", "121", []string{"Manager"}, 19, *DataBase)
	// man , _ Auth.Login(*DataBase , "121","121")
	
	cmd := greet()
	clear()
	for true {

		if cmd == 1 {

			user, NID, FirstName, Lastname, password, args, age, isPatient := signup_form()
			if isPatient {
				err := Auth.Signup(NID, FirstName, Lastname, password, args, age, *DataBase)
				clear()
				if err != nil {
					log.Fatalf("Error setting password for doctor: %v", err)
				}
			} else {
				ToAddStack.Push(user)
			}

		} else if cmd == 2 {

			NID, password := login_form()
			user, err := Auth.Login(*DataBase, NID, password)
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
							p, _ := currentUser.VisitQueue.Pop()
							patient := p.(Entities.Patient)
							visit(patient)
							currentUser.PatientList.AddToStart(patient)
							d, _ := patient.DrugAllergies.Peek()
							fmt.Println("Drugs : ", d)
							fmt.Println(patient.DrugAllergies)
							// currentUser.PatientList.Display()

						} else if choice == 2 {
							// currentUser.PatientList.Display()
							clear()

							fmt.Println("==Doctor Patient List==")
							fmt.Println("---------------------------------------------")
							Entities.DisplayPatList(currentUser.PatientList)
							fmt.Println("---------------------------------------------")
							fmt.Println("press any Key to continue")
							cmd_reader := bufio.NewReader(os.Stdin)
							cmd_reader.ReadString('\n')

							clear()

						} else if choice == 3 {
							edit_menu(currentUser)
						} else if choice == 4 {
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
							clinicDBInterface, _ := DoctorsDB.GetRecursive(clinic)
							clinicDB := clinicDBInterface.(*DataStructures.HashMap)
							choose_doc(*currentUser, clinicDB)

						} else if inner_cmd == 2 {
							currentUser.DoctorList.Display()
							cancel_appointment(*currentUser)

						} else if inner_cmd == 4 {
							edit_menu(currentUser)
						} else if inner_cmd == 3 {
							innercmd := DrugStore_menu_Patient()
							for true {
								if innercmd == 1 {
									searchDrugByID(drugBST)
								} else if innercmd == 2 {
									searchDrugByName(drugBST)
								} else if innercmd == 3 {
									searchDrugByType(drugBST)
								} else if innercmd == 4 {
									searchByPriceRange(drugBST)
								} else if innercmd == 5 {
									break
								}
								innercmd = DrugStore_menu_Patient()
							}	
						} else if inner_cmd == 5 {
							// back = true
							break
						}
						inner_cmd = Patient_menu()
					}

				} else if our_type == reflect.TypeOf(&Entities.Manager{}) {
					currentUser := user.(*Entities.Manager)
					// ToAddStack.PrintByPoppingCopy()
					currentUser.ToAddStack = &ToAddStack
					innerCMD := Manager_menu()
					for true {
						if innerCMD == 1 {
							Manager_add_employee(*currentUser, DataBase)
						} else if innerCMD == 2 {
							Manager_add_entity(*currentUser, DataBase)

						} else if innerCMD == 3 {
							Manager_delete_employee(*currentUser, DataBase)
						} else if innerCMD == 4 {
							break
						} else if innerCMD == 5 {
							break
						}
						innerCMD = Manager_menu()

					}

				} else if our_type == reflect.TypeOf(&Entities.DrugMan{}) {
					currentUser := user.(*Entities.DrugMan)
					innercmd := DrugStore_menu()
					for true {
						if innercmd == 1 {
							DrugStore_Csevent(*currentUser, DataBase , drugBST)
						} else if innercmd == 2 {
							addNewDrug(drugBST)
						} else if innercmd == 3 {
							searchDrugByID(drugBST)
						} else if innercmd == 4 {
							searchDrugByName(drugBST)
						} else if innercmd == 5 {
							searchDrugByType(drugBST)
						} else if innercmd == 6 {
							searchByPriceRange(drugBST)
						} else if innercmd == 7 {
							cheapest := drugBST.MinHeap.ExtractMin()
							if cheapest != nil {
								fmt.Printf("Cheapest Drug: ID=%s, Name=%s, Price=%.2f, Type=%s, Dose=%s, Count=%d\n" ,
									cheapest.ID, cheapest.Name, cheapest.Price, cheapest.Type, cheapest.Dose, cheapest.Count)
							} else {
								fmt.Println("No drugs in inventory")
							}
						} else if innercmd == 8 {
							cheapest := drugBST.MaxHeap.ExtractMax()
							if cheapest != nil {
								fmt.Printf("Most Expensive Drug : ID=%s, Name=%s, Price=%.2f, Type=%s , Dose : %s Count : %d\n",
									cheapest.ID, cheapest.Name, cheapest.Price, cheapest.Type , cheapest.Dose , cheapest.Count)
							} else {
								fmt.Println("No drugs in inventory")
							}
						} else if innercmd == 9 {
							deleteDrug(drugBST)
						} else if innercmd == 10 {
							seeAllDrugs(drugBST)
						} else if innercmd == 11 {
							edit_menu(currentUser)
						} else if innercmd == 12 {
							break
						} else if innercmd == 13 {
							getTotalDrugCount(drugBST)
						} else if innercmd == 14 {
							getBSTDepth(drugBST)
						}
						innercmd = DrugStore_menu()
					}
				} else if our_type == reflect.TypeOf(&Entities.Triage{}) {

					currentUser := user.(*Entities.Triage)
					innerCmd := Triage_menu()
					for true {
						if innerCmd == 1 {
							Triage_agent(*EmergencyDB, TriageList)
						} else if innerCmd == 2 {
							edit_menu(currentUser)
						} else if innerCmd == 3 {
							break
						}
						innerCmd = Triage_menu()

					}

				}

				// break
				// continue
			}
		} else if cmd == 3 {

			Triage_entry(*DataBase, &TriageList)
			// TriageList.Display()

		} else if cmd == 4 {

			break
		}
		cmd = greet()
		clear()
	}

}
