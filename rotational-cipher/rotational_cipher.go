package rotationalcipher

func RotationalCipher(plain string, shiftKey int) string {
	cipherRune := []rune{}

	for _, r := range plain {
		base := 0
		switch {
		case r >= 'A' && r <= 'Z':
			base = int('A')
		case r >= 'a' && r <= 'z':
			base = int('a')
		}

		if base != 0 {
			r = rune(((int(r) - base + shiftKey) % 26) + base)
		}
		cipherRune = append(cipherRune, r)
	}
	return string(cipherRune)
}

// first submit
/*func RotationalCipher(plain string, shiftKey int) string {
	plainByte := []byte(plain)
	ciphers := ""
	for _, b := range plainByte {
		tempByte := b
		tempSum := byte(0)
		if (b < 97 && b > 90) || b < 65 || b > 122 {
			ciphers += string(b)
			continue
		}
		tempByte = b + byte(shiftKey)
		if tempByte > byte(122) && b > 96 {
			tempSum = tempByte - byte(122)
			ciphers += string(96 + tempSum)
		} else if tempByte > byte(90) && b < 91 {
			tempSum = tempByte - byte(90)
			ciphers += string(64 + tempSum)
		} else {
			ciphers += string(b + byte(shiftKey))
		}
	}

	return ciphers
}*/
