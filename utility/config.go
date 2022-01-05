package utility

import (
	"os"
	"sync"
)

const (
	azureStorageAccountConnStr = "CUSTOMCONNSTR_ASA"
	jwtEncryptionKey           = "APPSETTING_JWT"
	postgreSQLConnStr          = "POSTGRESQLCONNSTR_AD"
)

type config struct {
	db      string
	storage string
	jwt     string
}

var onceConfig sync.Once
var instanceConfig *config

func GetConfig() *config {
	onceConfig.Do(func() {
		db, _ := os.LookupEnv(postgreSQLConnStr)
		storage, _ := os.LookupEnv(azureStorageAccountConnStr)
		key, _ := os.LookupEnv(jwtEncryptionKey)

		instanceConfig = &config{db, storage, key}
	})

	return instanceConfig
}

func (obj config) GetDB() string {
	return obj.db
}

func (obj config) GetStorage() string {
	return obj.storage
}

func (obj config) GetJWTKey() string {
	return obj.jwt
}
