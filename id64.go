// Copyright 2015 SeukWon Kang (kasworld@gmail.com)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// serial id gen package
package idgen

import (
	"fmt"
	"sort"
)

type IDList []IDInt

func (s IDList) Len() int {
	return len(s)
}
func (s IDList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s IDList) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s IDList) FindIndex(id IDInt) int {
	return sort.Search(len(s), func(i int) bool { return s[i] >= id })
}

type IDInt int64

func (id IDInt) String() string {
	return fmt.Sprintf("0x%0X", int64(id))
}

func NewGen(beginnum IDInt) func() <-chan IDInt {
	genCh := make(chan IDInt)
	go func(i IDInt) {
		for {
			i++
			genCh <- i
		}
	}(beginnum)
	return func() <-chan IDInt {
		return genCh
	}
}

var GenCh = NewGen(0)
