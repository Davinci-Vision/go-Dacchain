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
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"github.com/Dacchain/go-Dacchain/core/types"
	"github.com/Dacchain/go-Dacchain/log"
	"math"
	"strconv"
	"time"
)


func Shuffle(height int64, delegateNumber int) []int {
	var truncDelegateList []int

	for i := 0; i < delegateNumber; i++ {
		truncDelegateList = append(truncDelegateList, i)
	}

	seed := math.Floor(float64(height / int64(delegateNumber)))
	//seed := strconv.FormatFloat(math.Floor(float64(height/101)), 'E', -1, 64)

	if height%int64(delegateNumber) > 0 {
		seed += 1
	}
	seedSource := strconv.FormatFloat(seed, 'E', -1, 64)
	var buf bytes.Buffer
	buf.WriteString(seedSource)
	hash := sha256.New()
	hash.Write(buf.Bytes())
	md := hash.Sum(nil)
	currentSend := hex.EncodeToString(md)

	delCount := len(truncDelegateList)
	for i := 0; i < delCount; i++ {
		for x := 0; x < 4 && i < delCount; i++ {
			newIndex := int(currentSend[x]) % delCount
			truncDelegateList[newIndex], truncDelegateList[i] = truncDelegateList[i], truncDelegateList[newIndex]
			x++
		}
	}
	return truncDelegateList

}
