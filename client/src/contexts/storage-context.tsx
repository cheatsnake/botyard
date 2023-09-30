import React, { createContext, useContext, useState, ReactNode } from "react";
import { Bot, ServiceInfo } from "../api/types";
import ClientAPI from "../api/client-api";

interface StorageContext {
    bots: Bot[];
    serviceInfo: ServiceInfo | undefined;
    loadBots: () => Promise<Bot[]>;
    setServiceInfo: React.Dispatch<React.SetStateAction<ServiceInfo | undefined>>;
}

const StorageContext = createContext<StorageContext | undefined>(undefined);

export const StorageProvider = ({ children }: { children: ReactNode }) => {
    const [bots, setBots] = useState<Bot[]>([]);
    const [serviceInfo, setServiceInfo] = useState<ServiceInfo>();

    const loadBots = async () => {
        try {
            if (bots.length !== 0) return bots;
            const allBots = await ClientAPI.getAllBots();
            setBots(allBots);
            return allBots;
        } catch (error) {
            throw error;
        }
    };

    return (
        <StorageContext.Provider value={{ bots, serviceInfo, loadBots, setServiceInfo }}>
            {children}
        </StorageContext.Provider>
    );
};

export const useStorageContext = () => {
    const context = useContext(StorageContext);
    if (context === undefined) {
        throw new Error("useStorageContext must be used within a StorageProvider");
    }

    return context;
};
