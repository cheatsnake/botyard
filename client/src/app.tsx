import { ColorScheme, ColorSchemeProvider, MantineProvider } from "@mantine/core";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { GalleryPage } from "./modules/gallery/gallery-page";
import { ChatPage } from "./modules/chat/chat-page";
import { useHotkeys, useLocalStorage } from "@mantine/hooks";

const router = createBrowserRouter([
    {
        path: "/",
        element: <GalleryPage />,
    },
    {
        path: "/bot/:id",
        element: <ChatPage />,
    },
]);

function App() {
    const [colorScheme, setColorScheme] = useLocalStorage<ColorScheme>({
        key: "mantine-color-scheme",
        defaultValue: "light",
        getInitialValueInEffect: true,
    });

    const toggleColorScheme = (value?: ColorScheme) =>
        setColorScheme(value || (colorScheme === "dark" ? "light" : "dark"));

    useHotkeys([["mod+J", () => toggleColorScheme()]]);

    return (
        <ColorSchemeProvider colorScheme={colorScheme} toggleColorScheme={toggleColorScheme}>
            <MantineProvider theme={{ colorScheme, primaryColor: "green" }} withGlobalStyles withNormalizeCSS>
                <RouterProvider router={router} />
            </MantineProvider>
        </ColorSchemeProvider>
    );
}

export default App;
