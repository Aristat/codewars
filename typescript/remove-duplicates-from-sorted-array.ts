export {}

function removeDuplicates(nums: number[]): number {
    let previousNum = nums[0]
    let uniqArray = []
    let count = 1
    for (let i = 1; i < nums.length; i++) {
        if (nums[i] !== previousNum) {
            previousNum = nums[i]
            uniqArray.push(nums[i])
            count++
        }
    }

    nums.splice(1, nums.length, ...new Set(uniqArray))
    return count
};

let nums1 = [1, 1, 2]
console.log(removeDuplicates(nums1), nums1)

let nums2 = [0, 0, 1, 1, 1, 2, 2, 3, 3, 4]
console.log(removeDuplicates(nums2), nums2)
