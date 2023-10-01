import { Flex, Textarea, Tooltip, ActionIcon, FileButton, LoadingOverlay, useMantineTheme } from "@mantine/core";
import { SpotlightProvider, spotlight } from "@mantine/spotlight";
import { IconSend, IconLink } from "@tabler/icons-react";
import { FC } from "react";
import { BotCommand } from "../../api/types";
import { AttachmentList } from "./attachment-list";

interface ChatInputProps {
    body: string;
    attachments: File[];
    isBlockInput: boolean;
    setAttachments: React.Dispatch<React.SetStateAction<File[]>>;
    setBody: (value: React.SetStateAction<string>) => void;
    sendMessage: (value?: string) => Promise<void>;
}

const COMMANDS: BotCommand[] = [
    // { alias: "start", description: "Init a new bot conversation." },
    // { alias: "help", description: "Print some instructions." },
    // { alias: "ping", description: "Send pong message." },
];

export const ChatInput: FC<ChatInputProps> = (props) => {
    const { body, attachments, isBlockInput, setAttachments, setBody, sendMessage } = props;
    const { colors, colorScheme } = useMantineTheme();

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
            if ((body.length > 0 || attachments.length > 0) && !isBlockInput) sendMessage();
        }

        if (event.key === "Tab") {
            event.preventDefault();

            const { selectionStart, selectionEnd, value } = event.currentTarget;
            const newValue = value.substring(0, selectionStart) + "\t" + value.substring(selectionEnd);

            setBody(newValue);
            event.currentTarget.setSelectionRange(selectionStart + 1, selectionStart + 1);
        }
    };

    const addAttachments = (files: File[]) => {
        const attached = attachments.map(({ name }) => name);
        setAttachments([...attachments, ...files.filter(({ name }) => !attached.includes(name))]);
    };

    const removeAttachment = (index: number) => {
        setAttachments([...attachments.slice(0, index), ...attachments.slice(index + 1)]);
    };

    return (
        <SpotlightProvider
            centered
            actions={COMMANDS.map((cmd) => ({
                title: cmd.alias,
                description: cmd.description,
                onTrigger: () => commandTrigger(cmd.alias),
            }))}
            nothingFoundMessage="Command not found..."
            searchPlaceholder="Command..."
            overlayProps={{ blur: "none" }}
            limit={100}
        >
            {attachments.length > 0 ? <AttachmentList files={attachments} remover={removeAttachment} /> : null}

            <Flex align="center" gap="sm" p="sm" pt="lg" pos="relative">
                <LoadingOverlay
                    visible={isBlockInput}
                    overlayColor={colorScheme === "dark" ? colors.dark[7] : colors.gray[0]}
                    loaderProps={{ variant: "dots" }}
                />
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
                />
                <Flex direction="column" gap="5px" h="100%">
                    <Tooltip openDelay={700} withArrow label="Send message">
                        <ActionIcon
                            disabled={body.length === 0 && attachments.length === 0}
                            size="lg"
                            h="50%"
                            onClick={() => sendMessage()}
                        >
                            <IconSend />
                        </ActionIcon>
                    </Tooltip>

                    <Tooltip openDelay={700} label="Attach files">
                        <FileButton onChange={addAttachments} multiple>
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
