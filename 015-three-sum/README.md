# Problem

Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note:

The solution set must not contain duplicate triplets.

Example:

Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]

# Solutions

## First Solution:

1. Sort the array.
2. Have an index `left` to loop through the whole array.
3. Within the loop, have another the index `middle` which starts next to the index `left` and another index `right` starts from the end of array. middle++ and right-- to find 
each unique triple.


