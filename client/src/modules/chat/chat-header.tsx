import { ActionIcon, Box, Container, Flex, Text, useMantineColorScheme } from "@mantine/core";
import { IconArrowNarrowLeft, IconSettings } from "@tabler/icons-react";
import { useNavigate } from "react-router-dom";

export const ChatHeader = () => {
    const { toggleColorScheme } = useMantineColorScheme();
    const navigate = useNavigate();

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
                <Flex justify="space-between">
                    <ActionIcon variant="subtle" size="md" onClick={() => navigate("/")}>
                        <IconArrowNarrowLeft />
                    </ActionIcon>
                    <Text size="lg" fw={600}>
                        Bot calculator
                    </Text>
                    <ActionIcon variant="subtle" size="md" onClick={() => toggleColorScheme()}>
                        <IconSettings />
                    </ActionIcon>
                </Flex>
            </Container>
        </Box>
    );
};
