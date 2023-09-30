import React, { createContext, useContext, useState, ReactNode } from "react";

interface LoaderContext {
    isLoad: boolean;
    setIsLoad: React.Dispatch<React.SetStateAction<boolean>>;
}

const LoaderContext = createContext<LoaderContext | undefined>(undefined);

export function LoaderProvider({ children }: { children: ReactNode }) {
    const [isLoad, setIsLoad] = useState(false);

    return <LoaderContext.Provider value={{ isLoad, setIsLoad }}>{children}</LoaderContext.Provider>;
}

export function useLoaderContext() {
    const context = useContext(LoaderContext);
    if (context === undefined) {
        throw new Error("useLoader must be used within a LoaderProvider");
    }

    return context;
}
