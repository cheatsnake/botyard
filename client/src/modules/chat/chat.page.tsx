import { ActionIcon, Avatar, Box, Container, Flex, ScrollArea, Text, Textarea, Tooltip } from "@mantine/core";
// import { Header } from "../../components/header";
import { IconLink, IconSend } from "@tabler/icons-react";
import { useEffect, useRef, useState } from "react";

export const ChatPage = () => {
    const [messages, setMessages] = useState([
        {
            id: 1,
            body: "Lorem ipsum dolor sit, amet consectetur adipisicing elit. Magnam totam exercitationem corporis perspiciatis debitis quo at nobis distinctio iure harum enim nam, officia in amet sunt quisquam tenetur eligendi ipsum!",
            timestamp: new Date(),
        },
        {
            id: 2,
            body: "Lorem ipsum dolor sit, amet consectetur adipisicing elit.",
            timestamp: new Date(),
        },
    ]);
    const [body, setBody] = useState("");

    const viewport = useRef<HTMLDivElement>(null);

    const scrollToBottom = () => {
        viewport.current!.scrollTo({ top: viewport.current!.scrollHeight });
    };

    useEffect(() => {
        scrollToBottom();
    }, []);

    return (
        <>
            {/* <Header /> */}
            <Container pos="relative" p="0" size="md" h="100vh">
                <Flex direction="column" justify="end" w="100%" h="100%">
                    <ScrollArea
                        viewportRef={viewport}
                        h="auto"
                        sx={{ display: "flex", flexDirection: "column-reverse", justifyContent: "end" }}
                    >
                        {messages.map((msg, i) => (
                            <Box
                                key={msg.id}
                                p={window.screen.width > 960 ? "md" : "sm"}
                                sx={(theme) => ({
                                    borderRadius: window.screen.width > 960 ? "0.4rem" : "none",
                                    background:
                                        theme.colorScheme === "dark"
                                            ? i % 2 === 0
                                                ? "transparent"
                                                : theme.colors.dark[6]
                                            : i % 2 === 0
                                            ? "transparent"
                                            : theme.colors.gray[2],
                                })}
                            >
                                {i % 2 !== 0 ? (
                                    <Flex gap="sm" align="center" mb="sm">
                                        <Avatar color="cyan" size="md">
                                            BC
                                        </Avatar>
                                        <Text size="lg" fw={600}>
                                            Bot calculator
                                        </Text>
                                    </Flex>
                                ) : null}
                                <Text ta="justify">{msg.body}</Text>
                                <Text ta="end" opacity={0.5}>
                                    {msg.timestamp.toLocaleTimeString()}
                                </Text>
                            </Box>
                        ))}
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
                                    }}
                                >
                                    <IconSend />
                                </ActionIcon>
                            </Tooltip>
                            <Tooltip openDelay={700} offset={15} label="Attach files">
                                <ActionIcon c="gray" size="lg" h="50%">
                                    <IconLink />
                                </ActionIcon>
                            </Tooltip>
                        </Flex>
                    </Flex>
                </Flex>
            </Container>
        </>
    );
};
