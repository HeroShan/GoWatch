package leetcode

import "fmt"

func getNum(num string) int {
	switch num {
	case "I":
		return 1
	case "IV":
		return 4
	case "V":
		return 5
	case "IX":
		return 9
	case "X":
		return 10
	case "XL":
		return 40
	case "L":
		return 50
	case "XC":
		return 90
	case "C":
		return 100
	case "CD":
		return 400
	case "D":
		return 500
	case "CM":
		return 900
	case "M":
		return 1000
	default:
		return -1
	}
}

func compare(a, b uint8) bool {
	var str1 string = fmt.Sprintf("%c", a)
	var str2 string = fmt.Sprintf("%c", b)
	return getNum(str1) >= getNum(str2)
}

func getString(a ...uint8) string {
	var format string = "%c"
	if len(a) > 1 {
		format += "%c"
		return fmt.Sprintf(format, a[0], a[1])
	}
	return fmt.Sprintf(format, a[0])

}

func RomanToInt(s string) int {
	var strLen int = len(s)
	var sum int = 0
	var reachEnd bool = false
	for i := 0; i < strLen-1; i++ {
		if !compare(s[i], s[i+1]) {
			if i+1 == strLen-1 {
				reachEnd = true
			}
			sum += getNum(getString(s[i], s[i+1]))
			i++
		} else {
			sum += getNum(getString(s[i]))
		}
	}
	if !reachEnd {
		sum += getNum(getString(s[strLen-1]))
	}
	return sum
}
