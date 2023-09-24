import { Space, Text } from "@mantine/core";
import { openNewTab } from "../../helpers/link";

export const Watermark = () => {
    return (
        <>
            <Text
                ta="center"
                opacity={0.3}
                onClick={() => openNewTab("https://github.com/cheatsnake/botyard")}
                sx={{ cursor: "pointer" }}
            >
                Made with Botyard
            </Text>
            <Space h="xl" />
        </>
    );
};
