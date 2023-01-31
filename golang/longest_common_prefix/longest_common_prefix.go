package longest_common_prefix

import "sort"

type strArray []string

func (s strArray) Len() int {
	return len(s)
}

func (s strArray) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s strArray) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func longestCommonPrefix(strs []string) string {
	numStrs := len(strs)

	if numStrs > 1 {
		sort.Sort(strArray(strs))

		firstStr := strs[0]

		for i, c := range firstStr {
			for nStr := 1; nStr < numStrs; nStr++ {
				if byte(c) != strs[nStr][i] {
					return firstStr[0:i]
				}
			}
		}

		return firstStr

	} else {
		return strs[0]
	}
}
