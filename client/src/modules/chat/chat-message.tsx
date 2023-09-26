import { Box, Flex, Avatar, TypographyStylesProvider, Text, MantineTheme, Space } from "@mantine/core";
import { FC } from "react";
import { File, Message } from "./types";
import { abbreviateName } from "../../helpers/text";
import { BOT_COMMAND_REGEX, KNOWN_MIME_TYPES, NOT_HTML_NEWLINE_REXEG } from "./const";
import { AudioGroup, ImageGroup } from "./file-groups";

const mockFiles: File[] = [
    { id: Math.random().toFixed(7), path: "https://placehold.co/900x400/orange/white", mimeType: "image/svg+xml" },
    { id: Math.random().toFixed(7), path: "https://placehold.co/700x400/blue/white", mimeType: "image/svg+xml" },
    { id: Math.random().toFixed(7), path: "https://placehold.co/400x400/green/white", mimeType: "image/svg+xml" },
    { id: Math.random().toFixed(7), path: "https://placehold.co/200x800/aqua/white", mimeType: "image/svg+xml" },
    { id: Math.random().toFixed(7), path: "https://placehold.co/100x200/pink/white", mimeType: "image/svg+xml" },
    { id: Math.random().toFixed(7), path: "https://placehold.co/1920x1080/black/white", mimeType: "image/svg+xml" },
    {
        id: Math.random().toFixed(7),
        path: "https://diviextended.com/wp-content/uploads/2021/10/sound-of-waves-marine-drive-mumbai.mp3",
        mimeType: "audio/mpeg",
    },
    {
        id: Math.random().toFixed(7),
        path: "https://diviextended.com/wp-content/uploads/2021/10/sound-of-waves-marine-drive-mumbai.mp3",
        mimeType: "audio/mpeg",
    },
];

type ChatMessageTypes = "bot" | "user";

interface ChatMessageProps {
    message: Message;
    type: ChatMessageTypes;
    senderName: string;
    avatar?: string;
}

export const ChatMessage: FC<ChatMessageProps> = (props) => {
    const defineBg = (theme: MantineTheme) => {
        return props.type === "user"
            ? "transparent"
            : theme.colorScheme === "dark"
            ? theme.colors.dark[6]
            : theme.colors.gray[2];
    };

    return (
        <Box
            p={window.innerWidth > 960 ? "md" : "sm"}
            sx={(theme) => ({
                borderRadius: window.innerWidth > 960 ? "0.4rem" : "none",
                background: defineBg(theme),
            })}
        >
            <Flex gap="sm" align="center" mb="0.2rem">
                <Avatar color={props.type === "user" ? "gray" : "primary"} size="md" src={props.avatar ?? null}>
                    {props.type === "bot" && !props.avatar ? abbreviateName(props.senderName) : null}
                </Avatar>

                <Flex direction="column">
                    <Text size="md" fw={600}>
                        {props.senderName}
                    </Text>

                    <Text opacity={0.5} size="xs">
                        {props.message.timestamp.toLocaleString()}
                    </Text>
                </Flex>
            </Flex>
            <TypographyStylesProvider
                sx={{
                    textAlign: "justify",
                    fontSize: window.innerWidth > 600 ? "1.1rem" : "1rem",
                    overflowWrap: "anywhere",
                }}
            >
                <div
                    dangerouslySetInnerHTML={{
                        __html: props.message.body
                            .replaceAll(NOT_HTML_NEWLINE_REXEG, "<br>")
                            .replaceAll("\t", "&emsp;")
                            .replaceAll(BOT_COMMAND_REGEX, (s) => `<a>${s}</a>`),
                    }}
                />
            </TypographyStylesProvider>
            {mockFiles.length > 0 ? (
                <>
                    <Space h="sm" />
                    <ImageGroup files={mockFiles.filter((file) => KNOWN_MIME_TYPES[file.mimeType] === "image")} />
                    <Space h="sm" />
                    <AudioGroup files={mockFiles.filter((file) => KNOWN_MIME_TYPES[file.mimeType] === "audio")} />
                </>
            ) : null}
        </Box>
    );
};
