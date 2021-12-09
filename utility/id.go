package utility

import (
	"log"
	"strconv"
)

func ConvertID(id string) uint {
	res, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		// Should do logging
		log.Panicln(err)
	}
	return uint(res)
}
