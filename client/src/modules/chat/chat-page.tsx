import { Container, Flex, ScrollArea } from "@mantine/core";
import { ChatHeader } from "./chat-header";
import { useEffect, useRef, useState } from "react";
import { Message } from "./types";
import { ChatInput } from "./chat-input";
import { EmptyChatLabel } from "./empty-chat-label";
import { ChatMessage } from "./chat-message";

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
            <ChatHeader />
            <Container pos="relative" p="0" size="md" h="calc(100vh - 52px)">
                <Flex direction="column" justify="end" w="100%" h="100%">
                    <ScrollArea
                        pt="sm"
                        viewportRef={viewport}
                        h="auto"
                        sx={{ display: "flex", flexDirection: "column-reverse", justifyContent: "end" }}
                    >
                        {messages.map((msg, i) =>
                            i % 2 !== 0 ? (
                                <ChatMessage key={msg.id} message={msg} type="bot" senderName="Bot calculator" />
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
