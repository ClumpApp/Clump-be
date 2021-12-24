package service

import (
	"io"
	"strings"

	"github.com/clumpapp/clump-be/utility"
)

func (obj *Service) uploadMedia(name string, data io.ReadSeekCloser) string {
	dot := strings.LastIndex(name, ".")
	extension := name[dot+1:]
	newName := utility.GetNameGen().RandStringBytesMaskImprSrcSB() + extension
	utility.GetStorage().Upload(newName, data)
	return newName
}

func (obj *Service) deleteMedia(name string) {
	utility.GetStorage().Delete(name)
}
