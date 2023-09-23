import { Flex, Textarea, Tooltip, ActionIcon, FileButton } from "@mantine/core";
import { SpotlightProvider, spotlight } from "@mantine/spotlight";
import { IconSend, IconLink } from "@tabler/icons-react";
import { FC } from "react";
import { Command } from "./types";

interface ChatInputProps {
    commands: Command[];
    body: string;
    setBody: (value: React.SetStateAction<string>) => void;
    sendMessage: (value?: string) => void;
}

export const ChatInput: FC<ChatInputProps> = ({ commands, body, setBody, sendMessage }) => {
    const commandTrigger = (alias: string) => {
        sendMessage("/" + alias);
    };

    const handleInput = (event: React.ChangeEvent<HTMLTextAreaElement>) => {
        if (event.currentTarget.value === "/") {
            spotlight.open();
        }
        setBody(event.currentTarget.value);
    };

    const handleKeyPress = (event: React.KeyboardEvent<HTMLTextAreaElement>) => {
        if (event.key === "Enter" && !event.shiftKey) {
            event.preventDefault();
            if (body) sendMessage();
        }

        if (event.key === "Tab") {
            event.preventDefault();

            const { selectionStart, selectionEnd, value } = event.currentTarget;
            const newValue = value.substring(0, selectionStart) + "\t" + value.substring(selectionEnd);

            setBody(newValue);
            event.currentTarget.setSelectionRange(selectionStart + 1, selectionStart + 1);
        }
    };

    return (
        <SpotlightProvider
            centered
            actions={commands.map((cmd) => ({
                title: cmd.alias,
                description: cmd.description,
                onTrigger: () => commandTrigger(cmd.alias),
            }))}
            searchPlaceholder="Command..."
            overlayProps={{ blur: "none" }}
            limit={100}
        >
            <Flex align="center" gap="sm" p="sm" pt="lg">
                <Textarea
                    value={body}
                    onChange={handleInput}
                    onKeyDown={handleKeyPress}
                    w="100%"
                    size={window.innerWidth > 600 ? "lg" : "sm"}
                    placeholder="Your message..."
                    minRows={2}
                    maxRows={6}
                    autosize
                    // disabled
                />
                <Flex direction="column" gap="5px" h="100%">
                    <Tooltip openDelay={700} withArrow label="Send message">
                        <ActionIcon disabled={body.length === 0} size="lg" h="50%" onClick={() => sendMessage()}>
                            <IconSend />
                        </ActionIcon>
                    </Tooltip>

                    <Tooltip openDelay={700} label="Attach files">
                        <FileButton onChange={() => {}} multiple>
                            {(props) => (
                                <Tooltip openDelay={700} withArrow label="Attach files">
                                    <ActionIcon {...props} c="gray" size="lg" h="50%">
                                        <IconLink />
                                    </ActionIcon>
                                </Tooltip>
                            )}
                        </FileButton>
                    </Tooltip>
                </Flex>
            </Flex>
        </SpotlightProvider>
    );
};
