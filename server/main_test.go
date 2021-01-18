package main_test

import (
	"testing"

	"github.com/spaolacci/murmur3"
)

func TestGetmurmur(T *testing.T) {
	n32 := murmur3.New32()
	T.Log(n32)
}
