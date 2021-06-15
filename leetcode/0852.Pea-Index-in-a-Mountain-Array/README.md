# [852. Peak Index in a Mountain Array](https://leetcode-cn.com/problems/peak-index-in-a-mountain-array/)

## 题目

Let's call an array arr a mountain if the following properties hold:

- `arr.length >= 3` 
- There exists some `i` with `0 < i < arr.length - 1` such that:
    - `arr[0] < arr[1] < ... arr[i-1] < arr[i]`
    - `arr[i] > arr[i+1] > ... > arr[arr.length - 1]`

Given an integer array arr that is **guaranteed** to be a mountain, return any i such that `arr[0] < arr[1] < ... arr[i - 1] < arr[i] > arr[i + 1] > ... > arr[arr.length - 1]`.

 

**Example 1:**

    Input: arr = [0,1,0]
    Output: 1

**Example 2:**

    Input: arr = [0,2,1,0]
    Output: 1

**Example 3:**

    Input: arr = [0,10,5,2]
    Output: 1

**Example 4:**

    Input: arr = [3,4,5,1]
    Output: 2

**Example 5:**

    Input: arr = [24,69,100,99,79,78,67,36,26,19]
    Output: 2


**Constraints:**

- `3 <= arr.length <= 104`
- `0 <= arr[i] <= 106`
- `arr` is **guaranteed** to be a mountain array.


**Follow up:** Finding the O(n) is straightforward, could you find an O(log(n)) solution?


## 题目大意

我们把符合下列属性的数组 A 称作山脉：

- A.length >= 3
- 存在 0 < i < A.length - 1 使得A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1]
  给定一个确定为山脉的数组，返回任何满足 A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1] 的 i 的值。

提示：

- 3 <= A.length <= 10000
- 0 <= A[i] <= 10^6
- A 是如上定义的山脉
 
 
## 解题思路

- 从左往右找到峰值，O(n)
- 二分法找峰值，O(log(n))
