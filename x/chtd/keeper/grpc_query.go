package keeper

import (
	"github.com/ChronicNetwork/chtd/x/chtd/types"
)

var _ types.QueryServer = Keeper{}
