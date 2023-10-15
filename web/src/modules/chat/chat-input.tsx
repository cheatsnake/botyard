import { Flex, Textarea, Tooltip, ActionIcon, FileButton, LoadingOverlay, useMantineTheme } from "@mantine/core";
import { spotlight } from "@mantine/spotlight";
import { IconSend, IconLink } from "@tabler/icons-react";
import { FC, useRef } from "react";
import { Bot, Chat, Message } from "../../api/types";
import { AttachmentList } from "./attachment-list";
import ClientAPI from "../../api/client-api";
import { useUserContext } from "../../contexts/user-context";
import { errNotify, warnNotify } from "../../helpers/notifications";
import { delay } from "../../helpers/async";
import { CommandPicker } from "./command-picker";
import { TextareaHotkeys } from "./chat-hotkeys";

interface ChatInputProps {
    body: string;
    attachments: File[];
    isBlockInput: boolean;
    currentChat?: Chat;
    currentBot: Bot;
    setBody: (value: React.SetStateAction<string>) => void;
    setMessages: React.Dispatch<React.SetStateAction<Message[]>>;
    setAttachments: React.Dispatch<React.SetStateAction<File[]>>;
    setIsBlockInput: React.Dispatch<React.SetStateAction<boolean>>;
    scrollToBottom: () => void;
}

const START_POOLING_DELAY = 500;
const POOLING_DELAY = 1000;
const MAX_POOLS = 7;

export const ChatInput: FC<ChatInputProps> = (props) => {
    const { body, attachments, currentChat, currentBot, isBlockInput } = props;
    const { setBody, setMessages, setAttachments, setIsBlockInput, scrollToBottom } = props;

    const { colors, colorScheme } = useMantineTheme();
    const { user } = useUserContext();
    const textareaRef = useRef<HTMLTextAreaElement>(null);

    const sendMessage = async (value?: string) => {
        try {
            if (!currentChat || !user?.id) return;
            setIsBlockInput(true);

            const attachmentIds: string[] = [];
            if (attachments.length > 0) {
                const attachs = await ClientAPI.uploadFiles(attachments);
                attachmentIds.push(...attachs.map(({ id }) => id));
            }

            const newMsg = await ClientAPI.sendUserMessage({
                chatId: currentChat.id,
                senderId: user.id,
                body: value ?? body,
                attachmentIds,
            });

            setMessages((prev) => [...prev, newMsg]);
            setTimeout(scrollToBottom, 1);
            setBody("");
            setAttachments([]);

            await botPooling(newMsg.timestamp);
        } catch (error) {
            errNotify((error as Error).message);
        } finally {
            setIsBlockInput(false);
        }
    };

    const commandTrigger = (alias: string) => {
        if (isBlockInput) return;
        sendMessage("/" + alias);
    };

    const botPooling = async (since: number) => {
        try {
            let isFinish = false;
            let poolsCount = 0;

            await delay(START_POOLING_DELAY);

            while (!isFinish && poolsCount < MAX_POOLS) {
                const msgPage = await ClientAPI.getMessagesByChat(
                    currentChat?.id as string,
                    currentBot.id,
                    1,
                    1,
                    since
                );

                if (msgPage.messages.length > 0) {
                    setMessages((prev) => [...prev, msgPage.messages[0]]);
                    setTimeout(scrollToBottom, 1);
                    return;
                }

                poolsCount++;

                if (poolsCount === MAX_POOLS) {
                    warnNotify("Perhaps the bot is currently offline. Try again later.", "Bot is not responding");
                    return;
                }

                await delay(POOLING_DELAY);
            }
        } catch (error) {
            throw error;
        }
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

        if (event.key === "Escape") {
            event.preventDefault();
            event.currentTarget.blur();
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
        <CommandPicker commandTrigger={commandTrigger}>
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
                    ref={textareaRef}
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
            <TextareaHotkeys textarea={textareaRef.current} />
        </CommandPicker>
    );
};
