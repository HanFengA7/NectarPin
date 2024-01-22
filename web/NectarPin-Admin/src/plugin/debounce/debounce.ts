export function debounce(func:any, delay:any) {
    let timerId : any

    return function (...args: []) {
        clearTimeout(timerId);

        timerId = setTimeout(() => {
            func.apply(func, args);
        }, delay);
    };
}