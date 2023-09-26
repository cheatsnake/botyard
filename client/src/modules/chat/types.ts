export interface Chat {
    id: string;
    userId: string;
    botId: string;
}

export interface Message {
    id: string;
    chatId: string;
    senderId: string;
    body: string;
    fileIds?: File[];
    timestamp: Date;
}

export interface File {
    id: string;
    path: string;
    mimeType: string;
}

export interface Bot {
    id: string;
    name: string;
    description?: string;
    avatar?: string;
}

export interface Command {
    alias: string;
    description: string;
}

export type AttachmentType = "image" | "video" | "audio" | "file";
