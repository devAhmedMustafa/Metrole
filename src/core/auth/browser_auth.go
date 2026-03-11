package auth

import (
	"encoding/json"
	"fmt"
	"metrole/src/config"
	"net/http"
	"net/url"
	"os/exec"
	"runtime"
	"strings"
)

func StartCallbackListener() chan string {
	codeChan := make(chan string)

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {

		code := r.URL.Query().Get("code")
		fmt.Fprintln(w, "Login successful! You can close this window.")

		codeChan <- code

	})

	go http.ListenAndServe(":1294", nil)

	return codeChan
}

func InitiateLogin() error {

	client_id := config.AppSettings.GOOGLE_CLIENT_ID
	redirect_uri := "http://localhost:1294/login"
	response_type := "code"
	scope := "openid email"

	auth_url := fmt.Sprintf("https://accounts.google.com/o/oauth2/v2/auth?client_id=%s&redirect_uri=%s&response_type=%s&scope=%s",
		client_id, redirect_uri, response_type, scope)

	return openBrowser(auth_url)
}

func openBrowser(url string) error {

	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "rundll32"
		args = []string{"url.dll,FileProtocolHandler", url}

	case "darwin":
		cmd = "open"
		args = []string{url}

	default:
		cmd = "xdg-open"
		args = []string{url}
	}

	return exec.Command(cmd, args...).Start()

}

func ExchangeCodeForIdToken(code string) (string, error) {

	data := url.Values{}
	data.Set("code", code)
	data.Set("client_id", config.AppSettings.GOOGLE_CLIENT_ID)
	data.Set("client_secret", config.AppSettings.GOOGLE_CLIENT_SECRET)
	data.Set("redirect_uri", "http://localhost:1294/login")
	data.Set("grant_type", "authorization_code")

	resp, err := http.Post(
		"https://oauth2.googleapis.com/token",
		"application/x-www-form-urlencoded",
		strings.NewReader(data.Encode()),
	)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		IdToken string `json:"id_token"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	return result.IdToken, nil
}
