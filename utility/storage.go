package utility

import (
	"io"
	"os"
	"path/filepath"
	"sync"
)

type storage struct {
	folder string
}

var onceStorage sync.Once
var instanceStorage *storage

func GetStorage() *storage {
	onceStorage.Do(func() {
		folder := GetConfig().GetStorage()

		instanceStorage = &storage{folder}
	})

	return instanceStorage
}

func (obj storage) Upload(name string, data io.ReadSeekCloser) {

	file := filepath.Join(obj.folder, name)
	osFile, err := os.Create(file)
	if err != nil {
		println(err)
		return
	}
	defer osFile.Close()
	io.Copy(osFile, data) // This is the reason I love Go

}

func (obj storage) Delete(name string) {

	file := filepath.Join(obj.folder, name)
	err := os.Remove(file)
	if err != nil {
		println(err)
	}

}
