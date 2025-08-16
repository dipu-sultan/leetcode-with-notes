# 1323. Maximum 69 Number - Editorial

**Problem Link:** [https://leetcode.com/problems/maximum-69-number/description/?envType=daily-question&envId=2025-08-17](https://leetcode.com/problems/maximum-69-number/description/?envType=daily-question&envId=2025-08-17)

## Key Insight
Find the position of the first digit 6 from right to left, then add 3 × 10^position to convert it to 9.

## Approach
- Iterate through digits from right to left (least significant to most significant)
- Track the position of the first 6 encountered
- If no 6 found, return original number
- Otherwise, add 3 × 10^position to convert that 6 to 9

## Time Complexity: O(log n)
## Space Complexity: O(1)
