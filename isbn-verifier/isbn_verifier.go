package isbn

func isDigit(chr byte) (digit int, ok bool) {
	if chr < '0' || chr > '9' {
		return digit, ok
	}
	return int(chr - '0'), true
}

func IsValidISBN(isbn string) bool {
	multiplier := 10
	sum := 0

	for i := 0; i < len(isbn); i++ {
		chr := isbn[i]
		if digit, ok := isDigit(chr); ok {
			sum += digit * multiplier
			multiplier--
		} else if multiplier == 1 && chr == 'X' {
			sum += 10
			multiplier--
		} else if chr != '-' {
			return false
		}
	}

	return multiplier == 0 && sum%11 == 0
}

// first answer
/*func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	if len(isbn) != 10 {
		return false
	}

	isX := 0
	if strings.Contains(isbn, "X"){
		isX = 10
		isbn = strings.ReplaceAll(isbn, "X", "")
	}

	sum := 0
	multiplier := 10
	for _, n := range isbn{
		num, err := strconv.Atoi(string(n))
		if err != nil {
			return false
		}
		sum += num * multiplier
		multiplier--
	}

	if isX == 10 {
		sum += 10
	}

	return sum % 11 == 0
}*/
