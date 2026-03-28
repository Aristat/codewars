function maxArea(height: number[]): number {
    let left = 0;
    let right = height.length - 1;
    let max = 0;

    while (left < right) {
        const area = (right - left) * Math.min(height[left], height[right]);
        max = Math.max(max, area);

        if (height[left] < height[right]) {
            left++;
        } else {
            right--;
        }
    }

    return max;
}

console.log(maxArea([1,8,6,2,5,4,8,3,7]))
console.log(maxArea([2,3,4,5,18,17,6]))