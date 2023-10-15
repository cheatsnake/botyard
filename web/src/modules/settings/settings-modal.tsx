import {
    Modal,
    ActionIcon,
    useMantineColorScheme,
    Text,
    Flex,
    Switch,
    Tabs,
    Kbd,
    Box,
    useMantineTheme,
    Group,
    ColorSwatch,
} from "@mantine/core";
import { useDisclosure } from "@mantine/hooks";
import { IconSettings } from "@tabler/icons-react";
import { useState } from "react";
import { usePrimaryColorContext } from "../../contexts/primary-color";

export const SettingsModal = () => {
    const [opened, { open, close }] = useDisclosure(false);

    return (
        <>
            <Modal
                centered
                opened={opened}
                onClose={close}
                size="lg"
                title={
                    <Text size="xl" fw={700}>
                        Settings
                    </Text>
                }
            >
                <Modal.Body mih={"25vh"} p={0}>
                    <SettingsTabs />
                </Modal.Body>
            </Modal>

            <ActionIcon variant="subtle" size="md" onClick={open}>
                <IconSettings />
            </ActionIcon>
        </>
    );
};

const SettingsTabs = () => {
    const theme = useMantineTheme();
    return (
        <Tabs defaultValue="appearance" p={0} color={theme.primaryColor}>
            <Tabs.List>
                <Tabs.Tab value="appearance">Appearance</Tabs.Tab>
                <Tabs.Tab value="shortcuts">Shortcuts</Tabs.Tab>
            </Tabs.List>

            <Box px={window.innerWidth > 500 ? "sm" : 0}>
                <AppearanceTab />
                <ShortcutsTab />
            </Box>
        </Tabs>
    );
};

const AppearanceTab = () => {
    const theme = useMantineTheme();
    const { setPrimaryColor } = usePrimaryColorContext();
    const { colorScheme, toggleColorScheme } = useMantineColorScheme();
    const [isDark, setIsDark] = useState(colorScheme === "dark");

    return (
        <Tabs.Panel value="appearance" pt="md">
            <Flex justify="space-between">
                <Text>Enable dark theme</Text>

                <Switch
                    width="100%"
                    labelPosition="left"
                    size="md"
                    checked={isDark}
                    color={theme.primaryColor}
                    onChange={(event) => {
                        const val = event.currentTarget.checked;
                        toggleColorScheme(val ? "dark" : "light");
                        setIsDark(val);
                    }}
                />
            </Flex>

            <Text mt="md">Change primary color</Text>
            <Group spacing="xs" mt="sm">
                {Object.keys(theme.colors)
                    .reverse()
                    .map((color) => (
                        <ColorSwatch
                            radius="xs"
                            size={theme.primaryColor === color ? 36 : 28}
                            key={color}
                            color={theme.colors[color][6]}
                            onClick={() => {
                                setPrimaryColor(color);
                            }}
                            sx={{
                                cursor: "pointer",
                            }}
                        />
                    ))}
            </Group>
        </Tabs.Panel>
    );
};

const SHORTCUTS = [
    { label: "Show bot commands", key1: "Ctrl", key2: "K" },
    { label: "Copy last bot message", key1: "Ctrl", key2: "B" },
    { label: "Focus message textarea", key1: "Ctrl", key2: "M" },
    { label: "Toggle theme", key1: "Ctrl", key2: "J" },
];

const ShortcutsTab = () => {
    return (
        <Tabs.Panel value="shortcuts" pt="md">
            <Flex direction="column" gap="md">
                {SHORTCUTS.map((sc) => (
                    <Flex gap="lg" key={sc.label}>
                        <Box>
                            <Kbd>{sc.key1}</Kbd> + <Kbd>{sc.key2}</Kbd>
                        </Box>
                        <Text>{sc.label}</Text>
                    </Flex>
                ))}
            </Flex>
        </Tabs.Panel>
    );
};
