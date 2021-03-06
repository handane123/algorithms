package str

// Lrs returns the longest repeated substring of the specified string.
func Lrs(text string) string {
	n := len(text)
	sa := NewSuffixArray(text)
	lrs := ""
	for i := 1; i < n; i++ {
		length := sa.Lcp(i)
		if length > len(lrs) {
			// lrs = sa.Select(i)[:length]
			lrs = text[sa.Index(i) : sa.Index(i)+length]
		}
	}
	return lrs
}
