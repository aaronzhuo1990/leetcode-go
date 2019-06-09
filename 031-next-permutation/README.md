# Problem

Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.

If such arrangement is not possible, it must rearrange it as the lowest possible order (ie, sorted in ascending order).

The replacement must be in-place and use only constant extra memory.

Here are some examples. Inputs are in the left-hand column and its corresponding outputs are in the right-hand column.

1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1

# Solutions

## First Solution

Example: 

1　　2　　7　　4　　3　　1

1. Find the longest descending permutation from the back, which is `7 4 3 1` in this case.

2. Sort the permutation in ascending order, which becomes `1 3 4 7`

3. Switch the position for the element before the permutation (which is `2` in this case) and the first element which value is greater than it.

4. The final order becomes: 1 3 1 2 4 7