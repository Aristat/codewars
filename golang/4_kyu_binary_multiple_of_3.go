/*

In this kata, your task is to create a regular expression capable of evaluating binary strings (strings with only 1s and 0s) and determining whether the given string represents a number divisible by 3.

Take into account that:

An empty string might be evaluated to true (it's not going to be tested, so you don't need to worry about it - unless you want)
The input should consist only of binary digits - no spaces, other digits, alphanumeric characters, etc.
There might be leading 0s.

A = 1B + 0A
B = 1A + 0C
C = 1C + 0B

C = 1*0B Fix recursion

B = 1A + 0(1*0B)
B = 01*0B + 1A
B = (01*0)*1A Fix recursion

A = 1(01*0)*1A + 0A
A = (1(01*0)*1 + 0)A
A = (1(01*0)*1 + 0)* Fix recursion

*/

package kata

var MultipleOf3Regex = "^(1(01*0)*1|0)+$"
