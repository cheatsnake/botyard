import { Container, Textarea } from "@mantine/core";
import { Header } from "../../components/header";

export const ChatPage = () => {
    return (
        <>
            <Header />
            <Container size="md">{/* <Textarea maxRows={6} minRows={1} size="md" /> */}</Container>
        </>
    );
};
