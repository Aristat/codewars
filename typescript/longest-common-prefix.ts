function longestCommonPrefix(strs: string[]): string {
    let result = ''
    if (strs.length == 0) return result
    if (strs.length == 1) return strs[0]

    for (let i = 0; i < strs[0].length; i++) {
        let char = strs[0][i]
        for (let j = 1; j < strs.length; j++) {
            if (strs[j][i] != char) return result
        }
        result += char
    }

    return result
}

console.log(longestCommonPrefix(["flower","flow","flight"]))
console.log(longestCommonPrefix(["dog","racecar","car"]))
