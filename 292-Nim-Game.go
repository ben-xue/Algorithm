func canWinNim(n int) bool {
	var m int = 3
	if n <= m{
		return true
	}

	return 0 != (n % (m+1))
}