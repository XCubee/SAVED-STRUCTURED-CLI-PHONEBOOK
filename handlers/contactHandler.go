package handlers

import (
	"fmt"
	"strings"
	"structure-phonebook-cli/models"
	"structure-phonebook-cli/storage"
)

func AddNewContact(name, phone, email string) {
	contact := models.Contact{Name: name, Phone: phone, Email: email}
	storage.AddContact(contact)
	fmt.Println("Contact added!")
}
func ShowAllContacts() {
	contacts := storage.GetAllContacts()
	if len(contacts) == 0 {
		fmt.Println("No contacts found.")
		return
	}
	for _, i := range contacts {
		fmt.Printf("Name %s - PhoneNuumber %s - Email %s\n", i.Name, i.Phone, i.Email)
	}
}

func SearchContact(name string) {
	contact := storage.FindContactByName(strings.ToLower(name))
	if contact != nil {
		fmt.Printf("Found:: Name %s - PhoneNumber %s - Email %s\n", contact.Name, contact.Phone, contact.Email)
	} else {
		fmt.Println("Contact not found.")
	}
}


