package stringbuilder

func build(str []rune, random rune, i int) []rune {
	var res []rune
	runes := make([]rune, len(str)-1)
	for i := range runes {
		runes[i] = ' '
	}
	res = str[:i]
	res = append(res, random)
	res = append(res, runes[i+1:]...)

	return res
}
