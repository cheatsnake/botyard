import { Flex, Box, Text } from "@mantine/core";

export const EmptyChatLabel = () => {
    return (
        <Flex direction="column" justify="center" align="center" w="100%" h="100%">
            <Box
                p="md"
                sx={(theme) => ({
                    background: theme.colorScheme === "dark" ? theme.colors.dark[6] : theme.colors.gray[2],
                    borderRadius: "0.25rem",
                })}
            >
                <Text ta="center" opacity={0.5} size="lg" fw={500}>
                    Message history is empty
                </Text>
                <Text opacity={0.5}>Start typing "/" to view list of commands.</Text>
            </Box>
        </Flex>
    );
};
