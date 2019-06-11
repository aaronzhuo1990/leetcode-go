# Problem

Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.

Your algorithm's runtime complexity must be in the order of O(log n).

If the target is not found in the array, return [-1, -1].

Example 1:

Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
Example 2:

Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]

# Solutions

## First Solution

1. Find the first element which value is target.

2. Split the array into two by the element.

3. Find the left index from the first part of the array.

4. Find the right index from the second part of the array.

