export const jsonRequestParams = (method: string, body: { [key: string]: any }) => {
    return {
        method,
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(body),
    };
};

export const queryParams = (params: { [key: string]: any }) => {
    return new URLSearchParams(params).toString();
};
