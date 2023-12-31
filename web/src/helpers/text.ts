export const abbreviateName = (name: string): string => {
    const words = name.split(" ");
    let abbreviation = "";

    for (const word of words) {
        if (abbreviation.length > 1) break;
        if (word.length > 0) {
            abbreviation += word[0].toUpperCase();
        }
    }

    return abbreviation;
};

export const truncString = (str: string, maxLen: number) => {
    return str.length > maxLen ? str.slice(0, maxLen) + "..." : str;
};
