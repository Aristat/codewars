/*

This kata is the first of a sequence of four about "Squared Strings".

You are given a string of n lines, each substring being n characters long: For example:

s = "abcd\nefgh\nijkl\nmnop"

We will study some transformations of this square of strings.

Vertical mirror: vert_mirror (or vertMirror or vert-mirror)
vert_mirror(s) => "dcba\nhgfe\nlkji\nponm"
Horizontal mirror: hor_mirror (or horMirror or hor-mirror)
hor_mirror(s) => "mnop\nijkl\nefgh\nabcd"

*/

function vertMirror(strng) {
  return strng.map(s => [...s].reverse().join(''))
}
function horMirror(strng) {
  return strng.reverse()
}
function oper(fct, s) {
  return fct(s.split("\n")).join("\n")
}
