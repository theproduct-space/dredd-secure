// dateUtils.ts

/**
 * Convert a JavaScript Date object to epoch seconds.
 *
 * @param date - The JavaScript Date object.
 * @returns The epoch timestamp in seconds or null if the input is null.
 */
export const dateToEpochSeconds = (date: Date | null): number | null => {
  return date ? Math.floor(date.getTime() / 1000) : null;
};

/**
 * Convert an epoch timestamp in seconds to a JavaScript Date object.
 *
 * @param epoch - The epoch timestamp in seconds.
 * @returns The JavaScript Date object or null if the input is null/undefined.
 */
export const epochSecondsToDate = (
  epoch: number | undefined | null,
): Date | null => {
  return epoch !== null && epoch !== undefined ? new Date(epoch * 1000) : null;
};
