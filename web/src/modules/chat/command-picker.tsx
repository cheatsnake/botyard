import { SpotlightProvider, spotlight } from "@mantine/spotlight";
import { FC, ReactNode, useEffect, useState } from "react";
import { BotCommand } from "../../api/types";
import { errNotify, warnNotify } from "../../helpers/notifications";
import ClientAPI from "../../api/client-api";
import { useParams } from "react-router-dom";

interface CommandPickerProps {
    children: ReactNode;
    commandTrigger: (alias: string) => void;
}

export const CommandPicker: FC<CommandPickerProps> = (props) => {
    const [commands, setCommands] = useState<BotCommand[]>([]);
    const [showEmptyNotifier, setShowEmptyNotifier] = useState(true);
    const { id: botId } = useParams();

    useEffect(() => {
        (async () => {
            try {
                const cmds = await ClientAPI.getBotCommands(botId || "");
                setCommands(cmds);
            } catch (error) {
                errNotify((error as Error).message);
            }
        })();
    }, []);

    return (
        <SpotlightProvider
            centered
            actions={commands.map((cmd) => ({
                title: cmd.alias,
                description: cmd.description,
                onTrigger: () => props.commandTrigger(cmd.alias),
            }))}
            onSpotlightOpen={() => {
                if (commands.length === 0) {
                    setTimeout(() => spotlight.close(), 1);
                    if (!showEmptyNotifier) return;
                    warnNotify("The current bot has no commands found.", "No commands");
                    setShowEmptyNotifier(false);
                }
            }}
            nothingFoundMessage="Command not found..."
            searchPlaceholder="Command..."
            overlayProps={{ blur: "none" }}
            limit={100}
        >
            {props.children}
        </SpotlightProvider>
    );
};
