# Problem

33. Search in Rotated Sorted Array
Medium

2437

313

Favorite

Share
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).

You are given a target value to search. If found in the array return its index, otherwise return -1.

You may assume no duplicate exists in the array.

Your algorithm's runtime complexity must be in the order of O(log n).

Example 1:

Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4


Example 2:

Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1

# Solutions

## First Solution

It is advanced binary search. We need to find the rotated num and do the following calculation:

1. mid := (left + right) / 2.

2. rotatedMid := (mid + rotatedNum) % len(nums).

3. Do the rest of binary search.
