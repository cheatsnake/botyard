import { ActionIcon, Box, Container, useMantineColorScheme } from "@mantine/core";
import { IconMoonStars, IconSun } from "@tabler/icons-react";

export const Header = () => {
    const { colorScheme, toggleColorScheme } = useMantineColorScheme();
    const dark = colorScheme === "dark";

    return (
        <Box
            w="100%"
            py="sm"
            sx={(theme) => ({
                background: theme.colorScheme === "dark" ? theme.colors.dark[6] : theme.colors.gray[1],
                zIndex: 10,
            })}
        >
            <Container size="md">
                <ActionIcon
                    variant="outline"
                    // color={dark ? "yellow" : "blue"}
                    onClick={() => toggleColorScheme()}
                    title="Toggle color scheme"
                >
                    {dark ? <IconSun size="1.1rem" /> : <IconMoonStars size="1.1rem" />}
                </ActionIcon>
            </Container>
        </Box>
    );
};
