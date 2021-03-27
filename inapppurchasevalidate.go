package cafebazaar_dev_api_v2

type InAppPurchaseValidate struct {
	ConsumptionState int    `json:"consumptionState"`
	PurchaseState    int    `json:"purchaseState"`
	Kind             string `json:"kind"`
	DeveloperPayload string `json:"developerPayload"`
	PurchaseTime     int64  `json:"purchaseTime"`
	Error            string `json:"error,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}
