import { Button, Container, Flex, ScrollArea } from "@mantine/core";
import { ChatHeader } from "./chat-header";
import { useEffect, useRef, useState } from "react";
import { Bot, Chat, Message } from "../../api/types";
import { ChatInput } from "./chat-input";
import { EmptyChatLabel } from "./empty-chat-label";
import { ChatMessage } from "./chat-message";
import { useLoaderContext } from "../../contexts/loader-context";
import { errNotify } from "../../helpers/notifications";
import { useStorageContext } from "../../contexts/storage-context";
import { useNavigate } from "react-router-dom";
import { useUserContext } from "../../contexts/user-context";
import ClientAPI from "../../api/client-api";
import { debounce } from "../../helpers/debounce";
import { defineNextPageLimit } from "../../helpers/pagination";

const MAX_PAGE_LIMIT = 20;

const ChatPage = () => {
    const [currentBot, setCurrentBot] = useState<Bot>({ id: "", name: "" });
    const [currentChat, setCurrentChat] = useState<Chat>();
    const [messages, setMessages] = useState<Message[]>([]);
    const [attachments, setAttachments] = useState<File[]>([]);
    const [body, setBody] = useState("");
    const [hasOldMessages, setHasOldMessages] = useState(true);
    const [hasScroll, setHasScroll] = useState(false);
    const [isBlockInput, setIsBlockInput] = useState(false);

    const { isLoad, setIsLoad } = useLoaderContext();
    const { loadBots } = useStorageContext();
    const { user } = useUserContext();
    const navigate = useNavigate();
    const chatViewport = useRef<HTMLDivElement>(null);

    const scrollToBottom = () => {
        chatViewport.current!.scrollTo({ top: chatViewport.current!.scrollHeight, behavior: "smooth" });
    };

    const loadMessages = async (chatId: string) => {
        try {
            if (!hasOldMessages) return;

            setIsLoad(true);
            const { page, limit } = defineNextPageLimit(messages.length, MAX_PAGE_LIMIT);
            const msgPage = await ClientAPI.getMessagesByChat(chatId, "", page, limit);

            if (msgPage.total <= messages.length + msgPage.messages.length) setHasOldMessages(false);
            if (msgPage.messages.length > 0) setMessages([...msgPage.messages, ...messages]);
        } catch (error) {
            errNotify((error as Error).message);
        } finally {
            setIsLoad(false);
        }
    };

    const loadNewPage = async () => {
        if (!currentChat?.id) return;
        const prev = chatViewport.current!.scrollHeight;

        await loadMessages(currentChat.id);

        setTimeout(() => {
            chatViewport.current!.scrollTo({
                top: chatViewport.current!.scrollHeight - prev,
            });
        }, 1);
    };

    useEffect(() => {
        (async () => {
            try {
                setIsLoad(true);

                const bots = await loadBots();
                const parts = window.location.pathname.split("/");
                const botId = parts[parts.length - 1];
                const cb = bots?.find((b) => b.id === botId);

                if (!cb) return navigate("/");

                let chats = await ClientAPI.getChatsByBot(cb.id);
                if (chats.length === 0) {
                    const newChat = await ClientAPI.createChat(cb.id);
                    chats = [newChat];
                }

                await loadMessages(chats[0].id);
                setCurrentBot(cb);
                setCurrentChat(chats[0]);
            } catch (error) {
                errNotify((error as Error).message);
            } finally {
                setIsLoad(false);
                setTimeout(scrollToBottom, 1);
            }
        })();
    }, []);

    return (
        <>
            <ChatHeader bot={currentBot} />
            <Container pos="relative" p={0} pt={54} size="md" h="100vh">
                <Flex direction="column" justify="end" w="100%" h="100%">
                    <ScrollArea
                        pt="sm"
                        onScrollPositionChange={debounce(() => {
                            if (!hasScroll) setHasScroll(true);
                        }, 100)}
                        viewportRef={chatViewport}
                        sx={{ display: "flex", flexDirection: "column-reverse", justifyContent: "end" }}
                    >
                        {hasScroll && hasOldMessages ? (
                            <Flex py="sm" justify="center">
                                <Button onClick={loadNewPage} variant="light" size="xs" color="gray">
                                    Load more
                                </Button>
                            </Flex>
                        ) : null}

                        {messages.map((msg) => (
                            <ChatMessage
                                key={msg.id}
                                message={msg}
                                type={msg.senderId === user?.id ? "user" : "bot"}
                                senderName={msg.senderId === user?.id ? user.nickname : currentBot.name}
                                senderAvatar={msg.senderId === user?.id ? "" : currentBot.avatar}
                                attachments={msg.attachments || []}
                            />
                        ))}
                    </ScrollArea>

                    {messages.length === 0 && !isLoad ? <EmptyChatLabel /> : null}

                    <ChatInput
                        body={body}
                        attachments={attachments}
                        currentChat={currentChat}
                        currentBot={currentBot}
                        isBlockInput={isBlockInput}
                        setBody={setBody}
                        setMessages={setMessages}
                        setAttachments={setAttachments}
                        setIsBlockInput={setIsBlockInput}
                        scrollToBottom={scrollToBottom}
                    />
                </Flex>
            </Container>
        </>
    );
};

export default ChatPage;
