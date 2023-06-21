package models

import (
	"log"

	"github.com/dzoxploit/crud-contact-golang/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Contact struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Gender     string   `json:"gender"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}


func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Contact{})
}

func (b *Contact) CreateContact() *Contact {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllContact(pageNumber int, pageSize int, search string) ([]Contact, error) {


	// Calculate the offset based on the page number and page size
	offset := (pageNumber - 1) * pageSize

	// Query the database with pagination
	var contacts []Contact
	result := db.Where("name LIKE ?", "%"+search+"%").Offset(offset).Limit(pageSize).Find(&contacts)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return contacts, nil
}

// func GetAllContact() []Contact {
// 	var Contact []Contact
// 	db.Find(&Contact)
// 	return Contact
// }

func GetContactById(Id string) (*Contact , *gorm.DB){
	var getContact Contact
	db:=db.Where("ID = ?", Id).Find(&getContact)
	return &getContact, db
}

func DeleteContact(ID string) Contact {
	var contact Contact
	db.Where("ID = ?", ID).Delete(contact)
	return contact
}