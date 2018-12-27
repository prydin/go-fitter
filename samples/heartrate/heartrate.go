package main

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/prydin/go-fitter/client"
	"github.com/prydin/go-fitter/client/auth"

	"github.com/go-openapi/strfmt"
)

func authenticate(ac *auth.Client, clientID, clientSecret, clientCode string) (*auth.OauthTokenOK, error) {
	params := auth.NewOauthTokenParams()
	params.ClientID = clientID
	params.Code = &clientCode
	a := "Basic " + base64.StdEncoding.EncodeToString([]byte(clientID+":"+clientSecret))
	params.Authorization = &a
	params.GrantType = "authorization_code"
	fmt.Println(params)
	return ac.OauthToken(params)
}

func main() {
	clientID := os.Getenv("FITBIT_CLIENT_ID")
	clientSecret := os.Getenv("FITBIT_CLIENT_SECRET")
	clientCode := os.Getenv("FITBIT_CLIENT_CODE")
	c := client.NewHTTPClient(strfmt.NewFormats())
	t, err := authenticate(c.Auth, clientID, clientSecret, clientCode)
	if err != nil {
		fmt.Printf(*t)
		panic(err)
	}
}
