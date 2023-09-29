import { Carousel } from "@mantine/carousel";
import { File } from "./types";
import { FC, useState } from "react";
import { Box, Flex, Image, Loader, Paper, Space, Text } from "@mantine/core";
import { openNewTab } from "../../helpers/link";
import { IconFile } from "@tabler/icons-react";
import { truncString } from "../../helpers/text";
import { parseByteSize } from "../../helpers/size";
import { KNOWN_MIME_TYPES } from "./const";

interface AttachmentGroupProps {
    files: File[];
}

export const AttachmentGroups: FC<AttachmentGroupProps> = (props) => {
    const images = props.files.filter((file) => KNOWN_MIME_TYPES[file.mimeType] === "image");
    const videos = props.files.filter((file) => KNOWN_MIME_TYPES[file.mimeType] === "video");
    const audios = props.files.filter((file) => KNOWN_MIME_TYPES[file.mimeType] === "audio");
    const files = props.files.filter((file) => !KNOWN_MIME_TYPES[file.mimeType]);

    return (
        <>
            {images.length > 0 ? (
                <>
                    <Space h="xl" />
                    <ImageGroup files={props.files.filter((file) => KNOWN_MIME_TYPES[file.mimeType] === "image")} />
                </>
            ) : null}

            {videos.length > 0 ? (
                <>
                    <Space h="xl" />
                    <VideoGroup files={props.files.filter((file) => KNOWN_MIME_TYPES[file.mimeType] === "video")} />
                </>
            ) : null}

            {audios.length > 0 ? (
                <>
                    <Space h="xl" />
                    <AudioGroup files={props.files.filter((file) => KNOWN_MIME_TYPES[file.mimeType] === "audio")} />
                </>
            ) : null}

            {files.length > 0 ? (
                <>
                    <Space h="xl" />
                    <FileGroup files={props.files.filter((file) => !KNOWN_MIME_TYPES[file.mimeType])} />
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
            withControls={props.files.length > 1}
            withIndicators={props.files.length > 1}
            slideSize={props.files.length > 1 ? "50%" : "auto"}
            breakpoints={[
                { maxWidth: "md", slideSize: "50%" },
                { maxWidth: "sm", slideSize: "100%", slideGap: 0 },
            ]}
        >
            {props.files.map((file) => (
                <Carousel.Slide key={file.id}>
                    <ImageWithPlaceholder path={file.path} />
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
            {props.files.map((file) => (
                <video width={"100%"} controls src={file.path} style={{ borderRadius: "0.3rem" }} />
            ))}
        </>
    );
};

const AudioGroup: FC<AttachmentGroupProps> = (props) => {
    return (
        <Flex wrap="wrap" gap="sm">
            {props.files.map((file) => (
                <audio key={file.id} src={file.path} controls />
            ))}
        </Flex>
    );
};

const FileGroup: FC<AttachmentGroupProps> = (props) => {
    return (
        <Flex wrap="wrap" gap="sm">
            {props.files.map((file) => (
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
                        <Text
                            component="a"
                            href={file.path}
                            sx={{ cursor: "pointer", ":hover": { textDecoration: "underline" } }}
                        >
                            {truncString(file.name, 20)}
                        </Text>
                        <Text size="xs" opacity={0.7}>
                            {parseByteSize(file.size)}
                        </Text>
                    </Box>
                </Paper>
            ))}
        </Flex>
    );
};
