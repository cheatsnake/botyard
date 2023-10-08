import { LoadingOverlay } from "@mantine/core";
import { useLoaderContext } from "../contexts/loader-context";

export const Loader = () => {
    return (
        <LoadingOverlay
            visible={useLoaderContext().isLoad}
            transitionDuration={200}
            overlayColor="dark"
            loaderProps={{ size: "lg" }}
        />
    );
};
