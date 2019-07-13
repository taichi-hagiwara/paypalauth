package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/taichi-hagiwara/paypalauth"
)

var client = paypalauth.Client{
	ClientID:     os.Getenv("PAYPAL_CLIENT_ID"),
	ClientSecret: os.Getenv("PAYPAL_CLIENT_SECRET"),
	Endpoint:     paypalauth.SandboxEndpoint,
}

var config = client.OAuth2Config("http://localhost:4000/callback", "email")

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", demo)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/callback", callback)

	if err := http.ListenAndServe("localhost:4000", mux); err != nil {
		panic(err)
	}
}

func demo(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello, World!!")
	fmt.Fprintln(w, "To login with paypal, access `/login`")
}

func login(w http.ResponseWriter, r *http.Request) {
	url := config.AuthCodeURL("hello")
	w.Header().Add("Location", url)
	w.WriteHeader(http.StatusSeeOther)
}

func callback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		panic("oauth2 error: no code")
	}

	t, err := config.Exchange(context.Background(), code)
	if err != nil {
		panic(err)
	}

	token := client.Token(t)

	var user userInfo

	if err := token.UserInfo(context.Background(), &user); err != nil {
		panic(err)
	}

	w.Header().Add("Content-type", "text/plain")
	w.WriteHeader(http.StatusOK)

	k, _ := json.MarshalIndent(user, "", "  ")
	fmt.Fprintln(w, string(k))
}

type userInfo struct {
	UserID          string `json:"user_id"`
	Name            string `json:"name"`
	GivenName       string `json:"given_name"`
	FamilyName      string `json:"family_name"`
	VerifiedAccount bool   `json:"verified_account"`
	PayerID         string `json:"payer_id"`
	MailAddresses   []struct {
		Value   string `json:"value"`
		Type    string `json:"type"`
		Primary bool   `json:"primary"`
	} `json:"emails"`
	Address struct {
		StreetAddress string `json:"street_address"`
		Locality      string `json:"locality"`
		Region        string `json:"region"`
		PostalCode    string `json:"postal_code"`
		Country       string `json:"country"`
	} `json:"address"`
}
