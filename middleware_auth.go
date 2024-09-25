package main

import (
	"fmt"
	"net/http"

	"github.com/lghartmann/go-rss-web-scraper/internal/auth"
	"github.com/lghartmann/go-rss-web-scraper/internal/database"
)

type authorizedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authorizedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth Error: %v", err))
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Couldn't get user: %v", err))
		}

		handler(w, r, user)
	}
}
