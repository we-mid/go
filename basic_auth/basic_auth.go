package basic_auth

import (
	"encoding/json"
	"net/http"
	"os"
)

type User struct {
	Username string `json:"user"`
	Password string `json:"pass"`
}

var (
	userMap = map[string]User{}
)

func AddUser(user User) {
	userMap[user.Username] = user
}

func InitFromEnv() error {
	var list []User
	JSON := os.Getenv("BASICAUTH_USERLIST")
	if err := json.Unmarshal([]byte(JSON), &list); err != nil {
		return err
	}
	for _, user := range list {
		AddUser(user)
	}
	return nil
}

func Wrap(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := Check(r)
		if user == "" { // reject
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized.", http.StatusUnauthorized)
			return
		}
		next(w, r) // pass
	}
}

func Check(r *http.Request) string {
	username, password, ok := r.BasicAuth()
	if ok {
		if user, ok := userMap[username]; ok {
			if password == user.Password {
				return username
			}
		}
	}
	return ""
}
