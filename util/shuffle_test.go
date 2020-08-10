// Copyright 2018 The go-Dacchain Authors
// This file is part of the go-Dacchain library.
//
// The go-Dacchain library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-Dacchain library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-Dacchain library. If not, see <http://www.gnu.org/licenses/>.

package util

import (
	"fmt"
	"github.com/Dacchain/go-Dacchain/core/types"
	"testing"
	"time"
)

func TestShuffle(t *testing.T) {
	var delegateNumber = 4
	for i := 1; i < 100; i++ {
		fmt.Println(Shuffle(int64(i), delegateNumber))
		if i%delegateNumber == 0 {
			fmt.Println("=======================")
		}
	}
}
