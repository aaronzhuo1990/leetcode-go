# Problem

Given an array nums of n integers and an integer target, find three integers in nums such that the sum is closest to target. Return the sum of the three integers. You may assume that each input would have exactly one solution.

Example:

Given array nums = [-1, 2, 1, -4], and target = 1.

The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
# Solutions

## First Solution:

1. Sort the array.
2. Have an index `left` to loop through the whole array.
3. Within the loop, have another the index `middle` which starts next to the index `left` and another index `right` starts from the end of array. middle++ and right-- to find the closet sums 


