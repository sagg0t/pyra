package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/olehvolynets/pyra/pkg/auth"
)

var googleConfig = &oauth2.Config{
	ClientID:     "asdlfkj",
	ClientSecret: "asdlfkj",
	RedirectURL:  "http://localhost:42069/auth/google/callback",
	Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile",
		"openid",
	},
	Endpoint: google.Endpoint,
}

func (api *API) GoogleAuth(w http.ResponseWriter, r *http.Request) {
	url := googleConfig.AuthCodeURL("state")
	http.Redirect(w, r, url, http.StatusPermanentRedirect)
}

func (api *API) GoogleCallback(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	code := r.FormValue("code")
	if code == "" {
		http.Error(w, "invalid code", http.StatusInternalServerError)
		return
	}

	// Use the custom HTTP client when requesting a token.
	httpClient := &http.Client{Timeout: 2 * time.Second}
	ctx := context.WithValue(r.Context(), oauth2.HTTPClient, httpClient)

	tok, err := googleConfig.Exchange(ctx, code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client := googleConfig.Client(ctx, tok)

	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("responded with a %d trying to fetch user information", response.StatusCode), http.StatusInternalServerError)
		return
	}

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var u auth.GoogleUser
	if err := json.Unmarshal(responseBytes, &u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := api.authSvc.SignIn(r.Context(), u)
	if err != nil {
		api.log.ErrorContext(r.Context(), "sign in failed", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	prettyUser, _ := json.MarshalIndent(user, "", "	")
	fmt.Println(string(prettyUser))

	http.Redirect(w, r, "/signIn", http.StatusFound)
}
