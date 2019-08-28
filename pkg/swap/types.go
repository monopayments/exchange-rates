package swap

import (
	mex "github.com/monocah/exchange-rates/pkg/exchanger"
)

// Swap ... main struct
type Swap struct {
	exchangers []mex.Exchanger
	cache      interface{}
}
