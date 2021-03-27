package main

import (
	"fmt"

	cafebazaar_dev_api_v2 "github.com/aliforever/cafebazaar-dev-api-v2"
)

func main() {
	api := cafebazaar_dev_api_v2.NewCafeBazaarAPI("", "", "")
	code, err := api.Authorize("")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	token, err := api.GetAuthorizationTokens(code)
	if err != nil {
		fmt.Println(err)
		return
	}
	// err = api.RefreshToken()
	// fmt.Println(err)
	fmt.Println("code is: ", code)
	fmt.Println("lastToken: ", token)
}
