# Working with Python

This guide shows an example of developing a simple bot in Python using only the standard library of the language.

## Writting code

First, let's import some modules and define the necessary constants:

```py
import json
import threading
import http.server
import socketserver

port = 4000
api_server = "localhost:7007"
bot_key = "PASTE_BOT_KEY_HERE"
```

Next, let's write a function to send messages to the server on behalf of the bot:

```py
def send_message(chat_id, body):
    json_body = json.dumps({'chatId': chat_id, 'body': body})

    conn = http.client.HTTPConnection(api_server)
    conn.request("POST", "/v1/bot-api/chat/message", body=json_body, headers={
        'Content-Type': 'application/json',
        'Authorization': bot_key
    })

    response = conn.getresponse()
    if response.status >= 400:
        print(f'Get response with error: {response.status}, {response.reason}')
```

Now let's write a custom handler for user messages, with some simple logic:

```py
def message_handler(msg):
    reply = ""

    if msg['body'] == "/start":
        reply = "Hello!\n\nLet's play with me using /ping command."
    elif msg['body'] == "/ping":
        reply = "PONG"
    else:
        reply = "Sorry, but I don't understand you."

    # Send response to the user
    send_message(msg['chatId'], reply)
```

Let's create an HTTP handler for the webhook to receive messages from users and pass them to the handler described above:

```py
class CustomHandler(http.server.SimpleHTTPRequestHandler):
    def do_POST(self):
        if self.path == '/webhook':
            content_length = int(self.headers['Content-Length'])
            post_data = self.rfile.read(content_length)
            try:
                message = json.loads(post_data.decode('utf-8'))
                self.send_response(200)
                self.end_headers()

                # Handling message in other thread
                async_thread = threading.Thread(
                    target=message_handler,
                    args=([message])
                )
                async_thread.start()
            except ValueError:
                self.send_response(400)
                self.end_headers()
        else:
            self.send_response(404)
            self.end_headers()
```

And finally, let's get our bot up and running:

```py
with socketserver.TCPServer(("", port), CustomHandler) as httpd:
    print("Bot is running...")
    httpd.serve_forever()
```

To start our bot, run this command:

```
python3 main.py
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
