export {}

const mapping = {
    'I': 1,
    'V': 5,
    'X': 10,
    'L': 50,
    'C': 100,
    'D': 500,
    'M': 1000
}

function romanToInt(s: string): number {
    const characters = s.split('')
    const length = characters.length
    let total = 0
    let cache = 0
    for (let i = 0; i < characters.length; i++) {
        if (i == length - 1) {
            cache += mapping[characters[i]]
            continue
        }

        if (mapping[characters[i]] < mapping[characters[i + 1]]) {
            cache -= mapping[characters[i]]
        } else {
            cache += mapping[characters[i]]
            total += cache
            cache = 0
        }
    }

    if (cache != 0) {
        total += cache
    }

    return total
};

console.log(romanToInt('III'))
console.log(romanToInt('LVIII'))
console.log(romanToInt('MCMXCIV'))