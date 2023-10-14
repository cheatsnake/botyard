# Admin API

To manage bots, the platform provides a simple and convenient HTTP API for admins.

## Authorization

To access the admin functionality, a secret key is required, which is stored on the server in a `ADMIN_SECRET_KEY` [environment variable](../configuration.md#environment-variables). All requests described below should be executed with the [Authorization header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Authorization):

```nginx
Authorization: <YOUR_ADMIN_SECRET_KEY>
```

## Create a bot

The following route with a body in JSON format is used to create a new bot:

```nginx
POST /v1/admin-api/bot
```

```json
{
    "name": "Bot name",
    "description": "Just some info about bot.",
    "avatar": "https://google.com/image.png"
}
```

| Body parameter | Description                                                   |
| -------------- | ------------------------------------------------------------- |
| name           | `Required` <br> A name for the new bot.                       |
| description    | `Required` <br> A brief description of what your bot will do. |
| avatar         | `Optional` <br> Link to an image (avatar) for your bot.       |

As a result you will get response like this:

```json
{
    "bot": {
        "id": "01hc76p6ma76qfrs0bfr5x2pry",
        "name": "Test bot",
        "description": "Just testing bot.",
        "avatar": "https://google.com/image.png"
    },
    "key": "01hc76p6ma76qfrs0bfr5x2pry:hNBKVJQLAcD2S7xJMnVSkzKkfi"
}
```

-   `bot.id` a unique identifier that is required to perform other operations on this bot.
-   `key` is access key, which is required to [control](./bot.md#authorization) the new bot.

## Delete a bot

To remove a bot uses the following request:

```nginx
DELETE /v1/admin-api/bot/:id
```

## Get a bot's access key

To get the bot access key, use:

```nginx
GET /v1/admin-api/bot/:id/key
```

As a result you will get response like this:

```json
{
    "key": "01hc76p6ma76qfrs0bfr5x2pry:hNBKVJQLAcD2S7xJMnVSkzKkfi"
}
```

## Refresh a bot's access key

To recreate the access key and update its value, use:

```nginx
PUT /v1/admin-api/bot/:id/key
```

As a result you will get response like this:

```json
{
    "key": "01hc76p6ma76qfrs0bfr5x2pry:s0PWbovTQpmuCLKKaxuruOleaX"
}
```

## Reload server config

To apply the changes made to [config file](../configuration.md#config-file) without rebooting the server, use the following request:

```nginx
PUT /v1/admin-api/config
```
