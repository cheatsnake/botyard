import { MantineColor } from "@mantine/core";
import { createContext, useContext, ReactNode } from "react";

interface PrimaryColorContext {
    primaryColor: MantineColor;
    setPrimaryColor: React.Dispatch<React.SetStateAction<MantineColor>>;
}

const PrimaryColorContext = createContext<PrimaryColorContext | undefined>(undefined);

export const PrimaryColorProvider = ({
    children,
    primaryColor,
    setPrimaryColor,
}: {
    children: ReactNode;
    primaryColor: MantineColor;
    setPrimaryColor: (val: MantineColor | ((prevState: MantineColor) => MantineColor)) => void;
}) => {
    return (
        <PrimaryColorContext.Provider value={{ primaryColor, setPrimaryColor }}>
            {children}
        </PrimaryColorContext.Provider>
    );
};

export const usePrimaryColorContext = () => {
    const context = useContext(PrimaryColorContext);
    if (context === undefined) {
        throw new Error("usePrimaryColorContext must be used within a PrimaryColorProvider");
    }

    return context;
};
