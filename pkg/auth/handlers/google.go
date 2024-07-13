package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"pyra/pkg/auth"
)

var googleConfig = &oauth2.Config{
	ClientID:     os.Getenv("GOOGLE_OAUTH2_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_OAUTH2_CLIENT_SECRET"),
	RedirectURL:  "http://localhost:42069/auth/google/callback",
	Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile",
		"openid",
	},
	Endpoint: google.Endpoint,
}

func (api *API) GoogleAuth(w http.ResponseWriter, r *http.Request) {
	log := api.RequestLogger(r)
	session := api.Session(r)

	state := uuid.New()
	session.Values["state"] = state.String()
	url := googleConfig.AuthCodeURL(state.String())

	log.Debug("redirecting to Google for sign in")
	http.Redirect(w, r, url, http.StatusSeeOther)
}

func (api *API) GoogleCallback(w http.ResponseWriter, r *http.Request) {
	log := api.RequestLogger(r)

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	code := r.FormValue("code")
	if code == "" {
		log.ErrorContext(r.Context(), "invalid code")
		api.InternalServerError(w)
		return
	}

	// Use the custom HTTP client when requesting a token.
	httpClient := &http.Client{Timeout: 2 * time.Second}
	ctx := context.WithValue(r.Context(), oauth2.HTTPClient, httpClient)

	tok, err := googleConfig.Exchange(ctx, code)
	if err != nil {
		log.ErrorContext(ctx, "failed to exchange access grant", "error", err)
		api.InternalServerError(w)
		return
	}

	client := googleConfig.Client(ctx, tok)

	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		log.ErrorContext(ctx, "failed to fetch user details from Google", "error", err)
		api.InternalServerError(w)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.ErrorContext(
			ctx,
			fmt.Sprintf("Google responded with a %d trying to fetch user information", response.StatusCode),
		)
		api.InternalServerError(w)
		return
	}

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.ErrorContext(ctx, "failed to read response", "error", err)
		api.InternalServerError(w)
		return
	}

	var u auth.GoogleUser
	if err := json.Unmarshal(responseBytes, &u); err != nil {
		log.ErrorContext(ctx, "failed to unmarshal response", "error", err)
		api.InternalServerError(w)
		return
	}

	log.Inspect(u)

	user, err := api.authSvc.SignIn(r.Context(), u)
	if err != nil {
		log.ErrorContext(ctx, "sign in failed", "error", err)
		api.InternalServerError(w)
		return
	}

	session := api.Session(r)

	session.Values[auth.UserIDSessionKey] = user.ID

	http.Redirect(w, r, "/foodProducts", http.StatusSeeOther)
}
