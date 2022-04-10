package TextNowAPI

import (
	"os"
	"testing"
)

func TestAPI(t *testing.T) {
	api := New(os.Getenv("EMAIL"), os.Getenv("PASSWORD"), "")

	messages := api.GetMessages()
	for _, message := range *messages {
		t.Log(message)
	}

	message := api.WaitForResponse("13123123", false)
	t.Log(message)
}
