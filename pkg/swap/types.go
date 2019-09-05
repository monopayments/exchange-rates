package swap

import (
	mex "github.com/monocash/exchange-rates/pkg/exchanger"
)

// Swap ... main struct
type Swap struct {
	exchangers []mex.Exchanger
	cache      interface{}
}
