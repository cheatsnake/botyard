import { Flex, Avatar, Title, Badge, Text } from "@mantine/core";
import { openNewTab } from "../../helpers/link";
import { abbreviateName } from "../../helpers/text";
import { FC } from "react";

const CONFIG = {
    serviceName: "Botyard showcase",
    description: `Lorem ipsum, dolor sit amet consectetur adipisicing elit. Provident, sequi molestias dolorem
    voluptatibus aut, ullam tempora distinctio, a enim ratione quidem accusamus fugit cupiditate.
    Quisquam deserunt natus delectus soluta iste? Lorem ipsum dolor sit amet consectetur adipisicing
    elit. Beatae quod hic cumque.`,
    avatar: "",
    socials: [
        { title: "Website", url: "https://example.com" },
        { title: "GitHub", url: "https://example.com" },
        { title: "Discord", url: "https://example.com" },
        { title: "Twitter", url: "https://example.com" },
    ],
};

export const OrganizationInfo: FC = () => {
    return (
        <Flex gap="1rem" direction="column" sx={{ padding: "1rem 0" }}>
            <Flex justify="center" direction="column" align="center">
                <Avatar size="xl" radius="xl" src={CONFIG.avatar}>
                    {CONFIG.avatar ? null : abbreviateName(CONFIG.serviceName)}
                </Avatar>

                <Title mt="sm" order={1} size="h2">
                    {CONFIG.serviceName}
                </Title>
            </Flex>

            <Flex wrap="wrap" justify="center" gap="xs">
                {CONFIG.socials.map((link) => (
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

            <Text align="justify">{CONFIG.description}</Text>
        </Flex>
    );
};
