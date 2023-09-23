import { Avatar, Box, Divider, Flex, Modal, Text, Tooltip } from "@mantine/core";
import { useDisclosure } from "@mantine/hooks";
import { FC } from "react";
import { abbreviateName } from "../../helpers/test.helpers";
import { Bot } from "./types";
import { CopyBtn } from "../../components/CopyBtn";

export const BotInfoModal: FC<{ bot: Bot }> = ({ bot }) => {
    const [opened, { open, close }] = useDisclosure(false);

    return (
        <>
            <Modal opened={opened} onClose={close} title={<Text size="lg">Bot info</Text>}>
                <Modal.Body p={0} py="sm">
                    <Flex gap="md" align="start">
                        <Avatar color="primary" size="xl" src={bot.avatar ?? null}>
                            {!bot.avatar ? abbreviateName(bot.name) : null}
                        </Avatar>

                        <Box>
                            <Flex gap="sm">
                                <Text size="xl" fw={600}>
                                    {bot.name}
                                </Text>

                                <CopyBtn value={window.location.href} thing="Link" />
                            </Flex>

                            <Text opacity={0.7}>{`ID: ${bot.id}`}</Text>
                        </Box>
                    </Flex>

                    <Divider mt="lg" />

                    <Text component="pre" ta="justify" sx={{ whiteSpace: "pre-wrap", overflowWrap: "anywhere" }}>
                        {bot.description}
                    </Text>
                </Modal.Body>
            </Modal>
            <Tooltip label="Show bot info" openDelay={700}>
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
            </Tooltip>
        </>
    );
};
