package database

import (
	"github.com/clumpapp/clump-be/model"

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

func (obj *Database) Migrate() {
	obj.gormDB.AutoMigrate(
		&model.User{},
		&model.Group{},
		&model.UserGroups{},
		&model.Interest{},
		&model.Message{},
	)
}

func (obj *Database) Create(model, objects interface{}) {
	obj.gormDB.Model(model).Create(objects)
}

func (obj *Database) Read(model, ID, object interface{}) {
	obj.gormDB.Model(model).Where(ID).First(object)
}

func (obj *Database) Query(model, query, objects interface{}) {
	obj.gormDB.Model(model).Where(query).Find(objects)
}

func (obj *Database) Update(model, query, object interface{}) {
	obj.gormDB.Model(model).Where(query).Updates(object)
}

func (obj *Database) Delete(model, query interface{}) {
	obj.gormDB.Where(query).Delete(model)
}
