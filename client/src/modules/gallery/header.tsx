import { Box, Container } from "@mantine/core";

export const Header = () => {
    return (
        <Box
            pos="sticky"
            w="100%"
            py="sm"
            h={40}
            sx={(theme) => ({
                background: theme.colorScheme === "dark" ? theme.colors.dark[6] : theme.colors.gray[2],
                zIndex: 10,
            })}
        >
            <Container size="md"></Container>
        </Box>
    );
};
