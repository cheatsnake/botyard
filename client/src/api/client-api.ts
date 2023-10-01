import { jsonRequestParams, queryParams } from "./helpers";
import {
    Bot,
    BotCommand,
    Chat,
    CreateMessageBody,
    ResponseErr,
    Message,
    ResponseOK,
    User,
    Attachment,
    ServiceInfo,
    MessagePage,
} from "./types";

const CLIENT_API_PREFIX = "/v1/client-api";

class ClientAPI {
    constructor(private prefix: string) {
        this.prefix = prefix;
    }

    async getServiceInfo() {
        const resp = await fetch(this.prefix + "/service-info");

        if (!resp.ok) {
            const body: ResponseErr = await resp.json();
            throw new Error(body.message);
        }

        const info: ServiceInfo = await resp.json();
        return info;
    }

    async getCurrentUser() {
        const resp = await fetch(this.prefix + "/user");

        if (!resp.ok) {
            return;
        }

        const user: User = await resp.json();
        return user;
    }

    async createUser(nickname: string) {
        const resp = await fetch(this.prefix + "/user", jsonRequestParams("POST", { nickname }));

        if (!resp.ok) {
            const body: ResponseErr = await resp.json();
            throw new Error(body.message);
        }

        const user: User = await resp.json();
        return user;
    }

    async getAllBots() {
        const resp = await fetch(this.prefix + "/bots");

        if (!resp.ok) {
            const body: ResponseErr = await resp.json();
            throw new Error(body.message);
        }

        const bots: Bot[] = await resp.json();
        return bots;
    }

    async getBot(id: string) {
        const resp = await fetch(this.prefix + `/bot/${id}`);

        if (!resp.ok) {
            const body: ResponseErr = await resp.json();
            throw new Error(body.message);
        }

        const bot: Bot = await resp.json();
        return bot;
    }

    async getBotCommands(botId: string) {
        const resp = await fetch(this.prefix + `/bot/${botId}/commands`);

        if (!resp.ok) {
            const body: ResponseErr = await resp.json();
            throw new Error(body.message);
        }

        const cmds: BotCommand[] = await resp.json();
        return cmds;
    }

    async getChatsByBot(botId: string) {
        const resp = await fetch(this.prefix + `/chats?${queryParams({ bot_id: botId })}`);

        if (!resp.ok) {
            const body: ResponseErr = await resp.json();
            throw new Error(body.message);
        }

        const chats: Chat[] = await resp.json();
        return chats;
    }

    async getMessagesByChat(chatId: string, senderId = "", page = 1, limit = 20) {
        const resp = await fetch(
            this.prefix + `/chat/${chatId}/messages?${queryParams({ sender_id: senderId, page, limit })}`
        );

        if (!resp.ok) {
            const body: ResponseErr = await resp.json();
            throw new Error(body.message);
        }

        const msgs: MessagePage = await resp.json();
        return msgs;
    }

    async createChat(botId: string) {
        const resp = await fetch(this.prefix + "/chat", jsonRequestParams("POST", { botId }));

        if (!resp.ok) {
            const body: ResponseErr = await resp.json();
            throw new Error(body.message);
        }

        const chat: Chat = await resp.json();
        return chat;
    }

    async sendUserMessage(body: CreateMessageBody) {
        const resp = await fetch(this.prefix + "/chat/message", jsonRequestParams("POST", { ...body }));

        if (!resp.ok) {
            const body: ResponseErr = await resp.json();
            throw new Error(body.message);
        }

        const msg: Message = await resp.json();
        return msg;
    }

    async deleteChat(id: string) {
        const resp = await fetch(this.prefix + `/chat/${id}`, { method: "DELETE" });

        if (!resp.ok) {
            const body: ResponseErr = await resp.json();
            throw new Error(body.message);
        }

        const result: ResponseOK = await resp.json();
        return result;
    }

    async uploadFiles(files: File[]) {
        if (files.length === 0) {
            throw new Error("no files to upload");
        }

        const formData = new FormData();
        for (const file of files) {
            formData.append("file", file);
        }

        const resp = await fetch(this.prefix + "/files", { method: "POST", body: formData });
        if (!resp.ok) {
            const body: ResponseErr = await resp.json();
            throw new Error(body.message);
        }

        const attachments: Attachment[] = await resp.json();
        return attachments;
    }
}

export default new ClientAPI(CLIENT_API_PREFIX);
