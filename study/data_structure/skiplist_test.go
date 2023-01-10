package data_structure

import (
	"math/rand"
	"testing"
	"time"
)

func TestNewSkipList(t *testing.T) {
	skipList := NewSkipList()

	rand.Seed(time.Now().Unix())
	for i := 0; i < 25; i++ {
		n := rand.Int63n(100)
		skipList.Insert(n, int(n))
	}

	skipList.Print()
}
