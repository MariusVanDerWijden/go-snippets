package fuzz

func CoolFunction(index int, b []byte) byte {
	if index < 0 || index >= len(b) {
		return 0
	}
	return b[index]
}
