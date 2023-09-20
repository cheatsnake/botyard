import { Box, Text, TypographyStylesProvider } from "@mantine/core";
import { FC } from "react";
import { Message } from "./types";

export const UserMessage: FC<{ message: Message }> = (props) => {
    return (
        <Box
            p={window.screen.width > 960 ? "md" : "sm"}
            sx={{
                borderRadius: window.screen.width > 960 ? "0.4rem" : "none",
                background: "transparent",
            }}
        >
            <Box mb="0.2rem">
                <Text size="md" fw={600}>
                    You
                </Text>
                <Text opacity={0.5} size="xs">
                    {props.message.timestamp.toLocaleString()}
                </Text>
            </Box>
            <TypographyStylesProvider>
                <div dangerouslySetInnerHTML={{ __html: props.message.body.replaceAll("\n", "<br>") }} />
            </TypographyStylesProvider>
        </Box>
    );
};
