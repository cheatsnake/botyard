import { Carousel } from "@mantine/carousel";
import { File } from "./types";
import { FC, useState } from "react";
import { Box, Flex, Image, Loader, Paper, Text } from "@mantine/core";
import { openNewTab } from "../../helpers/link";
import { IconFile } from "@tabler/icons-react";
import { truncString } from "../../helpers/text";
import { parseByteSize } from "../../helpers/size";

interface AttachmentGroupProps {
    files: File[];
}

export const ImageGroup: FC<AttachmentGroupProps> = (props) => {
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

export const AudioGroup: FC<AttachmentGroupProps> = (props) => {
    return (
        <Flex wrap="wrap" gap="sm">
            {props.files.map((file) => (
                <audio key={file.id} src={file.path} controls />
            ))}
        </Flex>
    );
};

export const FileGroup: FC<AttachmentGroupProps> = (props) => {
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
