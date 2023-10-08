import { createContext, useContext, useState, ReactNode } from "react";

interface LoaderContext {
    isLoad: boolean;
    setIsLoad: React.Dispatch<React.SetStateAction<boolean>>;
}

const LoaderContext = createContext<LoaderContext | undefined>(undefined);

export const LoaderProvider = ({ children }: { children: ReactNode }) => {
    const [isLoad, setIsLoad] = useState(false);

    return <LoaderContext.Provider value={{ isLoad, setIsLoad }}>{children}</LoaderContext.Provider>;
};

export const useLoaderContext = () => {
    const context = useContext(LoaderContext);
    if (context === undefined) {
        throw new Error("useLoaderContext must be used within a LoaderProvider");
    }

    return context;
};
