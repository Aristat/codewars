function reverse(x: number): number {
    const maxNumber = 2147483647
    const minNumber = -2147483648

    let reverseNumber = parseInt(x.toString().split('').reverse().join(''))
    reverseNumber = x < 0 ? -reverseNumber : reverseNumber
    return (reverseNumber <= maxNumber && reverseNumber >= minNumber) ? reverseNumber : 0
}

console.log(reverse(123))
console.log(reverse(120))
console.log(reverse(-123))
