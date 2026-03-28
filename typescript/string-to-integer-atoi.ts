const allNumbers = new Set('0123456789'.split(''))

function myAtoi(s: string): number {
    let numbers = ""
    let readNumberRelated = false
    let negative = false

    for (const c of s.split('')) {
        if (c == ' ' && !readNumberRelated) {
        } else if (c == '-' && !readNumberRelated) {
            readNumberRelated = true
            negative = true
        } else if (c == '+' && !readNumberRelated) {
            readNumberRelated = true
            negative = false
        } else if (!allNumbers.has(c)) {
            break
        } else {
            readNumberRelated = true
            numbers += c
        }
    }

    if (numbers == "") {
        return 0
    }

    let result = Number(numbers) * (negative ? -1 : 1)

    if (result > 2147483647) {
        return 2147483647
    } else if (result < -2147483648) {
        return -2147483648
    } else {
        return result
    }
}

console.log(myAtoi('42'))
console.log(myAtoi(' -042'))
console.log(myAtoi('1337c0d3'))
console.log(myAtoi('0-1'))
console.log(myAtoi('words and 987'))
console.log(myAtoi('-91283472332'))
console.log(myAtoi('+-12'))
console.log(myAtoi('   +0 123'))