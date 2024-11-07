export const useDebounce = (fn: Function, delay: number) => {
  let timeout: NodeJS.Timeout;

  return (...args: any[]) => {
    if (timeout) clearTimeout(timeout);
    timeout = setTimeout(() => {
      fn(...args);
    }, delay);
  };
};
