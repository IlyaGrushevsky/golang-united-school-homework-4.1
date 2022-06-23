package reverse_string

func ReverseString(input string) (output string) {
	r := []rune(input)
	var src []rune
	for i := len(r) - 1; i >= 0; i-- {
		src = append(src, r[i])
	}
	output = string(src)
	return output
}
