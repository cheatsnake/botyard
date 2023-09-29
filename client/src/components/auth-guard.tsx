import { Button, Group, Modal, Text, TextInput } from "@mantine/core";
import { FC, useState } from "react";
import { useNavigate } from "react-router-dom";

export const AuthGuard: FC<{ children: JSX.Element }> = (props) => {
    const [isAuth, setIsAuth] = useState(false);

    return <>{isAuth ? props.children : <AuthModal setIsAuth={setIsAuth} />}</>;
};

const AuthModal: FC<{ setIsAuth: React.Dispatch<React.SetStateAction<boolean>> }> = (props) => {
    const navigate = useNavigate();

    const backToMainPage = () => {
        if ("startViewTransition" in document) {
            //@ts-ignore
            return document.startViewTransition(() => navigate("/"));
        }

        return navigate("/");
    };

    return (
        <>
            <Modal centered opened={true} onClose={close} withCloseButton={false}>
                <Modal.Header p={0}>
                    <Text fw={600} size="lg">
                        👋 Welcome!
                    </Text>
                </Modal.Header>
                <Modal.Body p={0} mt="md">
                    <Text>
                        Before you start communicating with the bot, you need to make up a nickname for yourself:
                    </Text>
                    <TextInput mt="md" size="md" placeholder="Your nickname..." />

                    <Group mt="lg" position="right" spacing="sm">
                        <Button variant="subtle" color="gray" onClick={backToMainPage}>
                            Cancel
                        </Button>

                        <Button onClick={() => props.setIsAuth(true)}>Submit</Button>
                    </Group>
                </Modal.Body>
            </Modal>
        </>
    );
};
