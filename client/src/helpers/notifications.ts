import { notifications } from "@mantine/notifications";

export const errNotify = (message: string) => {
    notifications.show({
        withBorder: true,
        title: "Error",
        color: "red",
        autoClose: 9000,
        message,
    });
};

export const successNotify = (message: string) => {
    notifications.show({
        withBorder: true,
        title: "Success",
        color: "green",
        autoClose: 6000,
        message,
    });
};
