package utility

import (
	"math/rand"
	"strings"
	"sync"
	"time"
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

type namegen struct {
	src rand.Source
}

var onceNameGen sync.Once
var instanceNameGen *namegen

func GetNameGen() *namegen {
	onceNameGen.Do(func() {
		src := rand.NewSource(time.Now().UnixNano())

		instanceNameGen = &namegen{src}
	})

	return instanceNameGen
}

func (obj namegen) GenerateName() string {
	sb := strings.Builder{}
	sb.Grow(16)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := 15, obj.src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = obj.src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}
