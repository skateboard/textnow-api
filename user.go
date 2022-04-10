package TextNowAPI

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// getUserInformation returns the user information for the given username
func (t *TextNowAPI) getUserInformation() bool {
	request, err := http.NewRequest("GET", "https://www.textnow.com/api/users/" + t.username, nil)
	if err != nil {
		fmt.Println("Error creating request: ", err)
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
		fmt.Println("Error sending request: ", err)
		return false
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		fmt.Println("Error getting user information: ", response.StatusCode)
		return false
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body: ", err)
		return false
	}

	var userResponse UserInformationResponse
	err = json.Unmarshal(body, &userResponse)
	if err != nil {
		fmt.Println("Error decoding user information: ", err)
		return false
	}

	t.phoneNumber = userResponse.PhoneNumber
	return true
}