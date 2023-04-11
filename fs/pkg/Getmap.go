package pkg

func Map(str []string) map[rune][]string {
	mymap := make(map[rune][]string)
	var temp []string
	for k := 32; k < 127; k++ {
		for i := 0; i < len(str); i++ {
			if i == len(str)-1 && len(temp) != 0 {
				temp = append(temp, str[i])
				mymap[rune(k)] = temp
				temp = []string{}
			}
			if str[i] != "" {
				temp = append(temp, str[i])
			} else if len(temp) != 0 {
				mymap[rune(k)] = temp
				temp = []string{}
				k++
			}
		}
	}
	return mymap
}
