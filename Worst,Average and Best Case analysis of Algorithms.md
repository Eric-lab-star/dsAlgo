---
date created: 2023-06-04 08:19
---

# Popular Notations is Complexity Analysis of Algorithms

## Big-O Notation

- Worst case time complexity
- maximum amount of time an algorithm requires to consider all input values

## Omega Notation

- Best case
- Minimum amount of time

## Theta Notation

- Average case

## Measurement of Complexity of an Algorithm

1. Worst Case Analysis
   1. Upper bound on the running time
2. Best Case Analysis
3. Average Case Analysis

## Example with theri complexiy analysis:

1. Linear Search Algorithm

```js
// javascript implementation of the approach

// Linearly search x in arr. If x is present then
// return the index, otherwise return -1
function search(arr, n, x) {
  var i;
  for (i = 0; i < n; i++) {
    if (arr[i] == x) {
      return i;
    }
  }
  return -1;
}

/* Driver program to test above functions */

var arr = [1, 10, 30, 15];
var x = 30;
var n = arr.length;
console.log(x + " is present at index " + search(arr, n, x));

```

> [!output]+ Output
> 30 is present at index 2

Time Complexity Analysis

- Best Case: O(1), element to be searched is ont the first index of the given list, So, the number of comparisons, in this cases, is 1.
