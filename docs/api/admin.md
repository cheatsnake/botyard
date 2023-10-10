# Admin API

To manage bots, the platform provides a simple and convenient HTTP API for admins.

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
| `name`         | _Required_ <br> A name for the new bot.                       |
| `description`  | _Required_ <br> A brief description of what your bot will do. |
| `avatar`       | _Optional_ <br> Link to an image (avatar) for your bot.       |

As a result you will

## Delete a bot

```nginx
DELETE /v1/admin-api/bot/:id
```

## Get a bot's secret key

```nginx
GET /v1/admin-api/bot/:id/key
```

## Refresh a bot's secret key

```nginx
PUT /v1/admin-api/bot/:id/key
```

## Reload server config

```nginx
PUT /v1/admin-api/config
```
