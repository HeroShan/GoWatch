package leetcode

func DefangIPaddr(address string) string {
	buf := ""
	begin := 0

	for i, c := range address {
		if c == '.' {
			buf += address[begin:i] + "[.]"
			begin = i + 1
		}
	}

	return buf + address[begin:]
}
