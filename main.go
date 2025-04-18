package main

import (
	"fmt"
	"structure-phonebook-cli/handlers"
	"structure-phonebook-cli/storage"
	"structure-phonebook-cli/utils"
)

func main() {
	for {
		fmt.Println("PHONEBOOK CLI")
		fmt.Println("1. Add Contact")
		fmt.Println("2. View Contacts")
		fmt.Println("3. Search Contact")
		fmt.Println("4. Exit")

		choice := utils.GetInput("Enter your choice: ")

		switch choice {
		case "1":
			name := utils.GetInput("Name: ")
			phone := utils.GetInput("Phone: ")
			email := utils.GetInput("Email: ")
			handlers.AddNewContact(name, phone, email)
			storage.JsonConvert()

		case "2":
			//handlers.ShowAllContacts()
			storage.Readfile()

		case "3":
			name := utils.GetInput("Enter name to search: ")
			handlers.SearchContact(name)

		case "4":
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println(" Invalid choice")
		}
	}
}
