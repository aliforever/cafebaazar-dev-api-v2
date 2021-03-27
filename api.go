package cafebazaar_dev_api_v2

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"
)

type CafeBazaarAPI struct {
	clientId           string
	clientSecret       string
	clientUri          string
	http               *http.Client
	lastToken          *Token
	lastTokenExpiresAt int64
}

func NewCafeBazaarAPI(clientId, clientSecret, clientUri string) *CafeBazaarAPI {
	j, _ := cookiejar.New(nil)
	c := &http.Client{}
	c.Jar = j
	return &CafeBazaarAPI{
		clientId:     clientId,
		clientSecret: clientSecret,
		clientUri:    clientUri,
		http:         c,
	}
}

func (ca *CafeBazaarAPI) LastToken() (token *Token) {
	return ca.lastToken
}

func (ca *CafeBazaarAPI) SetToken(accessToken, refreshToken string) {
	ca.lastToken = &Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    0,
		TokenType:    "Bearer",
		Scope:        "androidpublisher",
	}
}

func (ca *CafeBazaarAPI) Authorize(scope string) (token string, err error) {
	if scope == "" {
		scope = "androidpublisher"
	}
	address := fmt.Sprintf(authorizeUrl, ca.clientUri, ca.clientId)
	var resp *http.Response
	resp, err = ca.http.Get(address)
	if err != nil {
		return
	}
	b, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	data := string(b)
	search := `<form method="POST" action="."><input type='hidden' name='csrfmiddlewaretoken' value='`
	index := strings.Index(data, search)
	if index == -1 {
		err = errors.New("invalid_response_" + data)
		return
	}
	data = data[index+len(search):]
	csrfToken := strings.TrimSpace(data[:strings.Index(data, "'")])
	v := url.Values{}
	v.Set("csrfmiddlewaretoken", csrfToken)
	v.Set("scopes", scope)
	req, _ := http.NewRequest("POST", authorizeTokenUrl, strings.NewReader(v.Encode()))
	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	req.Header.Set("referer", address)
	resp, err = ca.http.Do(req)
	if err != nil {
		return
	}

	search = "?code="
	index = strings.Index(resp.Request.URL.String(), search)
	if index == -1 {
		b, _ = ioutil.ReadAll(resp.Body)
		err = errors.New("cant_authorize_" + string(b))
		resp.Body.Close()
		return
	} else {
		token = resp.Request.URL.String()[index+len(search):]
	}
	return
}

func (ca *CafeBazaarAPI) GetAuthorizationTokens(code string) (token *Token, err error) {
	var resp *http.Response

	v := url.Values{}
	v.Set("grant_type", "authorization_code")
	v.Set("code", code)
	v.Set("client_id", ca.clientId)
	v.Set("client_secret", ca.clientSecret)
	v.Set("redirect_uri", ca.clientUri)

	resp, err = ca.http.Post(tokenUrl, "application/x-www-form-urlencoded", strings.NewReader(v.Encode()))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	j := json.NewDecoder(resp.Body)
	err = j.Decode(&token)
	if err != nil {
		return
	}
	if token.Error != "" {
		err = errors.New(token.Error)
		token = nil
		return
	}
	ca.lastToken = token
	ca.lastTokenExpiresAt = time.Now().Unix() + ca.lastToken.ExpiresIn
	return
}

func (ca *CafeBazaarAPI) RefreshToken() (err error) {
	if ca.lastToken == nil {
		err = errors.New("empty_refresh_token")
		return
	}

	var resp *http.Response

	v := url.Values{}
	v.Set("grant_type", "refresh_token")
	v.Set("client_id", ca.clientId)
	v.Set("client_secret", ca.clientSecret)
	v.Set("refresh_token", ca.lastToken.RefreshToken)

	resp, err = ca.http.Post(refreshTokenUrl, "application/x-www-form-urlencoded", strings.NewReader(v.Encode()))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var token *Token
	j := json.NewDecoder(resp.Body)
	err = j.Decode(&token)
	if err != nil {
		return
	}
	if token.Error != "" {
		err = errors.New(token.Error)
		token = nil
		return
	}
	ca.lastToken = token
	ca.lastTokenExpiresAt = time.Now().Unix() + ca.lastToken.ExpiresIn
	return
}

func (ca *CafeBazaarAPI) InAppPurchaseValidate(packageName, productId, purchaseToken string) (response *InAppPurchaseValidate, err error) {
	if time.Now().Unix() < ca.lastTokenExpiresAt {
		err = ca.RefreshToken()
		if err != nil {
			return
		}
	}

	address := fmt.Sprintf(inAppPurchaseValidationUrl, packageName, productId, purchaseToken)

	req, _ := http.NewRequest("GET", address, nil)
	req.Header.Set("Authorization", ca.lastToken.header())

	var resp *http.Response
	resp, err = ca.http.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	j := json.NewDecoder(resp.Body)
	err = j.Decode(&response)
	if err != nil {
		return
	}
	if response.Error != "" {
		err = errors.New(response.Error + " => " + response.ErrorDescription)
		response = nil
		return
	}
	return
}
