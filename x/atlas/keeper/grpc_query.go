package keeper

import (
	"github.com/cosmonaut/Atlas/x/atlas/types"
)

var _ types.QueryServer = Keeper{}
