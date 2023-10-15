import { Grid, Avatar, Box, Title, Text, useMantineTheme } from "@mantine/core";
import { abbreviateName, truncString } from "../../helpers/text";
import { useNavigate } from "react-router-dom";
import { useEffect } from "react";
import { useLoaderContext } from "../../contexts/loader-context";
import { errNotify } from "../../helpers/notifications";
import { useStorageContext } from "../../contexts/storage-context";

export const BotList = () => {
    const { primaryColor } = useMantineTheme();
    const { bots, loadBots } = useStorageContext();
    const { setIsLoad } = useLoaderContext();
    const navigate = useNavigate();

    useEffect(() => {
        (async () => {
            try {
                setIsLoad(true);
                await loadBots();
            } catch (error) {
                errNotify((error as Error).message);
            } finally {
                setIsLoad(false);
            }
        })();
    }, []);

    return (
        <>
            {bots.length > 0 ? (
                <Grid grow gutter="sm" justify="center" py="sm" sx={{ gap: "1rem" }}>
                    {bots.map((bot) => (
                        <Grid.Col
                            p={window.innerWidth > 730 ? "lg" : "sm"}
                            key={bot.name}
                            span={4}
                            miw={300}
                            display="flex"
                            onClick={() => {
                                const link = "/bot/" + bot.id;
                                if ("startViewTransition" in document) {
                                    //@ts-ignore
                                    return document.startViewTransition(() => {
                                        navigate(link);
                                    });
                                }

                                return navigate(link);
                            }}
                            sx={(theme) => ({
                                backgroundColor:
                                    theme.colorScheme === "dark" ? theme.colors.dark[6] : theme.colors.gray[0],
                                gap: "0.75rem",
                                padding: theme.spacing.xl,
                                borderRadius: theme.radius.md,
                                cursor: "pointer",

                                "&:hover": {
                                    backgroundColor:
                                        theme.colorScheme === "dark" ? theme.colors.dark[5] : theme.colors.gray[1],
                                },
                            })}
                        >
                            <Avatar color={primaryColor} size="xl" src={bot.avatar}>
                                {!bot.avatar ? abbreviateName(bot.name) : ""}
                            </Avatar>

                            <Box>
                                <Title order={3} size="h4">
                                    {bot.name}
                                </Title>

                                <Text ta="justify" opacity={0.7}>
                                    {truncString(bot.description || "", 100)}
                                </Text>
                            </Box>
                        </Grid.Col>
                    ))}
                </Grid>
            ) : null}
        </>
    );
};
