package reverse

func Reverse(input string) string {
	reversed := ""
	for _, r := range input {
		reversed = string(r) + reversed
	}
	return reversed
}
