import { Container, Flex, ScrollArea } from "@mantine/core";
import { ChatHeader } from "./chat-header";
import { useEffect, useRef, useState } from "react";
import { Bot, Message } from "./types";
import { ChatInput } from "./chat-input";
import { EmptyChatLabel } from "./empty-chat-label";
import { ChatMessage } from "./chat-message";
import { notifications } from "@mantine/notifications";
import { debounce } from "../../helpers/debounce";

const BOT: Bot = {
    id: "52h7GxjB2pd56pJqaDrsq3pZ5I",
    name: "Bot calculator",
    description:
        "Lorem ipsum dolor sit amet consectetur adipisicing elit. Exercitationem ratione neque mollitia labore consequatur officiis at ut molestias eos, saepe dolorem culpa harum veritatis, velit aut qui repudiandae maxime! Iure! Lorem ipsum dolor sit amet consectetur adipisicing elit.\n\nExercitationem ratione neque mollitia labore consequatur officiis at ut molestias eos, saepe dolorem culpa harum veritatis, velit aut qui repudiandae maxime! Iure!",
    avatar: "data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHZpZXdCb3g9IjAgMCAyNTAgMjUwIiB3aWR0aD0iMjUwIiBoZWlnaHQ9IjI1MCI+CiAgPHJlY3Qgd2lkdGg9IjI1MCIgaGVpZ2h0PSIyNTAiIGZpbGw9IiMwMDAwMDBGRiI+PC9yZWN0PgogIDx0ZXh0IHg9IjUwJSIgeT0iNTAlIiBkb21pbmFudC1iYXNlbGluZT0ibWlkZGxlIiB0ZXh0LWFuY2hvcj0ibWlkZGxlIiBmb250LWZhbWlseT0ibW9ub3NwYWNlIiBmb250LXNpemU9IjQ4cHgiIGZpbGw9IiNGRkZGRkZGRiI+Qk9UPC90ZXh0PiAgIAo8L3N2Zz4=",
};

const ChatPage = () => {
    const [messages, setMessages] = useState<Message[]>([]);
    const [body, setBody] = useState("");

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
                senderId: Math.random().toFixed(5),
                body: value ?? body,
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
        scrollToBottom();
    }, [messages]);

    useEffect(() => {
        chatViewport.current?.addEventListener("scroll", debounce(loadOldMessages));
    }, []);

    return (
        <>
            <ChatHeader bot={BOT} />
            <Container pos="relative" p={0} pt={54} size="md" h="100vh" ref={pageViewport}>
                <Flex direction="column" justify="end" w="100%" h="100%">
                    <ScrollArea
                        pt="sm"
                        viewportRef={chatViewport}
                        sx={{ display: "flex", flexDirection: "column-reverse", justifyContent: "end" }}
                    >
                        {messages.map((msg, i) =>
                            i % 2 !== 0 ? (
                                <ChatMessage
                                    key={msg.id}
                                    message={msg}
                                    type="bot"
                                    senderName={BOT.name}
                                    avatar={BOT.avatar}
                                />
                            ) : (
                                <ChatMessage key={msg.id} message={msg} type="user" senderName="You" />
                            )
                        )}
                    </ScrollArea>

                    {messages.length === 0 ? <EmptyChatLabel /> : null}

                    <ChatInput body={body} setBody={setBody} sendMessage={sendMessage} />
                </Flex>
            </Container>
        </>
    );
};

export default ChatPage;
