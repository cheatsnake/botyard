# Configuration

## Config file

The main configuration file called `botyard.config.json` and located in the root of the project and looks like this:

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
            "maxNicknameLength": 32
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

## Environment variables
