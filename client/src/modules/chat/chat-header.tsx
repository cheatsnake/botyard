import { ActionIcon, Box, Container, Flex } from "@mantine/core";
import { IconArrowNarrowLeft } from "@tabler/icons-react";
import { useNavigate } from "react-router-dom";
import { BotInfoModal } from "./bot-info-modal";
import { FC } from "react";
import { Bot } from "./types";
import { SettingsModal } from "../settings/settings-modal";

export const ChatHeader: FC<{ bot: Bot }> = ({ bot }) => {
    const navigate = useNavigate();

    return (
        <Box
            pos="fixed"
            w="100%"
            py="sm"
            sx={(theme) => ({
                background: theme.colorScheme === "dark" ? theme.colors.dark[6] : theme.colors.gray[2],
                zIndex: 10,
            })}
        >
            <Container size="md">
                <Flex justify="space-between">
                    <ActionIcon
                        variant="subtle"
                        size="md"
                        onClick={() => {
                            if ("startViewTransition" in document) {
                                //@ts-ignore
                                return document.startViewTransition(() => {
                                    navigate("/");
                                });
                            }

                            return navigate("/");
                        }}
                    >
                        <IconArrowNarrowLeft />
                    </ActionIcon>

                    <BotInfoModal bot={bot} />

                    <SettingsModal />
                </Flex>
            </Container>
        </Box>
    );
};
