/**
 * im up and a racing like a forgeign entity
 * @param date The date to format
 * @returns Date formatted as yyyy-mm-dd
 */
export function dateToString(date: Date): string {
  const year = date.getFullYear();
  const month = date.getMonth() + 1;
  const day = date.getDate();
  return `${year}-${month}-${day}`;
}

/**
 * Format a date as a unix timestamp
 * @param date The date to format
 * @returns Date formatted as UNIX timestamp
 */
export function dateToTimestamp(date: Date): string {
  const dayStart = new Date(date);
  dayStart.setHours(0, 0, 0, 0);
  return String(dayStart.getTime() / 1000);
}
