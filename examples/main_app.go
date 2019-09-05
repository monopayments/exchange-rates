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
		AddExchanger(mex.NewyahooAPI(nil)).
		Build()

	usdToTryRate := SwapTest.Latest("USD/TRY")
	fmt.Println(usdToTryRate.GetRateValue())
	fmt.Println(usdToTryRate.GetRateDateTime())
	fmt.Println(usdToTryRate.GetExchangerName())
}
