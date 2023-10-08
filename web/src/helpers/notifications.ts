import { notifications } from "@mantine/notifications";

export const errNotify = (message: string, title?: string) => {
    notifications.show({
        withBorder: true,
        title: title ?? "Error",
        color: "red",
        autoClose: 9000,
        message,
    });
};

export const warnNotify = (message: string, title?: string) => {
    notifications.show({
        withBorder: true,
        title: title ?? "Warning",
        color: "yellow",
        autoClose: 7500,
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
