# 2264. Largest 3-Same-Digit Number in String - Editorial

**Problem Link:** [https://leetcode.com/problems/largest-3-same-digit-number-in-string/description/?envType=daily-question&envId=2025-08-15](https://leetcode.com/problems/largest-3-same-digit-number-in-string/description/?envType=daily-question&envId=2025-08-15)

## Key Insight
Find the largest substring of 3 consecutive identical digits in the given string.

## Approach
- Iterate through the string with a sliding window of size 3
- Check if all three characters in the window are the same
- Keep track of the largest valid substring found
- Return the largest one (or empty string if none found)

## Time Complexity: O(n)
## Space Complexity: O(1)
