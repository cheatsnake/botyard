# Working with Go

This guide shows an example of developing a simple bot in Go using only the standard library of the language.

## Writting code

First, let's import some packages and define the necessary constants:

```go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const port = "4000"
const botApi = "http://localhost:7007/v1/bot-api"
const botKey = "PASTE_BOT_KEY_HERE"
```

Let's define a structure for the message type to encode and decode JSON correctly:

```go
type Message struct {
	Id            string   `json:"id,omitempty"`
	ChatId        string   `json:"chatId"`
	SenderId      string   `json:"senderId,omitempty"`
	Body          string   `json:"body"`
	AttachmentIds []string `json:"attachmentIds,omitempty"`
	Timestamp     int64    `json:"timestamp,omitempty"`
}
```

Next, let's write a function to send messages to the server on behalf of the bot:

```go
func sendMessage(chatId, body string) {
	jsonBody, err := json.Marshal(&Message{
		ChatId: chatId,
		Body:   body,
	})
	if err != nil {
		fmt.Printf("can't marshal json %s\n", err.Error())
	}

	req, err := http.NewRequest(
		http.MethodPost,
		botApi+"/chat/message",
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		fmt.Printf("can't make a new request %s\n", err.Error())
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", botKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("can't send message to user %s\n", err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		fmt.Printf("got response with error code: %d\n", resp.StatusCode)
	}
}
```

Now let's write a custom handler for user messages, with some simple logic:

```go
func messageHandler(msg Message) {
	reply := ""

	switch msg.Body {
	case "/start":
		reply = "Hello!\n\nLet's play with me using /ping command."
	case "/ping":
		reply = "PONG"
	default:
		reply = "Sorry, but I don't understand you."
	}

	sendMessage(msg.ChatId, reply)
}
```

Let's create an HTTP handler for the webhook to receive messages from users and pass them to the handler described above:

```go
func webhookHandler(w http.ResponseWriter, r *http.Request) {
    // Accept only POST requests on /webhook path
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var msg Message

	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

    // Handling a user message in a separate goroutine
	go messageHandler(msg)

	w.WriteHeader(http.StatusOK)
}
```

Finally, let's declare the main function, which will be the entry point of our bot, where the HTTP server will be started:

```go
func main() {
	http.HandleFunc("/webhook", webhookHandler)

	fmt.Println("Bot is running...")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
```

To start our bot, run this command:

```
go run main.go
```

## Webhook setup

To start receiving messages from users, you need to connect a webhook for our bot. We have launched the bot on a local network on a port 4000, so our webhook will look like this:

```sh
curl --request POST \
  --url http://localhost:7007/v1/bot-api/bot/webhook \
  --header 'Authorization: PASTE_BOT_TOKEN_HERE' \
  --header 'Content-Type: application/json' \
  --data '{"url": "http://localhost:4000/webhook"}'
```

> Note that if you running Botyard using Docker, accessing to localhost from the container will not work. Watch [this video](https://youtu.be/NZGu-9KQVsE?si=n2KM4BKmIF4yDtox) to solve the problem.

See the full [bot API reference](../api/bot.md).
