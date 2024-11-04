package main

func reverseString(s []byte) []byte {
	return reverse(s, 0, len(s)-1)
}

func reverse(s []byte, l int, r int) []byte {
	if l >= r {
		return s
	}

	s[l], s[r] = s[r], s[l]
	return reverse(s, l+1, r-1)
}
