import { ColorScheme, ColorSchemeProvider, MantineThemeOverride, MantineProvider, MantineColor } from "@mantine/core";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { useHotkeys, useLocalStorage } from "@mantine/hooks";
import { Notifications } from "@mantine/notifications";
import GalleryPage from "./modules/gallery/gallery-page";
import ChatPage from "./modules/chat/chat-page";
import { AuthGuard } from "./components/auth-guard";
import { UserProvider } from "./contexts/user-context";
import { LoaderProvider } from "./contexts/loader-context";
import { Loader } from "./components/loader";
import { StorageProvider } from "./contexts/storage-context";
import { PrimaryColorProvider } from "./contexts/primary-color";

const router = createBrowserRouter([
    {
        path: "/",
        element: <GalleryPage />,
    },
    {
        path: "/bot/:id",
        element: (
            <AuthGuard>
                <ChatPage />
            </AuthGuard>
        ),
    },
]);

function App() {
    const [colorScheme, setColorScheme] = useLocalStorage<ColorScheme>({
        key: "color-scheme",
        defaultValue: "dark",
        getInitialValueInEffect: true,
    });

    const [primaryColor, setPrimaryColor] = useLocalStorage<MantineColor>({
        key: "primary-color",
        defaultValue: "green",
        getInitialValueInEffect: true,
    });

    const globalTheme: MantineThemeOverride = {
        colorScheme,
        primaryColor,
    };

    const toggleColorScheme = (value?: ColorScheme) =>
        setColorScheme(value || (colorScheme === "dark" ? "light" : "dark"));

    useHotkeys([["mod+J", () => toggleColorScheme()]]);

    return (
        <ColorSchemeProvider colorScheme={colorScheme} toggleColorScheme={toggleColorScheme}>
            <PrimaryColorProvider primaryColor={primaryColor} setPrimaryColor={setPrimaryColor}>
                <MantineProvider theme={globalTheme} withGlobalStyles withNormalizeCSS>
                    <Notifications limit={3} position="top-center" />
                    <LoaderProvider>
                        <Loader />
                        <UserProvider>
                            <StorageProvider>
                                <RouterProvider router={router} />
                            </StorageProvider>
                        </UserProvider>
                    </LoaderProvider>
                </MantineProvider>
            </PrimaryColorProvider>
        </ColorSchemeProvider>
    );
}

export default App;
