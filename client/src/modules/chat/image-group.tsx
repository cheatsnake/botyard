import { Carousel } from "@mantine/carousel";
import { File } from "./types";
import { FC } from "react";
import { Image } from "@mantine/core";
import { KNOWN_MIME_TYPES } from "./const";
import { openNewTab } from "../../helpers/link";

export const ImageGroup: FC<{ files: File[] }> = (props) => {
    return (
        <Carousel
            loop
            slideGap="sm"
            align="start"
            controlsOffset="xs"
            controlSize={20}
            withIndicators
            slideSize="50%"
            breakpoints={[
                { maxWidth: "md", slideSize: "50%" },
                { maxWidth: "sm", slideSize: "100%", slideGap: 0 },
            ]}
        >
            {props.files
                .filter((file) => KNOWN_MIME_TYPES[file.mimeType] === "image")
                .map((file) => (
                    <Carousel.Slide key={file.id}>
                        <Image src={file.path} height={300} onClick={() => openNewTab(file.path)} />
                    </Carousel.Slide>
                ))}
        </Carousel>
    );
};
