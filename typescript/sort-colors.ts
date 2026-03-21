export {}

function sortColors(nums: number[]): void {
    const array0: number[] = []
    const array1: number[] = []
    const array2: number[] = []

    for(const n of nums) {
        if (n == 0) {
            array0.push(n)
        } else if (n == 1) {
            array1.push(n)
        } else {
            array2.push(n)
        }
    }

    nums.splice(0, nums.length, ...array0, ...array1, ...array2)
}

let nums1 = [2,0,2,1,1,0]
sortColors(nums1)
console.log(nums1)

let nums2 = [2,0,1]
sortColors(nums2)
console.log(nums2)