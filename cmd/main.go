package main

import (
	"github.com/huobirdcenter/huobi_golang/cmd/accountclientexample"
	"github.com/huobirdcenter/huobi_golang/cmd/commonclientexample"
	"github.com/huobirdcenter/huobi_golang/cmd/crossmarginclientexample"
	"github.com/huobirdcenter/huobi_golang/cmd/etfclientexample"
	"github.com/huobirdcenter/huobi_golang/cmd/isolatedmarginclientexample"
	"github.com/huobirdcenter/huobi_golang/cmd/marketclientexample"
	"github.com/huobirdcenter/huobi_golang/cmd/orderclientexample"
	"github.com/huobirdcenter/huobi_golang/cmd/walletclientexample"
	"github.com/huobirdcenter/huobi_golang/logging/perflogger"
	"github.com/shopspring/decimal"
)

func main() {
	runAll()
}

// Run all examples
func runAll() {

	// 取消order
	// orderclientexample.CancelOrdersByCriteria()
	// btc
	rate := [5]float64{0.95, 0.9, 0.85, 0.8, 0.75}
	amount := [5]int32{8, 16, 24, 32, 40}
	payUsdt("btcusdt", rate, amount, 2, 6)

	// eth
	rate1 := [5]float64{0.95, 0.9, 0.85, 0.8, 0.75}
	amount1 := [5]int32{8, 16, 24, 32, 40}
	payUsdt("ethusdt", rate1, amount1, 2, 4)

	// mana
	rate2 := [5]float64{0.95, 0.9, 0.85, 0.8, 0.75}
	amount2 := [5]int32{10, 20, 30, 40, 50}
	payUsdt("manausdt", rate2, amount2, 2, 2)

	// shib
	// rate3 := [5]float64{0.95, 0.9, 0.85, 0.8, 0.75}
	// amount3 := [5]int32{8, 16, 24, 32, 10}
	// payUsdt("shibusdt", rate3, amount3, 7, 2)

	// sol
	// rate4 := [5]float64{0.95, 0.9, 0.85, 0.8, 0.75}
	// amount4 := [5]int32{8, 16, 24, 32, 10}
	// payUsdt("solusdt", rate4, amount4, 4, 2)

	// sandusdt
	rate5 := [5]float64{0.95, 0.9, 0.85, 0.8, 0.75}
	amount5 := [5]int32{8, 16, 24, 32, 40}
	payUsdt("sandusdt", rate5, amount5, 2, 4)

	// btc
	// var btcusdt string = "btcusdt"
	// btcusdtResp := marketclientexample.RunAllExamples(btcusdt)

	// btcClose, _ := btcusdtResp.Close.Float64()

	// btcPrice5, btcPrice10, btcPrice15, btcPrice20, btcPrice25 := btcClose * 0.95, btcClose * 0.9, btcClose * 0.85, btcClose * 0.8, btcClose * 0.75
	// btcAmount5, btcAmount10, btcAmount15, btcAmount20, btcAmount25 := 8 / btcClose, 12 / btcClose, 18 / btcClose, 24 / btcClose, 30 / btcClose * 0.75
	// var btcSymbool string = "btcusdt"

	// orderclientexample.RunAllExamples(btcSymbool, formattingFloatToString(btcPrice5, 2), formattingFloatToString(btcAmount5, 6))
	// orderclientexample.RunAllExamples(btcSymbool, formattingFloatToString(btcPrice10, 2), formattingFloatToString(btcAmount10, 6))
	// orderclientexample.RunAllExamples(btcSymbool, formattingFloatToString(btcPrice15, 2), formattingFloatToString(btcAmount15, 6))
	// orderclientexample.RunAllExamples(btcSymbool, formattingFloatToString(btcPrice20, 2), formattingFloatToString(btcAmount20, 6))
	// orderclientexample.RunAllExamples(btcSymbool, formattingFloatToString(btcPrice25, 2), formattingFloatToString(btcAmount25, 6))

	// var ethusdt string = "ethusdt"
	// ethusdtResp := marketclientexample.RunAllExamples(ethusdt)
	//
	// fmt.Println(ethusdtResp.Close)

	// var symbool string = "btcusdt"
	// var price string = "btcusdt"
	// var amount string = "btcusdt"
	// orderclientexample.RunAllExamples(symbool, price, amount)

	// 账户相关
	// accountclientexample.RunAllExamples()

	// algoorderclientexample.RunAllExamples()
	// commonclientexample.RunAllExamples()
	// isolatedmarginclientexample.RunAllExamples()
	// crossmarginclientexample.RunAllExamples()
	// walletclientexample.RunAllExamples()
	// subuserclientexample.RunAllExamples()
	// stablecoinclientexample.RunAllExamples()
	// etfclientexample.RunAllExamples()
	// marketwebsocketclientexample.RunAllExamples()
	// accountwebsocketclientexample.RunAllExamples()
	// orderwebsocketclientexample.RunAllExamples()
}

// Run performance test
func runPerfTest() {
	perflogger.Enable(true)

	commonclientexample.RunAllExamples()
	accountclientexample.RunAllExamples()
	// orderclientexample.RunAllExamples()
	// marketclientexample.RunAllExamples()
	isolatedmarginclientexample.RunAllExamples()
	crossmarginclientexample.RunAllExamples()
	walletclientexample.RunAllExamples()
	etfclientexample.RunAllExamples()
}

func payUsdt(usdtStr string, rate [5]float64, amount [5]int32, priceScale int32, mountScale int32) {
	usdtResp := marketclientexample.RunAllExamples(usdtStr)

	bClose, _ := usdtResp.Close.Float64()

	bPrice5, bPrice10, bPrice15, bPrice20, bPrice25 := bClose*rate[0], bClose*rate[1], bClose*rate[2], bClose*rate[3], bClose*rate[4]
	bAmount5, bAmount10, bAmount15, bAmount20, bAmount25 := float64(amount[0])/bClose, float64(amount[1])/bClose, float64(amount[2])/bClose, float64(amount[3])/bClose, float64(amount[4])/bClose

	orderclientexample.RunAllExamples(usdtStr, formattingFloatToString(bPrice5, priceScale), formattingFloatToString(bAmount5, mountScale))
	orderclientexample.RunAllExamples(usdtStr, formattingFloatToString(bPrice10, priceScale), formattingFloatToString(bAmount10, mountScale))
	orderclientexample.RunAllExamples(usdtStr, formattingFloatToString(bPrice15, priceScale), formattingFloatToString(bAmount15, mountScale))
	orderclientexample.RunAllExamples(usdtStr, formattingFloatToString(bPrice20, priceScale), formattingFloatToString(bAmount20, mountScale))
	orderclientexample.RunAllExamples(usdtStr, formattingFloatToString(bPrice25, priceScale), formattingFloatToString(bAmount25, mountScale))
}

// func doDiv(m int32, n float64) {
// 	return decimal.NewFromFloat(float64(m)).Div(decimal.NewFromFloat(n))
// }

func formattingFloatToString(num float64, scale int32) string {
	return decimal.NewFromFloat(num).Round(scale).String()
}
