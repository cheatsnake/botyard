# Configuration

## Config file

The main configuration file called `botyard.config.json`. It located in the `config` folder and looks like this:

```json
{
    "service": {
        "name": "Service name",
        "description": "Some info about your service.",
        "avatar": "https://google.com/image.png",
        "socials": [
            { "title": "Website", "url": "https://example.com" },
            { "title": "GitHub", "url": "https://github.com" },
            { "title": "Discord", "url": "https://discord.com" },
            { "title": "Twitter", "url": "https://twitter.com" }
        ]
    },
    "limits": {
        "user": {
            "minNicknameLength": 3,
            "maxNicknameLength": 32,
            "authTokenLifetime": 10080
        },
        "message": {
            "maxBodyLength": 4096,
            "maxAttachedFiles": 10
        },
        "file": {
            "maxImageSize": 2097152,
            "maxAudioSize": 5242880,
            "maxVideoSize": 26214400,
            "maxFileSize": 10485760
        }
    }
}
```

Below is a brief description for each field:

-   **service** - Basic information about your organization providing bots.
-   **service.name** - Name of your organization.
-   **service.description** - A brief description for your bots provided by your organization.
-   **service.avatar** - Link to your organization's logo or avatar.
-   **service.social** - Array of links to social networks and sites related to your organization.
-   **service.social.title** - Name of the social network or website (can be any).
-   **service.social.url** - Link to a social network profile or website.
-   **limits** - Data about limits on the platform.
-   **limits.user** - Data on limits that relate directly to users.
-   **limits.user.minNicknameLength** - Minimum length for user nickname.
-   **limits.user.maxNicknameLength** - Maximum length for user nickname.
-   **limits.user.authTokenLifetime** - The lifetime (in minutes) of the user's authorization token. You can learn more about it [here](./api/client.md#create-user-login). By default, it is equal to 10080 minutes, which corresponds to 1 week.
-   **limits.message** - Data on limits that relate directly to messages.
-   **limits.message.maxBodyLength** - The maximum length of the text part of the message (in characters).
-   **limits.message.maxAttachedFiles** - The maximum number of files that can be attached to a single message.
-   **limits.file** - Limit data that relates directly to uploaded files.
-   **limits.file.maxImageSize** - Maximum size (in bytes) for images. By default, the limit is 2 MB.
-   **limits.file.maxAudioSize** - Maximum size (in bytes) for audios. By default, the limit is 5 MB.
-   **limits.file.maxVideoSize** - Maximum size (in bytes) for videos. By default, the limit is 25 MB.
-   **limits.file.maxFileSize** - Maximum size (in bytes) for all other file types. By default, the limit is 10 MB.

## Environment variables

An example of `.env` file with environment variables required for the platform is shown below:

```
ADMIN_SECRET_KEY=YOUR_SECRET_KEY
JWT_SECRET_KEY=OTHER_SECRET_KEY
```

-   **ADMIN_SECRET_KEY** - It is a secret key for access to the [admin functionality](./api/admin.md).
-   **JWT_SECRET_KEY** - A random key required to sign [JWT](https://jwt.io) keys used to authorise users on the platform.

> Please do not specify keys that are too simple, it is not secure. Use random password generators to create strong keys.
