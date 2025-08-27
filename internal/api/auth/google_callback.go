package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/oauth2"

	"pyra/internal/api/base"
	"pyra/pkg/auth"
)

type GoogleCallbackHandler struct {
	*base.Handler
	authSvc *auth.AuthService
}

func (h *GoogleCallbackHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log := h.RequestLogger(r)

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	code := r.FormValue("code")
	if code == "" {
		log.ErrorContext(r.Context(), "invalid code")
		h.InternalServerError(w)
		return
	}

	// Use the custom HTTP client when requesting a token.
	httpClient := &http.Client{Timeout: 2 * time.Second}
	ctx := context.WithValue(r.Context(), oauth2.HTTPClient, httpClient)

	tok, err := googleConfig.Exchange(ctx, code)
	if err != nil {
		log.ErrorContext(ctx, "failed to exchange access grant", "error", err)
		h.InternalServerError(w)
		return
	}

	client := googleConfig.Client(ctx, tok)

	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		log.ErrorContext(ctx, "failed to fetch user details from Google", "error", err)
		h.InternalServerError(w)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.ErrorContext(
			ctx,
			fmt.Sprintf("Google responded with a %d trying to fetch user information", response.StatusCode),
		)
		h.InternalServerError(w)
		return
	}

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.ErrorContext(ctx, "failed to read response", "error", err)
		h.InternalServerError(w)
		return
	}

	var u auth.GoogleUser
	if err := json.Unmarshal(responseBytes, &u); err != nil {
		log.ErrorContext(ctx, "failed to unmarshal response", "error", err)
		h.InternalServerError(w)
		return
	}

	log.Inspect(u)

	user, err := h.authSvc.SignIn(r.Context(), u)
	if err != nil {
		log.ErrorContext(ctx, "sign in failed", "error", err)
		h.InternalServerError(w)
		return
	}

	session := h.Session(r)
	session.Values[base.UserIDSessionKey] = user.ID

	http.Redirect(w, r, "/products", http.StatusSeeOther)
}
