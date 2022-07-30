package utils

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

var ttl = 24 * time.Hour

func AddCookie(w http.ResponseWriter, name, value string) {
	expire := time.Now().Add(ttl)
	cookie := http.Cookie{
		Name:     name,
		Value:    value,
		Expires:  expire,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
}

func GetUserID(r *http.Request) (int, error) {
	cookie, err := r.Cookie("id")
	if err != nil || cookie == nil {
		return 0, fmt.Errorf("GetUserID no cookie or err: %w", err)
	}

	num, err := strconv.Atoi(cookie.Value)
	if err != nil || cookie == nil {
		return 0, fmt.Errorf("GetUserID no convert coonie err: %w", err)
	}

	return num, nil
}
