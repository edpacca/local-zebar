export function checkThreshold(value: number, threshold: number, trueAbove: boolean) {
    return trueAbove ? value >= threshold : value <= threshold;
}
