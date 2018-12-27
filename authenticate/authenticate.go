package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/prydin/go-fitter/client"
	"github.com/prydin/go-fitter/client/auth"
)

type handler struct {
	clientID     string
	clientSecret string
	scope        string
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		url := fmt.Sprintf("https://www.fitbit.com/oauth2/authorize?client_id=%s&scope=%s&response_type=code",
			url.QueryEscape(h.clientID), url.QueryEscape(h.scope))
		w.Write([]byte(fmt.Sprintf("<html><a href=\"%s\">Click to authenticate</a></html>", url)))
	case "/callback":
		fmt.Println(*r)
		codes, ok := r.URL.Query()["code"]
		if ok && len(codes) > 0 {
			code := codes[0]
			c := client.NewHTTPClient(nil)
			ac := c.Auth
			params := auth.NewOauthTokenParams()
			//params.ClientID = h.clientID
			params.Code = &code
			a := "Basic " + base64.StdEncoding.EncodeToString([]byte(h.clientID+":"+h.clientSecret))
			params.Authorization = &a
			params.GrantType = "authorization_code"
			fmt.Println(params)
			t, err := ac.OauthToken(params)
			if err != nil {
				fmt.Println(err.(*auth.OauthTokenBadRequest).Payload.Errors[0].Message)
				http.Error(w, err.Error(), 500)
				return
			}
			w.Write([]byte(fmt.Sprintf("<html>Authentication code: %s<br>Access Token: %s", codes[0], t.Payload.AccessToken)))

		} else {
			w.Write([]byte("No authentication code returned"))
		}
	default:
		w.WriteHeader(404)
	}
}

func main() {
	h := handler{
		clientID:     os.Getenv("FITBIT_CLIENT_ID"),
		clientSecret: os.Getenv("FITBIT_CLIENT_SECRET"),
		scope:        os.Getenv("FITBIT_SCOPE"),
	}
	s := &http.Server{
		Addr:           ":8080",
		Handler:        &h,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
