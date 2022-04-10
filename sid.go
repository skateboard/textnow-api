package TextNowAPI

import (
	"fmt"
	"net/http"
)

func (t *TextNowAPI) getSID() bool {
	request, err := http.NewRequest("GET", "https://www.textnow.com/api/init", nil)
	if err != nil {
		fmt.Print(err)
		return false
	}

	request.Header = http.Header{
		"user-agent": []string{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:99.0) Gecko/20100101 Firefox/99.0"},
		"accept": []string{"application/json, text/plain, */*"},
		"accept-language": []string{"en-US,en;q=0.5"},
		"referer": []string{"https://www.textnow.com/login"},
		"x-csrf-token": []string{t.xsrfToken},
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
		fmt.Print(err)
		return false
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		fmt.Print(response.StatusCode)
		return false
	}

	if response.Cookies()[0] != nil {
		t.sid = response.Cookies()[0].Value
		return true
	}

	return false
}