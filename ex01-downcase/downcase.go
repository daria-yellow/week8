package downcase

func Downcase(s string) (string, error) {
	m := []int32(s)

	for i, val := range m {
		if val >= 65 && val <= 90 {
			m[i] = m[i] + 32
		}
	}

	return string(m), nil
}
