import { test, describe, expect } from "vitest";
import { checkThreshold } from "../utils";

describe('checkThreshold', () => {
    test('should return true when value is greater than or equal to threshold and warnAbove is true', () => {
        expect(checkThreshold(50, 40, true)).toBe(true);
        expect(checkThreshold(40, 40, true)).toBe(true);
        expect(checkThreshold(30, 40, true)).toBe(false);
    });

    test('should return true when value is less than or equal to threshold and warnAbove is false', () => {
        expect(checkThreshold(30, 40, false)).toBe(true);
        expect(checkThreshold(40, 40, false)).toBe(true);
        expect(checkThreshold(50, 40, false)).toBe(false);
    });

    test('should handle edge cases correctly', () => {
        expect(checkThreshold(0, 0, true)).toBe(true);
        expect(checkThreshold(0, 0, false)).toBe(true);
        expect(checkThreshold(-10, -10, true)).toBe(true);
        expect(checkThreshold(-10, -10, false)).toBe(true);
    });
});
