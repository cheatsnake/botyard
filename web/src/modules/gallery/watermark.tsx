import { Box, Text } from "@mantine/core";
import { openNewTab } from "../../helpers/link";

export const Watermark = () => {
    return (
        <Box ta="center">
            <Text
                component="a"
                opacity={0.3}
                onClick={() => openNewTab("https://github.com/cheatsnake/botyard")}
                sx={{ cursor: "pointer", ":hover": { opacity: 0.6 } }}
            >
                Made with Botyard
            </Text>
        </Box>
    );
};
