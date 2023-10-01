import { Carousel } from "@mantine/carousel";
import { Attachment } from "../../api/types";
import { FC, useState } from "react";
import { Box, Flex, Image, Loader, Paper, Space, Text, Tooltip } from "@mantine/core";
import { openNewTab } from "../../helpers/link";
import { IconFile } from "@tabler/icons-react";
import { truncString } from "../../helpers/text";
import { parseByteSize } from "../../helpers/size";
import { KNOWN_MIME_TYPES } from "./const";

interface AttachmentGroupProps {
    attachments: Attachment[];
}

const getFullFilePath = (path: string) => `/${path}`;

export const AttachmentGroups: FC<AttachmentGroupProps> = (props) => {
    const images = props.attachments.filter((file) => KNOWN_MIME_TYPES[file.mimeType] === "image");
    const videos = props.attachments.filter((file) => KNOWN_MIME_TYPES[file.mimeType] === "video");
    const audios = props.attachments.filter((file) => KNOWN_MIME_TYPES[file.mimeType] === "audio");
    const files = props.attachments.filter((file) => !KNOWN_MIME_TYPES[file.mimeType]);

    return (
        <>
            {images.length > 0 ? (
                <>
                    <Space h="md" />
                    <ImageGroup
                        attachments={props.attachments.filter((file) => KNOWN_MIME_TYPES[file.mimeType] === "image")}
                    />
                </>
            ) : null}

            {videos.length > 0 ? (
                <>
                    <Space h="md" />
                    <VideoGroup
                        attachments={props.attachments.filter((file) => KNOWN_MIME_TYPES[file.mimeType] === "video")}
                    />
                </>
            ) : null}

            {audios.length > 0 ? (
                <>
                    <Space h="md" />
                    <AudioGroup
                        attachments={props.attachments.filter((file) => KNOWN_MIME_TYPES[file.mimeType] === "audio")}
                    />
                </>
            ) : null}

            {files.length > 0 ? (
                <>
                    <Space h="md" />
                    <FileGroup attachments={props.attachments.filter((file) => !KNOWN_MIME_TYPES[file.mimeType])} />
                </>
            ) : null}
        </>
    );
};

const ImageGroup: FC<AttachmentGroupProps> = (props) => {
    return (
        <Carousel
            loop
            slideGap="sm"
            align="start"
            controlsOffset="xs"
            controlSize={20}
            withControls={props.attachments.length > 1}
            withIndicators={props.attachments.length > 1}
            slideSize={props.attachments.length > 1 ? "50%" : "auto"}
            breakpoints={[
                { maxWidth: "md", slideSize: "50%" },
                { maxWidth: "sm", slideSize: "100%", slideGap: 0 },
            ]}
        >
            {props.attachments.map((img) => (
                <Carousel.Slide key={img.id}>
                    <ImageWithPlaceholder path={getFullFilePath(img.path)} />
                </Carousel.Slide>
            ))}
        </Carousel>
    );
};

const ImageWithPlaceholder: FC<{ path: string }> = ({ path }) => {
    const [isLoad, setIsLoad] = useState(true);

    return (
        <>
            <Image
                display={isLoad ? "none" : "block"}
                onLoad={() => setIsLoad(false)}
                src={path}
                height={300}
                onClick={() => openNewTab(path)}
                radius="sm"
            />
            {isLoad ? (
                <Flex justify="center" align="center" h={300}>
                    <Loader color="gray" />
                </Flex>
            ) : null}
        </>
    );
};

const VideoGroup: FC<AttachmentGroupProps> = (props) => {
    return (
        <>
            {props.attachments.map((video) => (
                <video width={"100%"} controls src={getFullFilePath(video.path)} style={{ borderRadius: "0.3rem" }} />
            ))}
        </>
    );
};

const AudioGroup: FC<AttachmentGroupProps> = (props) => {
    return (
        <Flex wrap="wrap" gap="sm">
            {props.attachments.map((audio) => (
                <audio key={audio.id} src={getFullFilePath(audio.path)} controls />
            ))}
        </Flex>
    );
};

const FileGroup: FC<AttachmentGroupProps> = (props) => {
    return (
        <Flex wrap="wrap" gap="sm">
            {props.attachments.map((file) => (
                <Paper
                    key={file.id}
                    p="sm"
                    withBorder
                    sx={{
                        display: "flex",
                        alignItems: "center",
                        gap: "0.3rem",
                        borderRadius: "0.3rem",
                        backgroundColor: "transparent",
                    }}
                >
                    <IconFile color="gray" size={40} />
                    <Box>
                        <Tooltip label={file.name}>
                            <Text
                                component="a"
                                href={getFullFilePath(file.path)}
                                sx={{ cursor: "pointer", ":hover": { textDecoration: "underline" } }}
                            >
                                {truncString(file.name, 20)}
                            </Text>
                        </Tooltip>
                        <Text size="xs" opacity={0.7}>
                            {parseByteSize(file.size)}
                        </Text>
                    </Box>
                </Paper>
            ))}
        </Flex>
    );
};
