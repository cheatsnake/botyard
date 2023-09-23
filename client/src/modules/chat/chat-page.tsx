import { Container, Flex, ScrollArea } from "@mantine/core";
import { ChatHeader } from "./chat-header";
import { useEffect, useRef, useState } from "react";
import { Bot, Message } from "./types";
import { ChatInput } from "./chat-input";
import { EmptyChatLabel } from "./empty-chat-label";
import { ChatMessage } from "./chat-message";

const BOT: Bot = {
    id: "52h7GxjB2pd56pJqaDrsq3pZ5I",
    name: "Bot calculator",
    description:
        "Lorem ipsum dolor sit amet consectetur adipisicing elit. Exercitationem ratione neque mollitia labore consequatur officiis at ut molestias eos, saepe dolorem culpa harum veritatis, velit aut qui repudiandae maxime! Iure! Lorem ipsum dolor sit amet consectetur adipisicing elit. Exercitationem ratione neque mollitia labore consequatur officiis at ut molestias eos, saepe dolorem culpa harum veritatis, velit aut qui repudiandae maxime! Iure!",
    avatar: "data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHZpZXdCb3g9IjAgMCAyNTAgMjUwIiB3aWR0aD0iMjUwIiBoZWlnaHQ9IjI1MCI+CiAgPHJlY3Qgd2lkdGg9IjI1MCIgaGVpZ2h0PSIyNTAiIGZpbGw9IiMwMDAwMDBGRiI+PC9yZWN0PgogIDx0ZXh0IHg9IjUwJSIgeT0iNTAlIiBkb21pbmFudC1iYXNlbGluZT0ibWlkZGxlIiB0ZXh0LWFuY2hvcj0ibWlkZGxlIiBmb250LWZhbWlseT0ibW9ub3NwYWNlIiBmb250LXNpemU9IjQ4cHgiIGZpbGw9IiNGRkZGRkZGRiI+Qk9UPC90ZXh0PiAgIAo8L3N2Zz4=",
};

export const ChatPage = () => {
    const [messages, setMessages] = useState<Message[]>([]);
    const [body, setBody] = useState("");

    const viewport = useRef<HTMLDivElement>(null);

    const scrollToBottom = () => {
        viewport.current!.scrollTo({ top: viewport.current!.scrollHeight, behavior: "smooth" });
    };

    const sendMessage = () => {
        setMessages((prev) => [...prev, { id: Math.random(), body, timestamp: new Date() }]);
        scrollToBottom();
        setBody("");
    };

    useEffect(() => {
        scrollToBottom();
    }, [messages]);

    return (
        <>
            <ChatHeader bot={BOT} />
            <Container pos="relative" p="0" size="md" h="calc(100vh - 54px)">
                <Flex direction="column" justify="end" w="100%" h="100%">
                    <ScrollArea
                        pt="sm"
                        viewportRef={viewport}
                        h="auto"
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
