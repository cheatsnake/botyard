# Bot API

This section describes the functionality of bots, which can be easily controlled using the HTTP API using almost any programming language.

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

This request adds new commands to the current bot.

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

| Body parameter | Description                                                                                                                                                                        |
| -------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| url            | `Required` <br> The address of the webhook to which user messages will be sent. Requests will come with a POST method and a message as described [here](./client.md#send-message). |
| secret         | `Optional` <br> Random data, in order to be able to verify that queries are coming from the correct source.                                                                        |

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

Removes the webhook information for the current bot, thus stopping user messages being sent to it:

```nginx
DELETE /v1/bot-api/bot/webhook
```

## Get chats

```nginx
GET /v1/bot-api/chats
```

## Send message

```nginx
POST /v1/bot-api/chat/message
```

## Get messages

```nginx
GET /v1/bot-api/chat/:id/messages
```

## Get message

```nginx
GET /v1/bot-api/chat/message/:id
```

## Upload files

```nginx
POST /v1/bot-api/files
```
