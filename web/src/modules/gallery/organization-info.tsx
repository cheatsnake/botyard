import { Flex, Avatar, Title, Badge, Text, Divider, useMantineTheme } from "@mantine/core";
import { openNewTab } from "../../helpers/link";
import { abbreviateName } from "../../helpers/text";
import { FC, useEffect } from "react";
import ClientAPI from "../../api/client-api";
import { useLoaderContext } from "../../contexts/loader-context";
import { errNotify } from "../../helpers/notifications";
import { useStorageContext } from "../../contexts/storage-context";

export const OrganizationInfo: FC = () => {
    const { primaryColor } = useMantineTheme();
    const { serviceInfo, setServiceInfo } = useStorageContext();
    const { setIsLoad } = useLoaderContext();

    useEffect(() => {
        (async () => {
            try {
                if (serviceInfo) return;
                setIsLoad(true);
                const info = await ClientAPI.getServiceInfo();
                setServiceInfo(info);
            } catch (error) {
                errNotify((error as Error).message);
            } finally {
                setIsLoad(false);
            }
        })();
    }, []);

    return (
        <>
            {serviceInfo ? (
                <Flex gap="md" direction="column">
                    <Flex justify="center" direction="column" align="center" gap="md">
                        <Avatar size="xl" radius="xl" src={serviceInfo.avatar}>
                            {serviceInfo.avatar ? null : abbreviateName(serviceInfo.name)}
                        </Avatar>

                        <Title order={1} size="h2">
                            {serviceInfo.name}
                        </Title>
                    </Flex>

                    <Flex wrap="wrap" justify="center" gap="xs">
                        {serviceInfo.socials.map((link) => (
                            <Badge
                                key={link.title}
                                component="button"
                                color={primaryColor}
                                radius="xs"
                                size="lg"
                                onClick={() => openNewTab(link.url)}
                                sx={{ cursor: "pointer" }}
                            >
                                {link.title}
                            </Badge>
                        ))}
                    </Flex>

                    <Text mt="md" align={serviceInfo.description.length > 64 ? "justify" : "center"}>
                        {serviceInfo.description}
                    </Text>

                    <Divider />
                </Flex>
            ) : null}
        </>
    );
};
