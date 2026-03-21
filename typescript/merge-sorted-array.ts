export {}

function merge(nums1: number[], m: number, nums2: number[], n: number): void {
    let num1Index = 0
    let num2Index = 0
    let result: number[] = []

    for (let i = 0; i < m + n; i++) {
        if (num1Index >= m) {
            result.push(nums2[num2Index])
            num2Index++
            continue
        }
        if (num2Index >= n) {
            result.push(nums1[num1Index])
            num1Index++
            continue
        }

        if (nums1[num1Index] >= nums2[num2Index]) {
            result.push(nums2[num2Index])
            num2Index++
        } else {
            result.push(nums1[num1Index])
            num1Index++
        }
    }

    nums1.splice(0, nums1.length, ...result)
}

let nums11 = [1,2,3,0,0,0], m1 = 3, nums12 = [2,5,6], n1 = 3
merge(nums11, m1, nums12, n1)
console.log(nums11)

let nums21 = [0], m2 = 0, nums22 = [1], n2 = 1
merge(nums21, m2, nums22, n2)
console.log(nums21)
