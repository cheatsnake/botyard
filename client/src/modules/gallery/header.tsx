import { Box, Container, Flex } from "@mantine/core";
import { SettingsModal } from "../settings/settings-modal";

export const Header = () => {
    return (
        <Box
            pos="sticky"
            w="100%"
            py="sm"
            sx={(theme) => ({
                background: theme.colorScheme === "dark" ? theme.colors.dark[6] : theme.colors.gray[2],
                zIndex: 10,
            })}
        >
            <Container size="md">
                <Flex justify="end">
                    <SettingsModal />
                </Flex>
            </Container>
        </Box>
    );
};
