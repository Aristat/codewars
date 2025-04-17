/*

Given a number, find the permutation with the smallest absolute value (no leading zeros).

-20 => -20
-32 => -23
0 => 0
10 => 10
29394 => 23499

The input will always be an integer.

*/

function minPermutation(n) {  
  const digits = Math.abs(n).toString().split('');
  if (digits.length == 1) {
    return n;
  }
  
  const isNegative = n < 0;
  const results = new Set();

  function permute(arr, l = 0) {
    if (l === arr.length) {
      const str = arr.join('');
      if (str[0] !== '0') {
        results.add(parseInt(str));
      }
      return;
    }

    for (let i = l; i < arr.length; i++) {
      [arr[l], arr[i]] = [arr[i], arr[l]];
      permute(arr, l + 1);
      [arr[l], arr[i]] = [arr[i], arr[l]];
    }
  }

  permute(digits);

  const min = Math.min(...results);
  return isNegative ? -min : min;
}
