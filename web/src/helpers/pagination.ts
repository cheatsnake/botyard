export const defineNextPageLimit = (totalItems: number, defaultLimit: number) => {
    let page = 1;
    let limit = 1;

    for (let divisor = defaultLimit; divisor >= 3; divisor--) {
        if (totalItems % divisor === 0) {
            limit = divisor;
            page = totalItems / divisor + 1;
            break;
        }
    }

    return { page, limit };
};
