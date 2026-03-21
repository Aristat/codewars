export {}

const mapping = new Map<number, string>(
    [[1000, 'M'], [500, 'D'], [100, 'C'], [50, 'L'], [10, 'X'], [5, 'V'], [1, 'I']]
)
const subtractiveMapping = {
    4: 'IV',
    9: 'IX',
    40: 'XL',
    90: 'XC',
    400: 'CD',
    900: 'CM'
}

function digitCount(num: number): number {
    if (num === 0) return 0

    return Math.floor(Math.log10(num))
}

function firstNum(num: number): number {
    while (num >= 10) {
        num = Math.floor(num / 10);
    }
    return num;
}

function parsing(num: number, result: string[]): number {
    if (firstNum(num) === 4) {
        const digits = digitCount(num)
        const divNumber = Math.pow(10,digits)

        num %= divNumber
        result.push(subtractiveMapping[4 * divNumber])
    } else if (firstNum(num) === 9) {
        const digits = digitCount(num)
        const divNumber = Math.pow(10,digits)

        num %= divNumber
        result.push(subtractiveMapping[9 * divNumber])
    } else {
        for (const [key, value] of mapping) {
            if (num < key) continue

            num -= key
            result.push(value)
            break
        }
    }

    if (num > 0) {
        num = parsing(num, result)
    }

    return num
}

function intToRoman(num: number): string {
    let r: string[] = []
    parsing(num, r)
    return r.join('')
}

console.log(intToRoman(3749))
console.log(intToRoman(58))
console.log(intToRoman(1994))

