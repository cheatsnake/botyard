import { Modal, ActionIcon, useMantineColorScheme, Text, Divider, Flex, Select, ColorScheme } from "@mantine/core";
import { useDisclosure } from "@mantine/hooks";
import { IconSettings } from "@tabler/icons-react";

export const SettingsModal = () => {
    const [opened, { open, close }] = useDisclosure(false);
    const { colorScheme, toggleColorScheme } = useMantineColorScheme();

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
                <Divider />
                <Modal.Body mih={"25vh"}>
                    <Flex py="lg" align="center" justify="space-between">
                        <Text>Global theme</Text>
                        <Select
                            maw={120}
                            placeholder="Pick one"
                            defaultValue={colorScheme}
                            data={[
                                { value: "dark", label: "Dark" },
                                { value: "light", label: "Light" },
                            ]}
                            dropdownPosition="bottom"
                            onChange={(event) => {
                                toggleColorScheme(event as ColorScheme);
                            }}
                        />
                    </Flex>
                </Modal.Body>
            </Modal>

            <ActionIcon variant="subtle" size="md" onClick={open}>
                <IconSettings />
            </ActionIcon>
        </>
    );
};
