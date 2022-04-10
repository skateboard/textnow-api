package TextNowAPI

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// login sometimes gets PX'd so maybe make a manual mode and auto login.
func (t *TextNowAPI) login() bool {
	if !t.getLoginParameters() {
		fmt.Println("Error: Could not get login parameters")
		return false
	}
	payload := LoginPayload{
		Json: fmt.Sprintf("{\"remember\":true,\"username\":\"%v\",\"password\":\"%v\",\"disable_session\":false}", t.Email, t.Password),
	}
	jsonBytes, _ := json.Marshal(payload)

	request, err := http.NewRequest("POST", "https://www.textnow.com/api/sessions", bytes.NewBuffer(jsonBytes))
	if err != nil {
		fmt.Println("Error: Could not create request")
		return false
	}

	request.Header = http.Header{
		"user-agent": []string{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:99.0) Gecko/20100101 Firefox/99.0"},
		"accept": []string{"application/json, text/plain, */*"},
		"accept-language": []string{"en-US,en;q=0.5"},
		"referer": []string{"https://www.textnow.com/login"},
		"x-csrf-token": []string{t.xsrfToken},
		"content-type": []string{"application/json;charset=utf-8"},
		"x-xsrf-token": []string{t.xsrfToken},
		"origin": []string{"https://www.textnow.com"},
		"dnt": []string{"1"},
		"sec-fetch-dest": []string{"empty"},
		"sec-fetch-mode": []string{"cors"},
		"sec-fetch-site": []string{"same-origin"},
		"te": []string{"trailers"},
	}

	response, err := t.client.Do(request)
	if err != nil {
		fmt.Println("Error: Could not send request")
		return false
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error: Could not read response body")
		return false
	}

	if response.StatusCode == 403 {
		fmt.Println("GETTING PX'D RETRYING IN 5 MINUTES")

		time.Sleep(time.Minute * 5)
		return t.login()
	} else if response.StatusCode != 200 {
		fmt.Println(response.StatusCode)
		fmt.Println(string(body))

		fmt.Println("Error: Could not login")
		return false
	}

	var loginResponse LoginResponse
	err = json.Unmarshal(body, &loginResponse)
	if err != nil {
		fmt.Println("Error: Could not unmarshal login response")
		return false
	}
	t.username = loginResponse.Username

	return true
}

func (t *TextNowAPI) getLoginParameters() bool {
	request, err := http.NewRequest("GET", "https://www.textnow.com/login", nil)
	if err != nil {
		fmt.Println(err)
		return false
	}

	request.Header = http.Header{
		"User-Agent": []string{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:99.0) Gecko/20100101 Firefox/99.0"},
		"Accept":     []string{"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8"},
		"Accept-Language": []string{"en-US,en;q=0.5"},
		"Accept-Encoding": []string{"gzip, deflate, br"},
		"DNT": []string{"1"},
		"Referer": []string{"https://www.textnow.com/"},
		"Upgrade-Insecure-Requests": []string{"1"},
		"Sec-Fetch-Dest": []string{"document"},
		"Sec-Fetch-Mode": []string{"navigate"},
		"Sec-Fetch-Site": []string{"same-origin"},
		"Sec-Fetch-User": []string{"?1"},
		"TE": []string{"trailers"},
	}

	response, err := t.client.Do(request)
	if err != nil {
		fmt.Println(err)
		return false
	}

	if response.StatusCode != 200 {
		fmt.Println(response.StatusCode)
		return false
	}
	t.xsrfToken = response.Cookies()[0].Value

	return true
}