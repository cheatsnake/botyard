export function debounce<T extends unknown[]>(func: (...args: T) => void, delay = 1000): (...args: T) => void {
    let timer: NodeJS.Timeout | null = null;
    return (...args: T) => {
        if (timer) clearTimeout(timer);
        timer = setTimeout(() => {
            func.call(null, ...args);
        }, delay);
    };
}
