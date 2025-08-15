# 326. Power of Three - Editorial

**Problem Link:** [https://leetcode.com/problems/power-of-three/description/?envType=daily-question&envId=2025-08-15](https://leetcode.com/problems/power-of-three/description/?envType=daily-question&envId=2025-08-15)

## Key Insight
A number is a power of three if it can be repeatedly divided by 3 until it becomes 1.

## Approach
- Check if `n` is 0 (not a power of three)
- Keep dividing by 3 as long as it's divisible
- If final result is 1, it was a power of three

## Time Complexity: O(log n)
## Space Complexity: O(1)
