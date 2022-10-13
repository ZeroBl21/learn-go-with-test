package iteration

// Repeat the given number as many times as specified.
func Repeat(character string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}

	return repeated
}
