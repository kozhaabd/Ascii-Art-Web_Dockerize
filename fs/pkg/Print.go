package pkg

func Print(str map[rune][]string, sstr []string) string {
	res := ""
	for i := 0; i < len(sstr); i++ {
		for j := 0; j < 8; j++ {
			for _, v := range sstr[i] {
				res += str[v][j]
			}
			if sstr[i] != "" {
				res += string(rune(10))
			}
		}
		if sstr[i] == "" {
			res += string(rune(10))
		}
	}
	return res
}
