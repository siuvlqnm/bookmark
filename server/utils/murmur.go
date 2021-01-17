package utils

import (
	"github.com/spaolacci/murmur3"
)

func GetMurmur32(str string) uint32 {
	return murmur3.Sum32([]byte(str))
}
