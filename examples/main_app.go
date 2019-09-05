// +build ignore

package main

import (
	"fmt"
	mex "github.com/monocash/exchange-rates/pkg/exchanger"
	"github.com/monocash/exchange-rates/pkg/swap"
)

func main() {
	SwapTest := swap.NewSwap()

	SwapTest.
		AddExchanger(mex.NewinvestingAPI(nil)).
		Build()

	euroToUsdRate := SwapTest.Latest("USD/TRY")
	fmt.Println(euroToUsdRate.GetRateValue())
	fmt.Println(euroToUsdRate.GetRateDateTime())
	fmt.Println(euroToUsdRate.GetExchangerName())
}
