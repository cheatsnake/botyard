import { Box, CloseButton, Flex, ScrollArea, Text } from "@mantine/core";
import { IconFile, IconMusic, IconPhoto, IconVideo } from "@tabler/icons-react";
import { FC } from "react";
import { truncString } from "../../helpers/text";

type FileType = "image" | "video" | "audio";

const KNOWN_FILE_TYPES: { [key: string]: FileType } = {
    "image/gif": "image",
    "image/jpeg": "image",
    "image/png": "image",
    "image/svg+xml": "image",
    "image/webp": "image",

    "video/mp4": "video",
    "video/webm": "video",
    "video/ogg": "video",
    "video/quicktime": "video",
    "video/x-flv": "video",

    "audio/mpeg": "audio",
    "audio/ogg": "audio",
    "audio/wav": "audio",
    "audio/aac": "audio",
};

interface AttachmentListProps {
    files: File[];
    remover: (i: number) => void;
}

export const AttachmentList: FC<AttachmentListProps> = (props) => {
    return (
        <Box>
            <ScrollArea>
                <Flex gap="sm" px="sm" pt="md">
                    {props.files.map((file, idx) => (
                        <Flex
                            key={file.name}
                            pos="relative"
                            direction="column"
                            align="center"
                            gap="sm"
                            bg="dark"
                            p="sm"
                            miw={96}
                            sx={{ borderRadius: "0.3rem" }}
                        >
                            <CloseButton
                                pos="absolute"
                                opacity={0.5}
                                right={1}
                                top={1}
                                onClick={() => props.remover(idx)}
                            />
                            {defineFileIcon(file.type)}
                            <Text size="sm" sx={{ whiteSpace: "nowrap" }}>
                                {truncString(file.name, 10)}
                            </Text>
                        </Flex>
                    ))}
                </Flex>
            </ScrollArea>
        </Box>
    );
};

const defineFileIcon = (type: string) => {
    if (KNOWN_FILE_TYPES[type] === "image") return <IconPhoto size={32} />;
    if (KNOWN_FILE_TYPES[type] === "video") return <IconVideo size={32} />;
    if (KNOWN_FILE_TYPES[type] === "audio") return <IconMusic size={32} />;

    return <IconFile size={32} />;
};
