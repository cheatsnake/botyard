import { Avatar, Badge, Box, Container, Divider, Flex, Grid, Space, Text, Title } from "@mantine/core";
import { abbreviateName } from "../../helpers/test.helpers";
import { Header } from "./header";
import { useNavigate } from "react-router-dom";

const SOCIAL_LINKS = [
    { title: "Website", url: "https://example.com" },
    { title: "GitHub", url: "https://example.com" },
    { title: "Discord", url: "https://example.com" },
    { title: "Twitter", url: "https://example.com" },
];

const BOTS = [
    { name: "Echo Bot", avatar: "", description: "Lorem ipsum, dolor sit amet consectetur adipisicing elit." },
    { name: "Bot Caclulator", avatar: "", description: "Lorem ipsum, dolor sit amet consectetur adipisicing elit." },
    { name: "Bot Googler", avatar: "", description: "Lorem ipsum, dolor sit amet consectetur adipisicing elit." },
    { name: "Echo Bot 2", avatar: "", description: "Lorem ipsum, dolor sit amet consectetur adipisicing elit." },
    { name: "Bot Caclulator 2", avatar: "", description: "Lorem ipsum, dolor sit amet consectetur adipisicing elit." },
    {
        name: "Bot Googler 2",
        avatar: "",
        description: `Lorem ipsum, dolor sit amet consectetur adipisicing elit. Provident, sequi molestias dolorem
    voluptatibus aut, ullam tempora distinctio, a enim ratione quidem accusamus fugit cupiditate.`,
    },
    { name: "Echo Bot 3", avatar: "", description: "Lorem ipsum, dolor sit amet consectetur adipisicing elit." },
    { name: "Bot Caclulator 3", avatar: "", description: "Lorem ipsum, dolor sit amet consectetur adipisicing elit." },
    { name: "Testing", avatar: "", description: "Lorem ipsum, dolor sit amet consectetur adipisicing elit." },
];

export const GalleryPage = () => {
    const navigate = useNavigate();

    const openNewTab = (url: string) => {
        window.open(url, "_blank")?.focus();
    };

    return (
        <>
            <Header />
            <Container size="md">
                <Flex gap="1rem" direction="column" sx={{ padding: "1rem 0" }}>
                    <Flex justify="center" direction="column" align="center">
                        <Avatar size="xl" radius="xl" />
                        <Title mt="sm" order={1} size="h2">
                            Botyard showcase
                        </Title>
                    </Flex>
                    <Flex wrap="wrap" justify="center" gap="xs">
                        {SOCIAL_LINKS.map((link) => (
                            <Badge
                                key={link.title}
                                component="button"
                                radius="xs"
                                size="lg"
                                onClick={() => openNewTab(link.url)}
                                sx={{ cursor: "pointer" }}
                            >
                                {link.title}
                            </Badge>
                        ))}
                    </Flex>
                    <Text align="justify">
                        Lorem ipsum, dolor sit amet consectetur adipisicing elit. Provident, sequi molestias dolorem
                        voluptatibus aut, ullam tempora distinctio, a enim ratione quidem accusamus fugit cupiditate.
                        Quisquam deserunt natus delectus soluta iste? Lorem ipsum dolor sit amet consectetur adipisicing
                        elit. Beatae quod hic cumque.
                    </Text>
                </Flex>
                <Divider p="md" />
                <Box>
                    <Grid grow gutter="sm" justify="center" sx={{ gap: "1rem" }}>
                        {BOTS.map((bot) => (
                            <Grid.Col
                                p={window.innerWidth > 730 ? "lg" : "sm"}
                                key={bot.name}
                                span={4}
                                miw={300}
                                display="flex"
                                onClick={() => {
                                    if ("startViewTransition" in document) {
                                        //@ts-ignore
                                        return document.startViewTransition(() => {
                                            navigate("/bot/1");
                                        });
                                    }

                                    return navigate("/bot/1");
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
                </Box>

                <Space h="xl" />

                <Text my="xl" ta="center" color="gray">
                    Made with Botyard
                </Text>
            </Container>
        </>
    );
};
