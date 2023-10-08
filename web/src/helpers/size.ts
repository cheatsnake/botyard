export const parseByteSize = (size: number) => {
    if (size >= 1048576) {
        return (size / 1048576).toFixed(2) + " MB";
    } else {
        return (size / 1024).toFixed(0) + " KB";
    }
};
