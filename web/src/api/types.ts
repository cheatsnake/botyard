export interface ResponseErr {
    message: string;
}

export interface ResponseOK {
    message: string;
}

export interface ServiceInfo {
    name: string;
    description: string;
    avatar?: string;
    socials: {
        title: string;
        url: string;
    }[];
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
    id: string;
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
    timestamp: number;
    attachments?: Attachment[];
}

export interface MessagePage {
    chatId: string;
    page: number;
    limit: number;
    total: number;
    messages: Message[];
}

export interface CreateMessageBody {
    chatId: string;
    senderId: string;
    body: string;
    attachmentIds?: string[];
}

export interface Attachment {
    id: string;
    path: string;
    name: string;
    size: number;
    mimeType: string;
}
