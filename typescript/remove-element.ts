function removeElement(nums: number[], val: number): number {
    const newArray: number[] = []
    for (const item of nums) {
        if (item == val) continue;

        newArray.push(item);
    }
    nums.splice(0, nums.length, ...newArray)

    return newArray.length
}

console.log(removeElement([3,2,2,3], 3))
console.log(removeElement([0,1,2,2,3,0,4,2], 2))
