package auth

import (
	"fmt"
	"github.com/waykiss/wkcomps/database"
	"log"
	"os"
)

type Dao interface {
	Create(v Model) error
	Update(v Model) error
	Delete(v Model) error
	Find(v Query) ([]Model, error)
	FindByEmail(email string) (Model, error)
}

//singleton instace of the database
var db *database.Database

type DaoDatabase struct{}

//GetDb return the instance of database connection
func GetDb() (r *database.Database) {
	var err error
	if db == nil {
		db, err = database.NewConnection(os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"))
		if err != nil {
			log.Fatal("database connection error")
		}
		err = db.Migrate(&Model{})
	}
	r = db
	return
}

func NewDao() (Dao, error) {
	return DaoDatabase{}, nil
}

func (d DaoDatabase) Create(v Model) error {
	db := GetDb()
	return db.Create(&v)
}

func (d DaoDatabase) Update(v Model) error {
	return nil
}

func (d DaoDatabase) Delete(v Model) error {
	db := GetDb()
	_, err := db.Delete(&v)
	return err
}

func (d DaoDatabase) Find(_ Query) (r []Model, err error) {
	db := GetDb()
	err = db.Find(&r)
	return
}

func (d DaoDatabase) FindByEmail(email string) (r Model, err error) {
	db := GetDb()
	db.GormInstance().Where(fmt.Sprintf("%s = ?", dbEmailField), email).First(&r)
	return
}
