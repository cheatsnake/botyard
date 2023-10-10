# Client API

## Get service info

```nginx
GET /v1/client-api/service-info
```

## Get all bots

```nginx
GET /v1/client-api/bots
```

## Get bot

```nginx
GET /v1/client-api/bot/:id
```

## Create user (login)

```nginx
POST /v1/client-api/user
```

## Get current user

```nginx
GET /v1/client-api/user
```

## Get bot commands

```nginx
GET /v1/client-api/bot/:id/commands
```

## Create chat

```nginx
POST /v1/client-api/chat
```

## Get chats

```nginx
GET /v1/client-api/chats
```

## Delete chat

```nginx
DELETE /v1/client-api/chat/:id
```

## Send message

```nginx
POST /v1/client-api/chat/message
```

## Get messages

```nginx
GET /v1/client-api/chat/:id/messages
```

## Upload files

```nginx
POST /v1/client-api/files
```
