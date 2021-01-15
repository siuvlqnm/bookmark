package utils

import "github.com/spaolacci/murmur3"

func GetMurmur128() (h1, h2 uint64) {
	return murmur3.New128().Sum128()
}
