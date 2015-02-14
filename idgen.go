// serial id gen package
package idgen

import (
	crand "crypto/rand"
	"fmt"
	"sort"
)

type UUID [16]uint8

func (u UUID) String() string {
	return fmt.Sprintf("%x-%x-%x-%x-%x",
		u[:4], u[4:6], u[6:8], u[8:10], u[10:])
}
func NewUUID() *UUID {
	u := UUID{}
	crand.Read(u[:])
	return &u
}

type IDList []int64

func (s IDList) Len() int {
	return len(s)
}
func (s IDList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s IDList) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s IDList) FindIndex(id int64) int {
	return sort.Search(len(s), func(i int) bool { return s[i] >= id })
}

var genCh chan int64

func init() {
	genCh = make(chan int64)
	go func() {
		var i int64
		for {
			i++
			genCh <- i
		}
	}()
}

func GenCh() <-chan int64 {
	return genCh
}
