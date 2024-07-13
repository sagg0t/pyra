package auth

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
	"pyra/pkg/log"
	"pyra/pkg/session"
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
	l := log.FromContext(r.Context())
	s := session.FromCtx(r.Context())

	state := uuid.New()
	s.Values["state"] = state.String()
	url := googleConfig.AuthCodeURL(state.String())

	l.Debug("redirecting to Google for sign in")
	http.Redirect(w, r, url, http.StatusSeeOther)
}

func (api *API) GoogleCallback(w http.ResponseWriter, r *http.Request) {
	l := log.FromContext(r.Context())

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
		l.ErrorContext(ctx, "failed to exchange access grant", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	client := googleConfig.Client(ctx, tok)

	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		l.ErrorContext(ctx, "failed to fetch user details from Google", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		http.Error(
			w,
			fmt.Sprintf("Google responded with a %d trying to fetch user information", response.StatusCode),
			http.StatusInternalServerError,
		)
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

	l.Inspect(u)

	user, err := api.authSvc.SignIn(r.Context(), u)
	if err != nil {
		l.ErrorContext(r.Context(), "sign in failed", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	s := session.FromCtx(r.Context())

	s.Values[auth.UserIDSessionKey] = user.ID

	http.Redirect(w, r, "/foodProducts", http.StatusSeeOther)
}
