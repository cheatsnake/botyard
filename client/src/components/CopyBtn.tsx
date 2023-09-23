import { CopyButton, ActionIcon, Tooltip } from "@mantine/core";
import { IconCopy, IconCheck } from "@tabler/icons-react";
import { FC } from "react";

interface CopyBtnProps {
    value: string;
    thing?: string;
}

export const CopyBtn: FC<CopyBtnProps> = (props) => {
    const beforeLabel = props.thing ? "Copy " + props.thing.toLowerCase() : "Copy";
    const afterLabel = props.thing ? props.thing + " copied" : "Copied";

    return (
        <CopyButton value={props.value} timeout={2000}>
            {({ copied, copy }) => (
                <Tooltip label={copied ? afterLabel : beforeLabel} withArrow position="left">
                    <ActionIcon color={copied ? "primary" : "gray"} onClick={copy}>
                        {copied ? <IconCheck size="1rem" /> : <IconCopy size="1rem" />}
                    </ActionIcon>
                </Tooltip>
            )}
        </CopyButton>
    );
};
