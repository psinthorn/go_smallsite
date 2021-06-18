package utils

import (
	"net/http"
)

func (u *Utils) IsAuthenticated(r *http.Request) bool {
	exists := appConfig.Session.Exists(r.Context(), "user_id")
	return exists
}
