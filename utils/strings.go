package utils

func Lower(str string) string {
	bytes := []byte(str)
	for i := 0; i < len(bytes); i++ {
		if bytes[i] >= 'A' && bytes[i] <= 'Z' {
			bytes[i] += 32
		}
	}
	return string(bytes)
}

func Upper(str string) string {
	bytes := []byte(str)
	for i := 0; i < len(bytes); i++ {
		if bytes[i] >= 'a' && bytes[i] <= 'z' {
			bytes[i] -= 32
		}
	}
	return string(bytes)
}