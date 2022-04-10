# TextNow Api
a GO lang implementation of textnow's API.

# Usage
```go
api := New("email", "password")

messages := api.GetMessages()
for _, message := range *messages {
  log(message)
}
```
