import { notifications } from "@mantine/notifications";

export const errNotify = (message: string) => {
    notifications.show({
        withBorder: true,
        title: "Error",
        color: "red",
        autoClose: 7000,
        message,
    });
};
