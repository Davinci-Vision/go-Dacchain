package types

import (
	"github.com/Dacchain/go-Dacchain/common"
	"math/big"
	"github.com/Dacchain/go-Dacchain/common/hexutil"
)

//go:generate gencodec -type InnerTx -field-override innertxMarshaling -out gen_innertx_json.go

type InnerTx struct {
	From		common.Address	`json:"from" gencodec:"required"`
	To 			common.Address	`json:"to" gencodec:"to" gencodec:"required"`
	AssetID		*common.Address	`json:"assetid" rlp:"nil"`
	Value		*big.Int		`json:"value" gencodec:"required"`
}


type innertxMarshaling struct {
	Value		*hexutil.Big
}