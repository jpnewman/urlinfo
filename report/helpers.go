package report

func convertStringToBytes(s string) []byte {
	b := []byte(s)

	if len(b) > 0 {
		if b[len(b)-1] != '\n' {
			b = append(b, '\n')
		}
	} else {
		b = append(b, '\n')
	}

	return b
}
