import { Grid, Avatar, Box, Title, Text } from "@mantine/core";
import { abbreviateName } from "../../helpers/text";
import { useNavigate } from "react-router-dom";

const BOTS = [
    { id: "1", name: "Echo Bot", avatar: "", description: "Lorem ipsum, dolor sit amet consectetur adipisicing elit." },
    {
        id: "2",
        name: "Bot Caclulator",
        avatar: "",
        description: "Lorem ipsum, dolor sit amet consectetur adipisicing elit.",
    },
    {
        id: "3",
        name: "Bot Googler",
        avatar: "",
        description: "Lorem ipsum, dolor sit amet consectetur adipisicing elit.",
    },
    {
        id: "4",
        name: "Echo Bot 2",
        avatar: "",
        description: "Lorem ipsum, dolor sit amet consectetur adipisicing elit.",
    },
    {
        id: "5",
        name: "Bot Caclulator 2",
        avatar: "",
        description: "Lorem ipsum, dolor sit amet consectetur adipisicing elit.",
    },
    {
        id: "6",
        name: "Bot Googler 2",
        avatar: "",
        description: `Lorem ipsum, dolor sit amet consectetur adipisicing elit. Provident, sequi molestias dolorem
    voluptatibus aut, ullam tempora distinctio, a enim ratione quidem accusamus fugit cupiditate.`,
    },
    {
        id: "7",
        name: "Echo Bot 3",
        avatar: "",
        description: "Lorem ipsum, dolor sit amet consectetur adipisicing elit.",
    },
    {
        id: "8",
        name: "Bot Caclulator 3",
        avatar: "",
        description: "Lorem ipsum, dolor sit amet consectetur adipisicing elit.",
    },
];

export const BotList = () => {
    const navigate = useNavigate();

    return (
        <Grid grow gutter="sm" justify="center" sx={{ gap: "1rem" }}>
            {BOTS.map((bot) => (
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
                        backgroundColor: theme.colorScheme === "dark" ? theme.colors.dark[6] : theme.colors.gray[0],
                        gap: "0.75rem",
                        padding: theme.spacing.xl,
                        borderRadius: theme.radius.md,
                        cursor: "pointer",

                        "&:hover": {
                            backgroundColor: theme.colorScheme === "dark" ? theme.colors.dark[5] : theme.colors.gray[1],
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
    );
};
