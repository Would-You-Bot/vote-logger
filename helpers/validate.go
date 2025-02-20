package helpers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Would-You-Bot/vote-logger/config"
)

func Validate(r *http.Request, w http.ResponseWriter) bool {
	ct := r.Header.Get("Content-Type")
	if ct != "" {
		mediaType := strings.ToLower(strings.TrimSpace(strings.Split(ct, ";")[0]))
		if mediaType != "application/json" {
			msg := "Content-Type header is not application/json"
			fmt.Println(msg)
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return false
		}
	}

	if r.Header.Get("Authorization") != config.Conf.BotList.Topgg.Auth {
		fmt.Println("Invalid authorization")
		http.Error(w, "Invalid authorization", http.StatusUnauthorized)
		return false
	}

	return true
}