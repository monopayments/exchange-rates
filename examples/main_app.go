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
		AddExchanger(mex.NewGoogleApi(nil)).
		AddExchanger(mex.NewYahooApi(nil)).
		Build()

	euroToUsdRate := SwapTest.Latest("EUR/USD")
	fmt.Println(euroToUsdRate.GetRateValue())
	fmt.Println(euroToUsdRate.GetRateDateTime())
	fmt.Println(euroToUsdRate.GetExchangerName())
}
