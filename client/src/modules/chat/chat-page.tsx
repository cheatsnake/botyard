import { Container, Flex, ScrollArea } from "@mantine/core";
import { ChatHeader } from "./chat-header";
import { useEffect, useRef, useState } from "react";
import { Bot, Message } from "../../api/types";
import { ChatInput } from "./chat-input";
import { EmptyChatLabel } from "./empty-chat-label";
import { ChatMessage } from "./chat-message";
import { notifications } from "@mantine/notifications";
import { debounce } from "../../helpers/debounce";
import { useLoaderContext } from "../../contexts/loader-context";
import { errNotify } from "../../helpers/notifications";
import { useStorageContext } from "../../contexts/storage-context";
import { useNavigate } from "react-router-dom";
import { useUserContext } from "../../contexts/user-context";

const ChatPage = () => {
    const [currentBot, setCurrentBot] = useState<Bot>({ id: "", name: "" });
    const [messages, setMessages] = useState<Message[]>([]);
    const [body, setBody] = useState("");

    const navigate = useNavigate();
    const { setIsLoad } = useLoaderContext();
    const { loadBots } = useStorageContext();
    const { user } = useUserContext();

    const chatViewport = useRef<HTMLDivElement>(null);
    const pageViewport = useRef<HTMLDivElement>(null);

    const scrollToBottom = () => {
        chatViewport.current!.scrollTo({ top: chatViewport.current!.scrollHeight, behavior: "smooth" });
        pageViewport.current!.scrollTo({ top: chatViewport.current!.scrollHeight });
    };

    const sendMessage = (value?: string) => {
        setMessages((prev) => [
            ...prev,
            {
                id: Math.random().toFixed(5),
                chatId: Math.random().toFixed(5),
                senderId: user?.id || "",
                body: value ?? body,
                attachments: [],
                timestamp: new Date(),
            },
        ]);
        scrollToBottom();
        setBody("");
    };

    const loadOldMessages = () => {
        if (chatViewport.current!.scrollTop < 200) {
            notifications.show({
                withBorder: true,
                title: "Loading older messages...",
                color: "green",
                loading: true,
                autoClose: 5000,
                message: "",
            });
        }
    };

    useEffect(() => {
        chatViewport.current?.addEventListener("scroll", debounce(loadOldMessages));

        (async () => {
            try {
                setIsLoad(true);
                const bots = await loadBots();

                const parts = window.location.pathname.split("/");
                const botId = parts[parts.length - 1];
                const cb = bots?.find((b) => b.id === botId);

                if (!cb) {
                    navigate("/");
                    return;
                }

                setCurrentBot(cb);
            } catch (error) {
                errNotify((error as Error).message);
            } finally {
                setIsLoad(false);
            }
        })();
    }, []);

    useEffect(scrollToBottom, [messages]);

    return (
        <>
            <ChatHeader bot={currentBot} />
            <Container pos="relative" p={0} pt={54} size="md" h="100vh" ref={pageViewport}>
                <Flex direction="column" justify="end" w="100%" h="100%">
                    <ScrollArea
                        pt="sm"
                        viewportRef={chatViewport}
                        sx={{ display: "flex", flexDirection: "column-reverse", justifyContent: "end" }}
                    >
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

                    {messages.length === 0 ? <EmptyChatLabel /> : null}

                    <ChatInput body={body} setBody={setBody} sendMessage={sendMessage} />
                </Flex>
            </Container>
        </>
    );
};

export default ChatPage;
