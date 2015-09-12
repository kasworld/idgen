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
	"testing"
)

func TestIDGen(t *testing.T) {
	idfn := NewGen32(22)
	idfn2 := NewGen32(42)
	for i := 0; i < 100; i++ {
		t.Logf("%v %v %v", <-idfn(), <-idfn2(), <-ID32GenCh())
	}
}
