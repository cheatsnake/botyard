import { Button, Group, Modal, Text, TextInput } from "@mantine/core";
import { FC, useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { useUserContext } from "../contexts/user-context";
import ClientAPI from "../api/client-api";
import { useLoaderContext } from "../contexts/loader-context";
import { errNotify } from "../helpers/notifications";

export const AuthGuard: FC<{ children: JSX.Element }> = (props) => {
    const { user } = useUserContext();

    return <>{user ? props.children : <AuthModal />}</>;
};

const AuthModal = () => {
    const [nickname, setNickname] = useState("");
    const { setUser } = useUserContext();
    const { isLoad, setIsLoad } = useLoaderContext();
    const navigate = useNavigate();

    const backToMainPage = () => {
        if ("startViewTransition" in document) {
            //@ts-ignore
            return document.startViewTransition(() => navigate("/"));
        }

        return navigate("/");
    };

    const login = async () => {
        try {
            setIsLoad(true);
            const newUser = await ClientAPI.createUser(nickname);
            setUser(newUser);
        } catch (error) {
            errNotify((error as Error).message);
        } finally {
            setIsLoad(false);
        }
    };

    useEffect(() => {
        (async () => {
            try {
                setIsLoad(true);
                const currentUser = await ClientAPI.getCurrentUser();
                setUser(currentUser);
            } finally {
                setIsLoad(false);
            }
        })();
    }, []);

    return (
        <>
            {!isLoad ? (
                <Modal centered opened={true} onClose={() => {}} withCloseButton={false}>
                    <Modal.Header p={0}>
                        <Text fw={600} size="lg">
                            👋 Welcome!
                        </Text>
                    </Modal.Header>
                    <Modal.Body p={0} mt="md">
                        <Text>
                            Before you start communicating with the bot, you need to make up a nickname for yourself:
                        </Text>
                        <TextInput
                            value={nickname}
                            onChange={(event) => setNickname(event.currentTarget.value)}
                            mt="md"
                            size="md"
                            placeholder="Your nickname..."
                        />

                        <Group mt="lg" position="right" spacing="xs">
                            <Button variant="subtle" color="gray" opacity={0.7} onClick={backToMainPage}>
                                Cancel
                            </Button>

                            <Button disabled={nickname.length === 0} onClick={login}>
                                Submit
                            </Button>
                        </Group>
                    </Modal.Body>
                </Modal>
            ) : null}
        </>
    );
};
