package main

import (
	"fmt"
	"net/http"

	"github.com/skrb7f16/rssagg/internal/auth"
	"github.com/skrb7f16/rssagg/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		fmt.Printf("%s%s", apiKey, err)

		if err != nil {
			responseWithErr(w, 403, fmt.Sprintf("Auth error: %s", err))
		}
		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			responseWithErr(w, 403, fmt.Sprintf("Auth error: %s", err))
		}
		handler(w, r, user)
	}
}
