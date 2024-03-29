package web

import (
	"net/http"
	"time"
)

func DeleteCookie(
	w http.ResponseWriter,
	cookie string,
	now time.Time,
) {
	http.SetCookie(w, &http.Cookie{
		Name:     cookie,
		Value:    "",
		Expires:  now.Add(-time.Hour),
		Secure:   true,
		Path:     "/",
		HttpOnly: true,
	})
}

func SetCookie(
	w http.ResponseWriter,
	cookie string,
	value string,
	exp time.Time,
) {
	http.SetCookie(w, &http.Cookie{
		Name:     cookie,
		Value:    value,
		Expires:  exp,
		Secure:   true,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
}

func SetSessionCookie(
	w http.ResponseWriter,
	cookie string,
	value string,
) {
	http.SetCookie(w, &http.Cookie{
		Name:     cookie,
		Value:    value,
		Secure:   true,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
}
