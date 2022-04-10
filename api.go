package TextNowAPI

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"time"
)

type TextNowAPI struct {
	Email string
	Password string

	client http.Client

	// The following are set by the login flow
	xsrfToken string
	username string
	phoneNumber string
	sid string
}

func New(email, password string) *TextNowAPI {
	jar, _ := cookiejar.New(nil);
	client := http.Client{
		Jar:           jar,
		Timeout:       10 * time.Second,
	}
	api := &TextNowAPI{
		Email: email,
		Password: password,
		client: client,
	}
	if !api.initialCookies() {
		return nil
	}

	if !api.login() {
		fmt.Println("Login failed")
		return nil
	}

	if !api.getSID() {
		fmt.Println("Failed to get SID")
		return nil
	}

	if !api.getUserInformation() {
		fmt.Println("Failed to get user information")
		return nil
	}

	return api
}

func (t *TextNowAPI) initialCookies() bool {
	_, err := t.client.Get("https://www.textnow.com/")
	if err != nil {
		fmt.Println("Error getting initial cookies: ", err)
		return false
	}

	return t.webFlow()
}

func (t *TextNowAPI) webFlow() bool {
	_, err := t.client.Get("https://www.textnow.com/webflow")
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}