# Bot API

This section describes the functionality of bots, which can be easily controlled using the HTTP API with almost any programming language.

## Authorization

Bots can be managed independently, so a unique access key is provided to access each one. Therefore, each request below must be executed with a [Authorization header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Authorization) with the value of the [access key](./admin.md#get-a-bots-access-key):

```nginx
Authorization: <BOT_ACCESS_KEY>
```

## Get bot

To get information about the current bot, use:

```nginx
GET /v1/bot-api/bot
```

Response will be like this:

```json
{
    "id": "01hc76p6ma76qfrs0bfr5x2pry",
    "name": "Test bot",
    "description": "Just testing bot.",
    "avatar": "https://google.com/image.png"
}
```

## Edit bot

To edit bot info, use this request with JSON body:

```nginx
PUT /v1/bot-api/bot
```

```json
{
    "name": "New bot name",
    "description": "New description.",
    "avatar": "https://google.com/new-avatar.jpeg"
}
```

| Body parameter | Description                                                   |
| -------------- | ------------------------------------------------------------- |
| name           | `Optional` <br> A name for the new bot.                       |
| description    | `Optional` <br> A brief description of what your bot will do. |
| avatar         | `Optional` <br> Link to an image (avatar) for your bot.       |

## Add bot commands

This request adds new commands to the current bot:

```nginx
POST /v1/bot-api/bot/commands
```

```json
[
    {
        "alias": "start",
        "description": "Start a new conversation."
    },
    {
        "alias": "help",
        "description": "Show usage guide."
    }
]
```

> Commands are an optional feature. Using commands, it is convenient to describe the functionality of the bot. Users seeing the list of available commands will be able to quickly understand how to interact with the bot. And developers will be able to easily write rules for processing user messages containing the described commands.

## Get bot commands

Get a list of available commands of the current bot:

```nginx
GET /v1/bot-api/bot/commands
```

Returned response:

```json
[
    {
        "id": "01hcjd34adz1va07vk97jfjmgq",
        "alias": "start",
        "description": "Start a new conversation."
    },
    {
        "id": "01hcjd34adz1va07vk99w3hzk3",
        "alias": "help",
        "description": "Show usage guide."
    }
]
```

## Remove bot command

Delete bot command by unique identifier:

```nginx
DELETE /v1/bot-api/bot/command/:id
```

## Add bot webhook

Connect the webhook to the current bot:

```nginx
POST /v1/bot-api/bot/webhook
```

```json
{
    "url": "http://localhost:4000/webhook",
    "secret": "SOME_SECRET"
}
```

| Body parameter | Description                                                                                                                                                                                       |
| -------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| url            | `Required` <br> The address of the webhook to which user messages will be sent. Requests will come with a POST method and JSON body with a message as described [here](./client.md#send-message). |
| secret         | `Optional` <br> Random data, in order to be able to verify that queries are coming from the correct source.                                                                                       |

## Get bot webhook

Get webhook info for current bot:

```nginx
GET /v1/bot-api/bot/webhook
```

Response:

```json
{
    "botId": "01hc76p6ma76qfrs0bfr5x2pry",
    "url": "http://localhost:4000/webhook",
    "secret": "SOME_SECRET"
}
```

## Delete bot webhook

Remove the webhook for the current bot:

```nginx
DELETE /v1/bot-api/bot/webhook
```

## Get chats

Bots can only receive a list of available chat rooms by user:

```nginx
GET /v1/bot-api/chats
```

| Query parameter | Description                             |
| --------------- | --------------------------------------- |
| user_id         | `Required` <br> Unique user identifier. |

Returned response:

```json
[
    {
        "id": "01hc7aa947s4fdwywwjqr14wh7",
        "userId": "01hc71yz5pny5bmf0b4055zbjw",
        "botId": "01hc76p6ma76qfrs0bfr5x2pry"
    }
]
```

## Send message

To send a message from a bot, use this request with a JSON body:

```nginx
POST /v1/bot-api/chat/message
```

```json
{
    "chatId": "01hc7aa947s4fdwywwjqr14wh7",
    "body": "Hello World!",
    "attachmentIds": []
}
```

| Body parameter | Description                                                                                                   |
| -------------- | ------------------------------------------------------------------------------------------------------------- |
| chatId         | `Required` <br> The unique ID of the chat to which the message is sent.                                       |
| body           | `Required` <br> The text body of the message.                                                                 |
| attachmentIds  | `Optional` <br> Array of unique identifiers for [files](#upload-files) to be attached to the current message. |

Returned response:

```json
{
    "id": "01hc7akt8z3w2a5ayfcrvkfy52",
    "chatId": "01hc7aa947s4fdwywwjqr14wh7",
    "senderId": "01hc76p6ma76qfrs0bfr5x2pry",
    "body": "Hello World!",
    "timestamp": 1696758098208
}
```

## Get messages

To get a list of messages from the chat, use the following request:

```nginx
GET /v1/bot-api/chat/:id/messages
```

| Query parameter | Description                                                                                                                                               |
| --------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------- |
| page            | `Optional` <br> Message page number. Default is 1.                                                                                                        |
| limit           | `Optional` <br> Maximum number of messages per page. Default is 20.                                                                                       |
| sender_id       | `Optional` <br> Filter messages by the ID of the specified sender. For example, you can filter messages to get only user messages, ignoring bot messages. |
| since           | `Optional` <br> The [timestamp](https://currentmillis.com/) (in milliseconds) from which you want to search for messages.                                 |

Returned response will be like this:

```json
{
    "chatId": "01hc7aa947s4fdwywwjqr14wh7",
    "total": 1,
    "page": 1,
    "limit": 20,
    "messages": [
        {
            "id": "01hc7akt8z3w2a5ayfcrvkfy52",
            "chatId": "01hc7aa947s4fdwywwjqr14wh7",
            "senderId": "01hc76p6ma76qfrs0bfr5x2pry",
            "body": "Hello World!",
            "timestamp": 1696758098208
        }
    ]
}
```

> The `total` field is an indicator of the number of all messages found.

## Get message

Get message by their unique indenifier:

```nginx
GET /v1/bot-api/chat/message/:id
```

Response:

```json
{
    "id": "01hc7akt8z3w2a5ayfcrvkfy52",
    "chatId": "01hc7aa947s4fdwywwjqr14wh7",
    "senderId": "01hc76p6ma76qfrs0bfr5x2pry",
    "body": "Hello!",
    "timestamp": 1696758098208
}
```

## Upload files

To upload files use this request:

```nginx
POST /v1/bot-api/files
```

with form data as a body:

```sh
Content-Disposition: form-data; name="file"; filename="yourfile.png"
Content-Type: image/png
# Content...
```

The parameter `name` is always should be equals to _file_.

> Unique file identifiers can be used as attachments for [messages](#send-message).
