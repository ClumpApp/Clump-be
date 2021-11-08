package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	gormDB *gorm.DB
}

func New() *Database {
	return &Database{}
}

func (obj *Database) Connect() {
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	obj.gormDB = db
}

func (obj *Database) Initialize() {

}
