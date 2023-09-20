import { ActionIcon, Box, Container, FileButton, Flex, ScrollArea, Text, Textarea, Tooltip } from "@mantine/core";
import { ChatHeader } from "./chat-header";
import { IconLink, IconSend } from "@tabler/icons-react";
import { useEffect, useRef, useState } from "react";
import { Message } from "./types";
import { BotMessage } from "./bot-message";
import { UserMessage } from "./user-message";

export const ChatPage = () => {
    const [messages, setMessages] = useState<Message[]>([]);
    const [body, setBody] = useState("");

    const viewport = useRef<HTMLDivElement>(null);

    const scrollToBottom = () => {
        viewport.current!.scrollTo({ top: viewport.current!.scrollHeight });
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
                                <BotMessage key={msg.id} message={msg} />
                            ) : (
                                <UserMessage key={msg.id} message={msg} />
                            )
                        )}
                    </ScrollArea>

                    {messages.length === 0 ? (
                        <Flex direction="column" justify="center" align="center" w="100%" h="100%">
                            <Box
                                p="md"
                                sx={(theme) => ({
                                    background:
                                        theme.colorScheme === "dark" ? theme.colors.dark[6] : theme.colors.gray[2],
                                    borderRadius: "0.25rem",
                                })}
                            >
                                <Text ta="center" opacity={0.5} size="lg" fw={500}>
                                    Message history is empty
                                </Text>
                                <Text opacity={0.5}>Start typing "/" to view list of commands.</Text>
                            </Box>
                        </Flex>
                    ) : null}
                    <Flex align="center" gap="sm" p="sm" pt="lg">
                        <Textarea
                            value={body}
                            onChange={(event) => setBody(event.currentTarget.value)}
                            w="100%"
                            placeholder="Your message..."
                            minRows={2}
                            maxRows={6}
                            autosize
                        />
                        <Flex direction="column" gap="5px" h="100%">
                            <Tooltip openDelay={700} offset={15} label="Send message">
                                <ActionIcon
                                    disabled={body.length === 0}
                                    size="lg"
                                    h="50%"
                                    onClick={() => {
                                        setMessages((prev) => [
                                            ...prev,
                                            { id: Math.random(), body, timestamp: new Date() },
                                        ]);
                                        scrollToBottom();
                                        setBody("");
                                    }}
                                >
                                    <IconSend />
                                </ActionIcon>
                            </Tooltip>
                            <Tooltip openDelay={700} offset={15} label="Attach files">
                                <FileButton onChange={() => {}} multiple>
                                    {(props) => (
                                        <ActionIcon {...props} c="gray" size="lg" h="50%">
                                            <IconLink />
                                        </ActionIcon>
                                    )}
                                </FileButton>
                            </Tooltip>
                        </Flex>
                    </Flex>
                </Flex>
            </Container>
        </>
    );
};
