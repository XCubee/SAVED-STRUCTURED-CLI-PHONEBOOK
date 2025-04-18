package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"structure-phonebook-cli/models"
)

var Contacts []models.Contact

func AddContact(contact models.Contact) {
	Contacts = append(Contacts, contact)

}
func GetAllContacts() []models.Contact {
	return Contacts
}

func FindContactByName(name string) *models.Contact {
	for i, contact := range Contacts {
		if contact.Name == name {
			return &Contacts[i]
		}
	}
	return nil
}

func JsonConvert() {
	jsonfile, err := json.MarshalIndent(Contacts, "", "")
	if err != nil {
		fmt.Println("Error in Converting")
	}

	err = os.WriteFile("data/jsonfile.json", jsonfile, 0644)
	if err != nil {
		fmt.Println("Error in writing:", err)
		return
	}
	fmt.Println("Contacts saved to JSON file.")

}

func Readfile() {
	readdata, err := os.ReadFile("data/jsonfile.json")
	if err != nil {
		fmt.Println("Error in reading")
	}
	fmt.Println(string(readdata))
}
