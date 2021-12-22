package database

import (
	"github.com/clumpapp/clump-be/model"
	"github.com/clumpapp/clump-be/utility"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	gormDB *gorm.DB
}

func New() *Database {
	return &Database{}
}

func (obj *Database) Connect() {
	dsn := utility.GetConfig().GetDB()
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	obj.gormDB = db
}

func (obj *Database) Migrate() {
	obj.gormDB.AutoMigrate(
		&model.Group{},
		&model.User{},
		&model.Interest{},
		&model.Message{},
		&model.IEUserGroup{},
		&model.IEUserInterest{},
		&model.IEGroupInterest{},
	)
}

func (obj *Database) Create(model, objects interface{}) {
	obj.gormDB.Model(model).Create(objects)
}

func (obj *Database) Read(model, ID, object interface{}) bool {
	res := obj.gormDB.Model(model).Where(ID).First(object)
	return res.RowsAffected != 0
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
