# 342. Power of Four - Editorial

**Problem Link:** [https://leetcode.com/problems/power-of-four/description/?envType=daily-question&envId=2025-08-15](https://leetcode.com/problems/power-of-four/description/?envType=daily-question&envId=2025-08-15)

## Key Insight
A number is a power of four if it can be repeatedly divided by 4 until it becomes 1.

## Approach
- Check if `n` is 0 (not a power of four)
- Keep dividing by 4 as long as it's divisible
- If final result is 1, it was a power of four

## Time Complexity: O(log n)
## Space Complexity: O(1)
