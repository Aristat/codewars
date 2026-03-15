class ListNode {
    val: number
    next: ListNode | null
    constructor(val?: number, next?: ListNode | null) {
        this.val = (val===undefined ? 0 : val)
        this.next = (next===undefined ? null : next)
    }
}

function generateList(arr: number[]): ListNode {
    let currentNode = null
    for (const n of arr) {
        currentNode = new ListNode(n, currentNode)
    }
    return currentNode
}

function reverseParseListToNumber(node: ListNode): bigint {
    let result = 0n;
    let multiplier = 1n;
    while (node) {
        result += BigInt(node.val) * multiplier;
        multiplier *= 10n;
        node = node.next;
    }
    return result;
}

function addTwoNumbers(l1: ListNode | null, l2: ListNode | null): ListNode | null {
    let outputNumber = reverseParseListToNumber(l1) + reverseParseListToNumber(l2)
    if (outputNumber == 0n) return generateList([0])

    const outputArray = [];
    while (outputNumber > 0n) {
        outputArray.unshift(Number(outputNumber % 10n));
        outputNumber = outputNumber / 10n;
    }

    return generateList(outputArray)
}

console.log(addTwoNumbers(generateList([2,4,3]), generateList([5,6,4])))
console.log(addTwoNumbers(generateList([0]), generateList([0])))
console.log(addTwoNumbers(generateList([9,9,9,9,9,9,9]), generateList([9,9,9,9])))
console.log(addTwoNumbers(generateList([2,4,9]), generateList([5,6,4,9])))
console.log(addTwoNumbers(generateList([2,4,3]), generateList([5,6,4])))
