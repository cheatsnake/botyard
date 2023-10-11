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
        "alias": "start",
        "description": "Start a new conversation."
    },
    {
        "alias": "help",
        "description": "Show usage guide."
    }
]
```

## Remove bot command

```nginx
DELETE /v1/bot-api/bot/command
```

## Add bot webhook

```nginx
POST /v1/bot-api/bot/webhook
```

## Get bot webhook

```nginx
GET /v1/bot-api/bot/webhook
```

## Delete bot webhook

```nginx
DELETE /v1/bot-api/bot/webhook
```

## Get chats

```nginx
GET /v1/bot-api/bot/chats
```

## Send message

```nginx
POST /v1/bot-api/bot/chat/message
```

## Get messages

```nginx
GET /v1/bot-api/bot/chat/:id/messages
```

## Get message

```nginx
GET /v1/bot-api/bot/chat/message/:id
```

## Upload files

```nginx
POST /v1/bot-api/bot/files
```
