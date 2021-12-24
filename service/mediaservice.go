package service

import (
	"github.com/clumpapp/clump-be/utility"
)

func (obj *Service) CreateMedia(extension string, data []byte) string {
	// Determine file format, then assign a 16 char random name and upload
	name := "asd" + extension
	utility.GetStorage().Upload(name, data)
	name = utility.GetStorage().GetURL() + name
	return name
}

func (obj *Service) DeleteMedia(name string) {
	utility.GetStorage().Delete(name)
}
