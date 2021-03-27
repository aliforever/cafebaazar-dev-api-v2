package cafebazaar_dev_api_v2

const (
	authorizeUrl               = "https://pardakht.cafebazaar.ir/devapi/v2/auth/authorize/?response_type=code&access_type=offline&redirect_uri=%s&client_id=%s"
	authorizeTokenUrl          = "https://pardakht.cafebazaar.ir/devapi/v2/auth/authorize/"
	tokenUrl                   = "https://pardakht.cafebazaar.ir/devapi/v2/auth/token/"
	refreshTokenUrl            = "https://pardakht.cafebazaar.ir/devapi/v2/auth/token/"
	inAppPurchaseValidationUrl = "https://pardakht.cafebazaar.ir/devapi/v2/api/validate/%s/inapp/%s/purchases/%s/"
)
