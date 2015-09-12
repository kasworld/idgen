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
)

type ID32 int32

func (id ID32) String() string {
	return fmt.Sprintf("0x%0X", int32(id))
}

func NewGen32(beginnum ID32) func() <-chan ID32 {
	genCh := make(chan ID32)
	go func(i ID32) {
		// var i ID32
		for {
			i++
			genCh <- i
		}
	}(beginnum)
	return func() <-chan ID32 {
		return genCh
	}
}

var ID32GenCh = NewGen32(0)
