import { Carousel } from "@mantine/carousel";
import { File } from "./types";
import { FC } from "react";
import { Flex, Image } from "@mantine/core";
import { openNewTab } from "../../helpers/link";

export const ImageGroup: FC<{ files: File[] }> = (props) => {
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
                    <Image src={file.path} height={300} onClick={() => openNewTab(file.path)} radius="sm" />
                </Carousel.Slide>
            ))}
        </Carousel>
    );
};

export const AudioGroup: FC<{ files: File[] }> = (props) => {
    return (
        <Flex wrap="wrap" gap="sm">
            {props.files.map((file) => (
                <audio key={file.id} src={file.path} controls />
            ))}
        </Flex>
    );
};
