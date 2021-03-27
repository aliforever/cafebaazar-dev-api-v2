package main

import (
	"fmt"

	cafebazaar_dev_api_v2 "github.com/aliforever/cafebazaar-dev-api-v2"
)

func main() {
	accessToken := ""
	refreshToken := ""
	api := cafebazaar_dev_api_v2.NewCafeBazaarAPI("", "", "")
	api.SetToken(accessToken, refreshToken)
	iapv, err := api.InAppPurchaseValidate("", "", "")
	if err != nil {
		fmt.Println(err)
		return
	}
	if iapv.ConsumptionState == 0 {
		fmt.Printf("consumed")
	} else {
		fmt.Printf("not consumed")
	}
	if iapv.PurchaseState == 0 {
		fmt.Printf("success")
	} else {
		fmt.Printf("refunded")
	}
	fmt.Printf("%+v", iapv)
}
