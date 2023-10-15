# Working with Node.js

Node.js is a great choice for bot development. It is able to handle a huge number of requests thanks to an asynchronous behavior model. In the current guide, an example of creating a simple bot using only the capabilities of the standard library.

## Preparations

1. Start with a new folder:

    ```sh
    mkdir ping-pong-bot
    ```

    ```sh
    cd ./ping-pong-bot
    ```

2. Init a new project:

    ```sh
    npm init -y
    ```

3. Create the main project file:

    ```sh
    touch main.js
    ```

## Writting code

First, let's import some modules and define the necessary constants:

```js
import http from "node:http";

const PORT = "4000";
const BOT_API = "http://localhost:7007/v1/bot-api";
const BOT_KEY = "PASTE_BOT_KEY_HERE";
```

To send messages to users, you need to create an appropriate function:

```js
const sendMessage = async (chatId, body) => {
    try {
        const jsonBody = JSON.stringify({ chatId, body });

        const resp = await fetch(`${BOT_API}/chat/message`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: BOT_KEY,
            },
            body: jsonBody,
        });

        if (resp.status >= 400) {
            console.error(`Got response with error code: ${resp.status}`);
        }
    } catch (error) {
        throw error;
    }
};
```

Next let's implement a custom message handler. This is the place where all the logic of your bot's work is performed:

```js
const messageHandler = async (userMsg) => {
    try {
        let reply = "";

        switch (userMsg.body) {
            case "/start":
                reply = "Hello!\n\nLet's play with me using /ping command.";
                break;
            case "/ping":
                reply = `PONG`;
                break;

            default:
                reply = "Sorry, but I don't understand you.";
                break;
        }

        // Send response to user
        await sendMessage(userMsg.chatId, reply);
    } catch (error) {
        console.error(`Message handling failed: ${String(error)}`);
    }
};
```

Then, let's declare a handler for the webhook to receive messages from users:

```js
const webhookHandler = async (req, res) => {
    try {
        // Accept only POST requests on /webhook path
        if (!req.method === "POST" && !req.url === "/webhook") {
            res.writeHead(404);
            res.end();
            return;
        }

        // Read request body -----------------------------------
        let body = "";
        req.on("data", (chunk) => (body += chunk));
        await new Promise((resolve) => req.on("end", resolve));
        // -----------------------------------------------------

        const message = JSON.parse(body);

        res.writeHead(200);
        res.end();

        // Handling user messages
        messageHandler(message);
    } catch (error) {
        res.writeHead(400, { "Content-Type": "application/json" });
        res.end(JSON.stringify({ error: String(error) }));
    }
};
```

And finally, we will create an entry point for our bot, where the HTTP server will be started:

```js
const main = async () => {
    try {
        const server = new http.Server(webhookHandler);
        server.listen(PORT, () => console.log("Bot is running..."));
    } catch (error) {
        console.error(error);
        process.exit(1);
    }
};

main();
```

To start our bot, run this command:

```
node main.js
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
