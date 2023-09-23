import { Box, Flex, Avatar, TypographyStylesProvider, Text, MantineTheme } from "@mantine/core";
import { FC } from "react";
import { Message } from "./types";
import { abbreviateName } from "../../helpers/test.helpers";

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
                        __html: props.message.body.replaceAll("\n", "<br>").replaceAll("\t", "&emsp;"),
                    }}
                />
            </TypographyStylesProvider>
        </Box>
    );
};
