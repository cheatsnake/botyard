import { Container, Divider, Space } from "@mantine/core";
import { Header } from "./header";
import { Watermark } from "./watermark";
import { OrganizationInfo } from "./organization-info";
import { BotList } from "./bot-list";

const GalleryPage = () => {
    return (
        <>
            <Header />
            <Container size="md">
                <OrganizationInfo />
                <Divider p="md" />
                <BotList />
                <Space h="xl" p="xl" />
                <Watermark />
            </Container>
        </>
    );
};

export default GalleryPage;
