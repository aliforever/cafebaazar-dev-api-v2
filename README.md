# cafebaazar-dev-api-v2
Go [CafeBazaar Developer API Version 2](https://developers.cafebazaar.ir/en/docs/developer-api-v2-introduction/developer-api-v2-getting-started/)

## Install:
```go get -u github.com/aliforever/cafebaazar-dev-api-v2```

---

## Usage:
### 1) To initialise the library use: 
```go
api := cafebazaar_dev_api_v2.NewCafeBazaarAPI("clientId", "clientSecret", "clientUri")
```
Fill in the parameters clientId, clientSecret, clientUri, set in your [api section of developer console](https://pishkhan.cafebazaar.ir/settings/api)

### 2) You will need a code to get access and refresh tokens, you can get the code by visiting the url:
`https://pardakht.cafebazaar.ir/devapi/v2/auth/authorize/?response_type=code&access_type=offline&redirect_uri=<REDIRECT_URI>&client_id=<CLIENT_ID>`

Or calling `Authorize` method of the API:
```go
code, err := api.Authorize("androidpublisher")
if err != nil {
  fmt.Println(err)
  return
}
```

### 3) To get accessToken & refreshToken you should pass the code received in the second step to `GetAuthorizationTokens`:
```go
token, err := api.GetAuthorizationTokens(code)
if err != nil {
  fmt.Println(err)
  return
}
// This will store the token inside the struct for future use, or you can store the tokens and manually import them
```

To manually import the tokens you can use `SetToken`:
```go
accessToken := ""
refreshToken := ""
api.SetToken(accessToken, refreshToken)
```

---
## Supported Functions:
#### [In-app Purchase Validation](https://developers.cafebazaar.ir/en/docs/developer-api-v2-introduction/developer-api-v2-ref-validate/)
```go
iapv, err := api.InAppPurchaseValidate("packageName", "productId", "purchaseToken")
if err != nil {
  fmt.Println(err)
  return
}
fmt.Printf("%+v", iapv)
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
```

TODO: 
### [Subscription Validation](https://developers.cafebazaar.ir/en/docs/developer-api-v2-introduction/developer-api-v2-ref-get-subs/)
### [Subscription Cancellation](https://developers.cafebazaar.ir/en/docs/developer-api-v2-introduction/developer-api-v2-ref-cancel-subs/)

Pull requests are welcome for the missing functions.
