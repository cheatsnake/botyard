import { Flex, Box, Text, Kbd } from "@mantine/core";

export const EmptyChatLabel = () => {
    return (
        <Flex direction="column" justify="center" align="center" w="100%" h="100%">
            <Box
                maw={280}
                p="md"
                py="xl"
                sx={(theme) => ({
                    background: theme.colorScheme === "dark" ? theme.colors.dark[6] : theme.colors.gray[2],
                    borderRadius: "0.3rem",
                })}
            >
                <Text ta="center" opacity={0.5} size="lg" fw={500}>
                    Message history is empty
                </Text>
                <Text ta="center" mt="md" opacity={0.5}>
                    Start typing <Kbd>/</Kbd> or press <Kbd>Ctrl</Kbd>+<Kbd>K</Kbd> to view list of bot commands
                </Text>
            </Box>
        </Flex>
    );
};
