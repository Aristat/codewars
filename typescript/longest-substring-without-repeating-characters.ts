function lengthOfLongestSubstring(s: string): number {
    const seen = new Map<string, number>();
    let max = 0;
    let left = 0;
    let right = 0;

    for (right; right < s.length; right++) {
        const c = s[right];

        if (seen.has(c) && seen.get(c) >= left) {
            left = seen.get(c) + 1;
        }

        seen.set(c, right);
        max = Math.max(max, right - left + 1);
    }

    return max;
}

console.log(lengthOfLongestSubstring("abcabcbb"))
console.log(lengthOfLongestSubstring("bbbbb"))
console.log(lengthOfLongestSubstring("pwwkew"))
console.log(lengthOfLongestSubstring(" "))
console.log(lengthOfLongestSubstring("dvdf"))