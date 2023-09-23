import { Flex, Textarea, Tooltip, ActionIcon, FileButton } from "@mantine/core";
import { IconSend, IconLink } from "@tabler/icons-react";
import { FC } from "react";

interface ChatInputProps {
    body: string;
    setBody: (value: React.SetStateAction<string>) => void;
    sendMessage: () => void;
}

export const ChatInput: FC<ChatInputProps> = ({ body, setBody, sendMessage }) => {
    return (
        <Flex align="center" gap="sm" p="sm" pt="lg">
            <Textarea
                value={body}
                onChange={(event) => setBody(event.currentTarget.value)}
                onKeyDown={(event) => {
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
                }}
                w="100%"
                size={window.innerWidth > 600 ? "md" : "sm"}
                placeholder="Your message..."
                minRows={2}
                maxRows={6}
                autosize
            />
            <Flex direction="column" gap="5px" h="100%">
                <Tooltip openDelay={700} withArrow label="Send message">
                    <ActionIcon disabled={body.length === 0} size="lg" h="50%" onClick={sendMessage}>
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
    );
};
