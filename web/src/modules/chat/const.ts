import { AttachmentType } from "./types";

export const BOT_COMMAND_REGEX = /\/[a-z0-9_\-]+(?= |$)/gm;
export const NOT_HTML_NEWLINE_REXEG = /^(?![^<]*<\/[^>]*>)\n/gm;

export const KNOWN_MIME_TYPES: { [key: string]: AttachmentType } = {
    "image/gif": "image",
    "image/jpeg": "image",
    "image/png": "image",
    "image/svg+xml": "image",
    "image/webp": "image",

    "video/mp4": "video",
    "video/webm": "video",
    "video/ogg": "video",
    "video/quicktime": "video",
    "video/x-flv": "video",

    "audio/mpeg": "audio",
    "audio/ogg": "audio",
    "audio/wav": "audio",
    "audio/aac": "audio",
};
