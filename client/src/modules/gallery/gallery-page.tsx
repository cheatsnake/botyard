import { Container, Flex } from "@mantine/core";
import { Watermark } from "./watermark";
import { OrganizationInfo } from "./organization-info";
import { BotList } from "./bot-list";

const GalleryPage = () => {
    return (
        <>
            <Container size="md">
                <Flex direction="column" mih="100vh" justify="center" py="2.4rem" gap="md">
                    <OrganizationInfo />
                    <BotList />
                    <Watermark />
                </Flex>
            </Container>
        </>
    );
};

export default GalleryPage;
