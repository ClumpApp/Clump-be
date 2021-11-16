package utility

import (
	"github.com/jinzhu/copier"
)

func Convert(in, out interface{}) {
	copier.Copy(out, in)
}
