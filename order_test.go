package wxpay

import (
	"testing"
	"fmt"
)

func TestWXPay_UnifiedOrder(t *testing.T) {
	fmt.Println("========== UnifiedOrder ==========")
	var p = &UnifiedOrderParam{}
	p.Body = "test product"
	p.NotifyURL = "http://www.test.com"
	p.TradeType = K_TRADE_TYPE_APP
	p.SpbillCreateIP = "220.112.233.229"
	p.TotalFee = 10
	p.OutTradeNo = "test-1111"

	result, err := client.UnifiedOrder(p)
	if err != nil {
		fmt.Println(err)
		t.Fatal(err)
	}
	fmt.Println(result.PrepayId)
}