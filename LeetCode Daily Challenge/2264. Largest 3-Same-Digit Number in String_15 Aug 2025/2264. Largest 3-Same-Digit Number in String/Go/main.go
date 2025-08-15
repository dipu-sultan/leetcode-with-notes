package main

func largestGoodInteger(num string) string {

	ans := ""
	for i := 0; i+2 < len(num); i++ {
		if num[i] == num[i+1] && num[i] == num[i+2] {
			current := num[i : i+3]
			if current > ans {
				ans = current
			}
		}
	}
	return ans
}