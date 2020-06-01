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

package params

import (
	"math/big"
	"reflect"
	"testing"
)

func TestCheckCompatible(t *testing.T) {
	type test struct {
		stored, new *ChainConfig
		head        uint64
		wantErr     *ConfigCompatError
	}
	tests := []test{
		{stored: AllDavinciProtocolChanges, new: AllDavinciProtocolChanges, head: 0, wantErr: nil},
		{stored: AllDavinciProtocolChanges, new: AllDavinciProtocolChanges, head: 100, wantErr: nil},
		{
			stored:  &ChainConfig{},
			new:     &ChainConfig{AresBlock: big.NewInt(20)},
			head:    9,
			wantErr: nil,
		},
		{
			stored:  &ChainConfig{ChainId:big.NewInt(1), ByzantiumBlock: big.NewInt(4370000), AresBlock: big.NewInt(373300)},
			new:     &ChainConfig{ChainId:big.NewInt(1), ByzantiumBlock: big.NewInt(4370000), AresBlock: big.NewInt(373300), EpiphronBlock: big.NewInt(400000)},
			head:    390000,
			wantErr: nil,
		},
		{
			stored:  &ChainConfig{ChainId:big.NewInt(1), ByzantiumBlock: big.NewInt(4370000), AresBlock: big.NewInt(373300), BlockInterval:big.NewInt(10)},
			new:     &ChainConfig{ChainId:big.NewInt(1), ByzantiumBlock: big.NewInt(4370000), AresBlock: big.NewInt(373300), BlockInterval:big.NewInt(5)},
			head:    390000,
			wantErr: nil,
		},
		{
			stored: AllDavinciProtocolChanges,
			new:    &ChainConfig{},
			head:   3,
			wantErr: &ConfigCompatError{
				What:         "Epiphron fork block",
				StoredConfig: big.NewInt(0),
				NewConfig:    nil,
				RewindTo:     0,
			},
		},
		{
			stored: AllDavinciProtocolChanges,
			new:    &ChainConfig{ByzantiumBlock: big.NewInt(10000),AresBlock:big.NewInt(90), EpiphronBlock:big.NewInt(1)},
			head:   3,
			wantErr: &ConfigCompatError{
				What:         "Epiphron fork block",
				StoredConfig: big.NewInt(0),
				NewConfig:    big.NewInt(1),
				RewindTo:     0,
			},
		},
	}

	for _, test := range tests {
		err := test.stored.CheckCompatible(test.new, test.head)
		if !reflect.DeepEqual(err, test.wantErr) {
			t.Errorf("error mismatch:\nstored: %v\nnew: %v\nhead: %v\nerr: %v\nwant: %v", test.stored, test.new, test.head, err, test.wantErr)
		}
	}
}
