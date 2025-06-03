package utils

// "000000000000000000000000a614f803b6fd780986a42c78ec9c7f77e6ded13c" -> "a614f803b6fd780986a42c78ec9c7f77e6ded13c"
func TrimZeroes(address string) string {
	idx := 0
	for ; idx < len(address); idx++ {
		if address[idx] != '0' {
			break
		}
	}
	return address[idx:]
}
