export interface ResponseErr {
    message: string;
}

export interface ResponseOK {
    message: string;
}

export interface User {
    id: string;
    nickname: string;
}

export interface Bot {
    id: string;
    name: string;
    description?: string;
    avatar?: string;
}

export interface BotCommand {
    alias: string;
    description: string;
}

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
    timestamp: Date;
    files?: Attachment[];
}

export interface CreateMessageBody {
    chatId: string;
    senderId: string;
    body: string;
    fileIds?: string[];
}

export interface Attachment {
    id: string;
    path: string;
    name: string;
    size: number;
    mimeType: string;
}
