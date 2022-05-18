package keeper

import (
	"github.com/ChronicNetwork/cht/x/cht/types"
)

var _ types.QueryServer = Keeper{}
