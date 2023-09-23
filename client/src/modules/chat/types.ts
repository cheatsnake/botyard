export interface Message {
    id: number;
    body: string;
    timestamp: Date;
}

export interface Bot {
    id: string;
    name: string;
    description?: string;
    avatar?: string;
}
