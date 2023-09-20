import { Avatar, Box, Flex, Text, TypographyStylesProvider } from "@mantine/core";
import { FC } from "react";
import { Message } from "./types";

export const BotMessage: FC<{ message: Message }> = (props) => {
    return (
        <Box
            p={window.screen.width > 960 ? "md" : "sm"}
            sx={(theme) => ({
                borderRadius: window.screen.width > 960 ? "0.4rem" : "none",
                background: theme.colorScheme === "dark" ? theme.colors.dark[6] : theme.colors.gray[2],
            })}
        >
            <Flex gap="sm" align="center" mb="0.2rem">
                <Avatar color="cyan" size="md">
                    BC
                </Avatar>
                <Flex direction="column">
                    <Text size="md" fw={600}>
                        Bot calculator
                    </Text>
                    <Text opacity={0.5} size="xs">
                        {props.message.timestamp.toLocaleString()}
                    </Text>
                </Flex>
            </Flex>
            <TypographyStylesProvider sx={{ textAlign: "justify", fontSize: "1rem" }}>
                <div dangerouslySetInnerHTML={{ __html: props.message.body.replaceAll("\n", "<br>") }} />
            </TypographyStylesProvider>
        </Box>
    );
};
