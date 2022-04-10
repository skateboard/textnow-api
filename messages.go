package TextNowAPI

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// GetMessages This gets most of the messages both sent and received. However It won't get all of them just the past 10-15
func (t *TextNowAPI) GetMessages() *[]Message {
	request, err := http.NewRequest("GET", fmt.Sprintf("https://www.textnow.com/api/users/%v/messages", t.username), nil)
	if err != nil {
		fmt.Println(err)
		return nil
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
		fmt.Println(err)
		return nil
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		fmt.Println(response.StatusCode)
		return nil
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var messagesResponse MessagesResponse
	err = json.Unmarshal(body, &messagesResponse)
	if err != nil {
		fmt.Println(err)
		return nil
	}


	return &messagesResponse.Messages
}

// GetUnreadMessages This gets the unread messages
func (t *TextNowAPI) GetUnreadMessages() *[]Message {
	messages := t.GetMessages()
	if messages == nil {
		return nil
	}

	var unreadMessages []Message
	for _, message := range *messages {
		if !message.Read {
			unreadMessages = append(unreadMessages, message)
		}
	}

	return &unreadMessages
}

func (t *TextNowAPI) markMessageAsRead(message Message) bool {
	url := fmt.Sprintf("https://www.textnow.com/api/users/%v/conversations/%v?latest_message_id=%v&http_method=PATCH", t.username, t.phoneNumber, message.Id)

	request, err := http.NewRequest("PATCH", url, nil)
	if err != nil {
		fmt.Println(err)
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
		fmt.Println(err)
		return false
	}

	return response.StatusCode == 200
}

// WaitForResponse This waits for a response from a specific phone number, if timeout is reached it returns nil. (timeout is 5 minutes)
func (t *TextNowAPI) WaitForResponse(phoneNumber string, timeout bool) *Message {
	receivedChannel := make(chan *Message)

	go func() {
		for {
			messages := t.GetUnreadMessages()
			if messages == nil {
				continue
			}

			for _, message := range *messages {
				t.markMessageAsRead(message)
				if message.ContactValue == phoneNumber {
					receivedChannel <- &message
					break
				}
			}

			time.Sleep(time.Second * 5)
		}
	}()

	select {
	case message := <-receivedChannel:
		return message
	case <-time.After(time.Minute * 5):
		if timeout {
			return nil
		}

		return t.WaitForResponse(phoneNumber, true)
	}
}