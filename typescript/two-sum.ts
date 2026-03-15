function twoSum(nums: number[], target: number): number[] {
    const mapIndexes: Map<number, number> = new Map();
    for (let i = 0; i < nums.length; i++) {
        mapIndexes.set(nums[i], i)
    }

    for (let i = 0; i < nums.length; i++) {
        const diff = target - nums[i];
        if (mapIndexes.has(diff) && mapIndexes.get(diff) !== i) {
            return [i, mapIndexes.get(diff)];
        }
    }

    return []
};

console.log(twoSum([2,7,11,15], 9))
console.log(twoSum([3,2,4], 6))
console.log(twoSum([3,3], 6))
console.log(twoSum([-3,4,3,90], 0))