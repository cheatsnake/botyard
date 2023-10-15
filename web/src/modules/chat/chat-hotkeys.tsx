import { FC } from "react";
import { Message } from "../../api/types";
import { useClipboard, useHotkeys } from "@mantine/hooks";
import { successNotify, warnNotify } from "../../helpers/notifications";

export const BotHotkeys: FC<{ messages: Message[]; botId?: string }> = (props) => {
    const clipboard = useClipboard();

    const copyLastBotMessage = () => {
        let lastMsgBody = "";

        for (let i = props.messages.length - 1; i !== 0; i--) {
            if (props.messages[i].senderId !== props.botId) continue;
            lastMsgBody = props.messages[i].body;
            break;
        }

        if (lastMsgBody.length === 0) {
            warnNotify("Last bot message not found.", "Nothing to copy");
            return;
        }

        clipboard.copy(lastMsgBody);

        if (clipboard.error) {
            warnNotify("Failed to copy the bot's last message. You may be using an outdated browser.", "Copy failed");
            return;
        }

        successNotify("The bot's last response is copied to the clipboard.");
    };

    useHotkeys([["mod+B", copyLastBotMessage]]);

    return <></>;
};

export const TextareaHotkeys: FC<{ textarea: HTMLTextAreaElement | null }> = (props) => {
    const focus = () => {
        props.textarea?.focus();
    };

    useHotkeys([["mod+M", focus]]);
    return <></>;
};
