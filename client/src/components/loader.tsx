import { LoadingOverlay } from "@mantine/core";
import { useLoaderContext } from "../contexts/loader-context";

export const Loader = () => {
    return <LoadingOverlay visible={useLoaderContext().isLoad} loaderProps={{ size: "lg" }} />;
};
