import { Avatar, Box, CopyButton, Divider, Flex, Modal, Text } from "@mantine/core";
import { useDisclosure } from "@mantine/hooks";
import { FC } from "react";
import { abbreviateName } from "../../helpers/test.helpers";
import { Bot } from "./types";

export const BotInfoModal: FC<{ bot: Bot }> = ({ bot }) => {
    const [opened, { open, close }] = useDisclosure(false);

    return (
        <>
            <Modal opened={opened} onClose={close} title={<Text size="lg">Bot info</Text>}>
                <Modal.Body p={0} py="sm">
                    <Flex gap="md" align="center">
                        <Avatar color="primary" size="xl" src={bot.avatar ?? null}>
                            {!bot.avatar ? abbreviateName(bot.name) : null}
                        </Avatar>

                        <Box>
                            <Text size="xl" fw={600}>
                                {bot.name}
                            </Text>

                            <Text opacity={0.7}>{`ID: ${bot.id}`}</Text>
                            <CopyButton value={window.location.href}>
                                {({ copied, copy }) => (
                                    <Text td={copied ? "none" : "underline"} onClick={copy} sx={{ cursor: "pointer" }}>
                                        {copied ? "Link copied" : "Copy link"}
                                    </Text>
                                )}
                            </CopyButton>
                        </Box>
                    </Flex>

                    <Divider mt="lg" />

                    <Text component="pre" ta="justify" sx={{ whiteSpace: "pre-wrap", overflowWrap: "anywhere" }}>
                        {bot.description}
                    </Text>
                </Modal.Body>
            </Modal>
            <Text
                component="button"
                onClick={open}
                size="lg"
                fw={600}
                bg="transparent"
                sx={{ border: "none", cursor: "pointer" }}
            >
                {bot.name}
            </Text>
        </>
    );
};
