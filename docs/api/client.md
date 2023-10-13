# Client API

The requests described in this section can be used to build client applications. Currently this API is [used](https://github.com/cheatsnake/botyard/blob/master/web/src/api/client-api.ts) in the [official web-client](https://github.com/cheatsnake/botyard/tree/master/web) of the platform.

## Get service info

Get service info defined in [`botyard.config.json`](../configuration.md#config-file) file:

```nginx
GET /v1/client-api/service-info
```

Response:

```json
{
    "name": "Service name",
    "description": "Some info about your service.",
    "socials": [
        {
            "title": "Website",
            "url": "https://example.com"
        },
        {
            "title": "GitHub",
            "url": "https://github.com"
        },
        {
            "title": "Discord",
            "url": "https://discord.com"
        },
        {
            "title": "Twitter",
            "url": "https://twitter.com"
        }
    ]
}
```

## Get all bots

Get a list of all available bots on the current server:

```nginx
GET /v1/client-api/bots
```

Response:

```json
[
    {
        "id": "01hc76p6ma76qfrs0bfr5x2pry",
        "name": "Test bot",
        "description": "Just testing bot.",
        "avatar": "https://google.com/img.png"
    }
]
```

## Get bot

Get a single bot by its ID:

```nginx
GET /v1/client-api/bot/:id
```

Response:

```json
{
    "id": "01hc76p6ma76qfrs0bfr5x2pry",
    "name": "Test bot",
    "description": "Just testing bot.",
    "avatar": "https://google.com/img.png"
}
```

## Create user (login)

Use this request with a JSON body to create a new user:

```nginx
POST /v1/client-api/user
```

```json
{
    "nickname": "user"
}
```

Response:

```json
{
    "id": "01hc71yz5pny5bmf0b4055zbjw",
    "nickname": "user"
}
```

⚠️ After executing this request, a response with cookies will be returned in which a JWT token with information about the current user will be sewn. All the requests listed below should be executed only with cookies setted.

> The lifetime of cookies is currently 7 days. If the user does not use the platform for more than 7 days, cookies will not be able to be updated, which will make them invalid. Therefore, the user will lose the ability to view his past data and he will have to use this request again.

## Get current user

Authorised users can perform the following request:

```nginx
GET /v1/client-api/user
```

Response:

```json
{
    "id": "01hc71yz5pny5bmf0b4055zbjw",
    "nickname": "user"
}
```

## Get bot commands

Get a list of available commands of the specified bot:

```nginx
GET /v1/client-api/bot/:id/commands
```

Response will be like this:

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

## Create chat

Users can create chats to communicate with bots. It only requires the ID of a specific bot:

```nginx
POST /v1/client-api/chat
```

```json
{
    "botId": "01hc76p6ma76qfrs0bfr5x2pry"
}
```

Returned response:

```json
{
    "id": "01hc7aa947s4fdwywwjqr14wh7",
    "userId": "01hc71yz5pny5bmf0b4055zbjw",
    "botId": "01hc76p6ma76qfrs0bfr5x2pry"
}
```

## Get chats

To get a list of all created chats with a particular bot, use the following request:

```nginx
GET /v1/client-api/chats
```

| Query parameter | Description                            |
| --------------- | -------------------------------------- |
| bot_id          | `Required` <br> Unique bot identifier. |

Returned response:

```json
[
    {
        "id": "01hc7aa947s4fdwywwjqr14wh7",
        "userId": "01hc71yz5pny5bmf0b4055zbjw",
        "botId": "01hc76p6ma76qfrs0bfr5x2pry"
    },
    {
        "id": "01hc7aa947s4fdwyww29pny5bm",
        "userId": "01hc71yz5pny5bmf0b4055zbjw",
        "botId": "01hc76p6ma76qfrs0bfr5x2pry"
    }
]
```

## Delete chat

You can delete a chat by its ID:

```nginx
DELETE /v1/client-api/chat/:id
```

## Send message

To send a message to the bot use the following request with JSON body:

```nginx
POST /v1/client-api/chat/message
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
    "senderId": "01hc71yz5pny5bmf0b4055zbjw",
    "body": "Hello World!",
    "timestamp": 1696758105336
}
```

## Get messages

To get a list of messages from the chat, use the following request:

```nginx
GET /v1/client-api/chat/:id/messages
```

| Query parameter | Description                                                                                                                                               |
| --------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------- |
| page            | `Optional` <br> Message page number. Default is 1.                                                                                                        |
| limit           | `Optional` <br> Maximum number of messages per page. Default is 20.                                                                                       |
| sender_id       | `Optional` <br> Filter messages by the ID of the specified sender. For example, you can filter messages to get only bot messages, ignoring user messages. |
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
            "senderId": "01hc71yz5pny5bmf0b4055zbjw",
            "body": "Hello World!",
            "timestamp": 1696758105336
        }
    ]
}
```

> The `total` field is an indicator of the number of all messages found.

## Upload files

To upload files use this request:

```nginx
POST /v1/client-api/files
```

with form data as a body:

```sh
Content-Disposition: form-data; name="file"; filename="yourfile.png"
Content-Type: image/png
# Content...
```

The parameter `name` is always should be equals to _file_.

> Unique file identifiers can be used as attachments for [messages](#send-message).
