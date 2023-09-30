import { Grid, Avatar, Box, Title, Text } from "@mantine/core";
import { abbreviateName } from "../../helpers/text";
import { useNavigate } from "react-router-dom";
import { useEffect, useState } from "react";
import { Bot } from "../../api/types";
import { useLoaderContext } from "../../contexts/loader-context";
import ClientAPI from "../../api/client-api";
import { errNotify } from "../../helpers/notifications";

export const BotList = () => {
    const [bots, setBots] = useState<Bot[]>([]);
    const { setIsLoad } = useLoaderContext();
    const navigate = useNavigate();

    useEffect(() => {
        (async () => {
            try {
                setIsLoad(true);
                const allBots = await ClientAPI.getAllBots();
                setBots(allBots);
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
                            <Avatar color="primary" size="xl" src={bot.avatar}>
                                {!bot.avatar ? abbreviateName(bot.name) : ""}
                            </Avatar>

                            <Box>
                                <Title order={3} size="h4">
                                    {bot.name}
                                </Title>

                                <Text ta="justify" opacity={0.7}>
                                    {bot.description}
                                </Text>
                            </Box>
                        </Grid.Col>
                    ))}
                </Grid>
            ) : null}
        </>
    );
};
