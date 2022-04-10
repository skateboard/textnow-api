# TextNow Api
a GO lang implementation of textnow's API.

# Usage
```go
api := New("email", "password", "proxy")

messages := api.GetMessages()
for _, message := range *messages {
  log(message)
}
```
