package problem0003

func lengthOfLongestSubstring(s string) int {
	m, max, left := make(map[rune]int), 0, 0
	for idx, c := range s {
		if _, ok := m[c]; ok {
			if idx-left > max {
				max = idx - left
			}
			left = m[c] + 1
		}
		m[c] = idx
	}
	return max
}
